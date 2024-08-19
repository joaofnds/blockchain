package main

import (
	"flag"
	"io"
	"os"

	"github.com/joaofnds/blockchain/blockchain"
	"github.com/joaofnds/blockchain/blockchain/storage"
	"github.com/joaofnds/blockchain/clock"
	"github.com/joaofnds/blockchain/hash"
	"github.com/joaofnds/blockchain/mine"
)

func main() {
	inputFile := flag.String("file", "blockchain.json", "path to the blockchain file")
	help := flag.Bool("help", false, "print the help message")
	flag.CommandLine.Parse(os.Args[2:])

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if len(os.Args) < 2 {
		println("missing command")
		os.Exit(1)
	}

	assertFileExists(*inputFile)

	time := clock.NewTimeClock()
	miner := mine.New(hash.NewSHA256())
	chain := blockchain.New(time, miner)

	jsonStorage := storage.NewJSON(*inputFile)
	jsonStorage.LoadBlocks(chain)

	command := os.Args[1]
	switch command {
	case "print":
		println(chain.String())
	case "add":
		b, readErr := io.ReadAll(os.Stdin)
		if readErr != nil {
			println("error reading stdin: " + readErr.Error())
			os.Exit(1)
		}

		chain.AddBlock(b)

		saveErr := jsonStorage.Save(chain)
		if saveErr != nil {
			println("error saving blockchain: " + saveErr.Error())
			os.Exit(1)
		}
	default:
		println("invalid command")
		os.Exit(1)
	}
}

func assertFileExists(file string) {
	_, statsErr := os.Stat(file)
	if os.IsNotExist(statsErr) {
		println("file does not exist: " + file)
		os.Exit(1)
	}
}
