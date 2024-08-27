package mine

import (
	"bytes"
	"strconv"

	"github.com/joaofnds/blockchain/block"
)

type Miner interface {
	Mine(blk *block.Block, difficulty int)
}

func hashPrefix(size int) string {
	prefix := make([]byte, size)

	for i := range prefix {
		prefix[i] = '0'
	}

	return string(prefix)
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

func blockSerializer(blk *block.Block) func(uint64) []byte {
	var buf bytes.Buffer

	buf.Write(blk.Data)
	buf.WriteString(blk.PrevHash)
	buf.WriteString(blk.Timestamp.String())
	lenBeforeNonce := buf.Len()

	return func(nonce uint64) []byte {
		buf.Truncate(lenBeforeNonce)
		buf.WriteString(strconv.FormatUint(nonce, 10))
		return buf.Bytes()
	}
}
