package main

import (
	"github.com/joaofnds/blockchain/blockchain"
	"github.com/joaofnds/blockchain/blockchain/storage"
	"github.com/joaofnds/blockchain/clock"
	"github.com/joaofnds/blockchain/hash"
	"github.com/joaofnds/blockchain/mine"
)

// TODO:
//   - add `IsValid` to `Blockchain`
//   - `challenge` interface (takes a block and returns a bool) to be used in `ProofOfWork`

func main() {
	time := clock.NewTimeClock()
	miner := mine.New(hash.NewSHA256())
	jsonStorage := storage.NewJSON("./blockchain.json")

	chain := blockchain.New(time, miner)
	// chain.AddGenesisBlock()
	// chain.AddBlock("Hello, World!")
	// chain.AddBlock("Hello, Blockchain!")
	// chain.AddBlock("Hello, Go!")
	// jsonStorage.Save(chain)

	jsonStorage.LoadBlocks(chain)

	println(chain.String())
}
