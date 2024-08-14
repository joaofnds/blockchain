package blockchain

import (
	"strings"

	"github.com/joaofnds/blockchain/block"
)

type Blockchain struct {
	Blocks []*block.Block
}

func Init() *Blockchain {
	blockchain := &Blockchain{}

	genesisBlock := block.NewBlock("Genesis Block", strings.Repeat("0", 64))
	genesisBlock.Mine(blockchain.Difficulty())

	blockchain.Blocks = []*block.Block{genesisBlock}

	return blockchain
}

func (blockchain *Blockchain) Len() int {
	return len(blockchain.Blocks)
}

func (blockchain *Blockchain) LastBlock() *block.Block {
	if blockchain.Len() == 0 {
		return nil
	}

	return blockchain.Blocks[blockchain.Len()-1]
}

func (blockchain *Blockchain) Difficulty() int {
	return blockchain.Len()/10 + 2
}

func (blockchain *Blockchain) AddBlock(data string) {
	prevBlock := blockchain.LastBlock()
	newBlock := block.NewBlock(data, prevBlock.Hash)
	newBlock.Mine(blockchain.Difficulty())
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

func (blockchain *Blockchain) String() string {
	var str strings.Builder

	for _, block := range blockchain.Blocks {
		str.WriteString(block.String() + "\n")
	}

	return str.String()
}
