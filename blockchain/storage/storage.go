package storage

import "github.com/joaofnds/blockchain/blockchain"

type Storage interface {
	Save(chain *blockchain.Blockchain) error
	LoadBlocks(chain *blockchain.Blockchain) error
}
