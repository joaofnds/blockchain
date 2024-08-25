package block

import (
	"bytes"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Data      []byte    `json:"Data"`
	Timestamp time.Time `json:"Timestamp"`
	PrevHash  string    `json:"PrevHash"`
	Hash      string    `json:"Hash"`
	Nonce     uint64    `json:"Nonce"`
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

func (block *Block) Serialize() []byte {
	var buf bytes.Buffer

	buf.Write(block.Data)
	buf.WriteString(block.PrevHash)
	buf.WriteString(block.Timestamp.String())
	buf.WriteString(strconv.FormatUint(block.Nonce, 10))

	return buf.Bytes()
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
