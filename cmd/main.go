package main

import (
	"github.com/joaofnds/blockchain/blockchain"
)

const difficulty = 2

func main() {
	chain := blockchain.Init()
	chain.AddBlock("Hello, World!")
	chain.AddBlock("Hello, Blockchain!")
	chain.AddBlock("Hello, Go!")
	println(chain.String())
}
