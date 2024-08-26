package mine

import (
	"strconv"

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
	buf := preNonceBuffer(blk)
	lenBeforeNonce := buf.Len()

	prefix := hashPrefix(difficulty)

	for !hasPrefix(blk.Hash, prefix) {
		blk.Nonce++

		buf.Truncate(lenBeforeNonce)
		buf.WriteString(strconv.FormatUint(blk.Nonce, 10))

		blk.Hash = miner.hasher.Hash(buf.Bytes())
	}
}
