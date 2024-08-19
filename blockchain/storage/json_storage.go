package storage

import (
	"encoding/json"
	"os"

	"github.com/joaofnds/blockchain/blockchain"
)

type JSONStorage struct {
	path string
}

var _ Storage = (*JSONStorage)(nil)

func NewJSON(path string) JSONStorage {
	return JSONStorage{path: path}
}

func (jsonStorage JSONStorage) Save(chain *blockchain.Blockchain) error {
	bytes, marshalErr := json.Marshal(chain.Blocks)
	if marshalErr != nil {
		return marshalErr
	}

	return os.WriteFile(jsonStorage.path, bytes, 0644)
}

func (jsonStorage JSONStorage) LoadBlocks(chain *blockchain.Blockchain) error {
	file, openErr := os.Open(jsonStorage.path)
	if openErr != nil {
		return openErr
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	decodeErr := decoder.Decode(&chain.Blocks)
	if decodeErr != nil {
		return decodeErr
	}

	return nil
}
