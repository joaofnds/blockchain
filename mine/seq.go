package mine

import (
	"github.com/joaofnds/blockchain/block"
	"github.com/joaofnds/blockchain/hash"
)

type SeqMiner struct {
	hasher hash.Hasher
}

var _ Miner = &SeqMiner{}

func NewSeq(hasher hash.Hasher) *SeqMiner {
	return &SeqMiner{hasher: hasher}
}

func (miner *SeqMiner) Mine(blk *block.Block, difficulty int) {
	serialize := blockSerializer(blk)

	prefix := hashPrefix(difficulty)

	for !hasPrefix(blk.Hash, prefix) {
		blk.Nonce++
		blk.Hash = miner.hasher.Hash(serialize(blk.Nonce))
	}
}
