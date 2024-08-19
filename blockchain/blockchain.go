package blockchain

import (
	"strings"

	"github.com/joaofnds/blockchain/block"
	"github.com/joaofnds/blockchain/clock"
	"github.com/joaofnds/blockchain/mine"
)

type Blockchain struct {
	Clock  clock.Clock
	miner  *mine.Miner
	Blocks []*block.Block
}

func New(clock clock.Clock, miner *mine.Miner) *Blockchain {
	return &Blockchain{
		Clock:  clock,
		miner:  miner,
		Blocks: []*block.Block{},
	}
}

func (blockchain *Blockchain) AddGenesisBlock() {
	if blockchain.Len() > 0 {
		return
	}

	genesisBlock := block.NewBlock([]byte("Genesis Block"), blockchain.Clock.Now(), strings.Repeat("0", 64))
	blockchain.miner.Mine(genesisBlock, blockchain.Difficulty())

	blockchain.Blocks = append(blockchain.Blocks, genesisBlock)
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
	if prevBlock == nil {
		return
	}

	newBlock := block.NewBlock([]byte(data), blockchain.Clock.Now(), prevBlock.Hash)
	blockchain.miner.Mine(newBlock, blockchain.Difficulty())
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

func (blockchain *Blockchain) String() string {
	var str strings.Builder

	for _, block := range blockchain.Blocks {
		str.WriteString(block.String() + "\n")
	}

	return str.String()
}
