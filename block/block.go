package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Data      []byte
	Timestamp time.Time
	PrevHash  string
	Hash      string
	Nonce     uint64
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

func (block *Block) CalculateHash() {
	hash := sha256.New()
	hash.Write(block.Serialize())
	block.Hash = hex.EncodeToString(hash.Sum(nil))
}

func (block *Block) Mine(difficulty int) {
	target := strings.Repeat("0", difficulty)

	for !strings.HasPrefix(block.Hash, target) {
		block.Nonce++
		block.CalculateHash()
	}
}

func NewBlock(data string, prevHash string) *Block {
	block := &Block{
		Data:      []byte(data),
		Timestamp: time.Now(),
		PrevHash:  prevHash,
		Hash:      "",
		Nonce:     0,
	}

	block.CalculateHash()

	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", strings.Repeat("0", 64))
}
