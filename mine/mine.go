package mine

import (
	"bytes"

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

func preNonceBuffer(blk *block.Block) bytes.Buffer {
	var buf bytes.Buffer

	buf.Write(blk.Data)
	buf.WriteString(blk.PrevHash)
	buf.WriteString(blk.Timestamp.String())

	return buf
}
