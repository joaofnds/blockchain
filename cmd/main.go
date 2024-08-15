package main

import (
	"github.com/joaofnds/blockchain/blockchain"
	"github.com/joaofnds/blockchain/clock"
	"github.com/joaofnds/blockchain/hash"
	"github.com/joaofnds/blockchain/mine"
)

func main() {
	time := clock.NewTimeClock()
	miner := mine.New(hash.NewSHA256())

	chain := blockchain.New(time, miner)
	chain.AddGenesisBlock()
	chain.AddBlock("Hello, World!")
	chain.AddBlock("Hello, Blockchain!")
	chain.AddBlock("Hello, Go!")

	println(chain.String())
}
