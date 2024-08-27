package mine

import (
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/joaofnds/blockchain/block"
	"github.com/joaofnds/blockchain/hash"
)

type Concurrent struct {
	hasher     hash.Hasher
	numWorkers int
	batchSize  uint64
}

var _ Miner = &Concurrent{}

func NewConcurrent(hasher hash.Hasher) *Concurrent {
	return &Concurrent{
		hasher:     hasher,
		numWorkers: runtime.NumCPU(),
		batchSize:  10_000_000,
	}
}

func (miner *Concurrent) Mine(blk *block.Block, difficulty int) {
	var nonce uint64
	found := make(chan bool)

	prefix := hashPrefix(difficulty)

	var once sync.Once

	for i := 0; i < miner.numWorkers; i++ {
		go func() {

			serialize := blockSerializer(blk)

			for {
				startNonce := atomic.AddUint64(&nonce, miner.batchSize) - miner.batchSize

				for localNonce := startNonce; localNonce < startNonce+miner.batchSize; localNonce++ {
					select {
					case <-found:
						return
					default:
						hash := miner.hasher.Hash(serialize(localNonce))

						if hasPrefix(hash, prefix) {
							once.Do(func() { close(found) })
							blk.Hash = hash
							blk.Nonce = localNonce
							return
						}
					}
				}
			}
		}()
	}

	<-found
}
