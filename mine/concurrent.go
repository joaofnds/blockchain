package mine

import (
	"runtime"
	"strconv"
	"sync"

	"github.com/joaofnds/blockchain/block"
	"github.com/joaofnds/blockchain/hash"
)

type Concurrent struct {
	hasher     hash.Hasher
	numWorkers int
}

var _ Miner = &Concurrent{}

func NewConcurrent(hasher hash.Hasher) *Concurrent {
	return &Concurrent{
		hasher:     hasher,
		numWorkers: runtime.NumCPU(),
	}
}

func (miner *Concurrent) Mine(blk *block.Block, difficulty int) {
	const batchSize = 10_000_000
	var nonce uint64

	prefix := hashPrefix(difficulty)

	nonceChan := make(chan uint64)
	resultChan := make(chan struct{}, 1)
	doneChan := make(chan struct{})

	var once sync.Once

	for i := 0; i < miner.numWorkers; i++ {
		go func() {
			buf := preNonceBuffer(blk)
			lenBeforeNonce := buf.Len()

			for {
				select {
				case startNonce, ok := <-nonceChan:
					if !ok {
						return
					}
					for localNonce := startNonce; localNonce < startNonce+batchSize; localNonce++ {
						buf.Truncate(lenBeforeNonce)
						buf.WriteString(strconv.FormatUint(localNonce, 10))
						hash := miner.hasher.Hash(buf.Bytes())

						if hasPrefix(hash, prefix) {
							select {
							case resultChan <- struct{}{}:
								blk.Nonce = localNonce
								blk.Hash = hash
								once.Do(func() { close(doneChan) })
							case <-doneChan:
							}
							return
						}

						select {
						case <-doneChan:
							return
						default:
						}
					}
				case <-doneChan:
					return
				}
			}
		}()
	}

	go func() {
		for {
			select {
			case nonceChan <- nonce:
				nonce += batchSize
			case <-doneChan:
				close(nonceChan)
				return
			}
		}
	}()

	<-resultChan
}
