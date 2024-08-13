package main

import "github.com/joaofnds/blockchain/block"

const difficulty = 2

func main() {
	b0 := block.NewGenesisBlock()
	b0.Mine(difficulty)
	println(b0.String())

	b1 := block.NewBlock("Hello, World!", b0.Hash)
	b1.Mine(difficulty)
	println(b1.String())
}
