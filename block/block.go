package block

import (
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Data      []byte    `json:"data"`
	Timestamp time.Time `json:"timestamp"`
	PrevHash  string    `json:"prev_hash"`
	Hash      string    `json:"hash"`
	Nonce     uint64    `json:"nonce"`
}

func NewBlock(data []byte, timestamp time.Time, prevHash string) Block {
	return Block{
		Data:      data,
		Timestamp: timestamp,
		PrevHash:  prevHash,
		Hash:      "",
		Nonce:     0,
	}
}

func (block *Block) String() string {
	var buf strings.Builder

	buf.WriteString("Block {\n")
	buf.WriteString("\tData: " + string(block.Data) + "\n")
	buf.WriteString("\tTimestamp: " + block.Timestamp.Format(time.RFC3339) + "\n")
	buf.WriteString("\tPrevHash: " + block.PrevHash + "\n")
	buf.WriteString("\tHash: " + block.Hash + "\n")
	buf.WriteString("\tNonce: " + strconv.FormatUint(block.Nonce, 10) + "\n")
	buf.WriteString("}")

	return buf.String()
}
