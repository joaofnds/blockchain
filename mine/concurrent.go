package mine

import (
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"

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
	var nonce uint64
	var found int32 = 0

	prefix := hashPrefix(difficulty)

	var wg sync.WaitGroup
	wg.Add(miner.numWorkers)
	for i := 0; i < miner.numWorkers; i++ {
		go func() {
			defer wg.Done()

			buf := preNonceBuffer(blk)
			lenBeforeNonce := buf.Len()

			for atomic.LoadInt32(&found) == 0 {
				localNonce := atomic.AddUint64(&nonce, 1)

				buf.Truncate(lenBeforeNonce)
				buf.WriteString(strconv.FormatUint(localNonce, 10))
				hash := miner.hasher.Hash(buf.Bytes())

				if hasPrefix(hash, prefix) {
					if atomic.CompareAndSwapInt32(&found, 0, 1) {
						blk.Hash = hash
						blk.Nonce = localNonce
					}
					break
				}
			}
		}()
	}

	wg.Wait()
}
