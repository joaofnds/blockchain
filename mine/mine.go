package mine

import (
	"strings"

	"github.com/joaofnds/blockchain/block"
	"github.com/joaofnds/blockchain/hash"
)

type Miner struct {
	hasher hash.Hasher
}

func New(hasher hash.Hasher) *Miner {
	return &Miner{hasher: hasher}
}

func (miner *Miner) Mine(blk block.Block, difficulty int) {
	prefix := miner.makePrefix(difficulty)

	for !strings.HasPrefix(blk.Hash, prefix) {
		blk.IncNonce()
		blk.SetHash(miner.hasher.Hash(blk.Serialize()))
	}
}

func (miner *Miner) Validate(blk block.Block, difficulty int) bool {
	return strings.HasPrefix(blk.Hash, miner.makePrefix(difficulty))
}

func (miner *Miner) makePrefix(size int) string {
	prefix := make([]byte, size)

	for i := range prefix {
		prefix[i] = '0'
	}

	return string(prefix)
}
