package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

type Hasher interface {
	Hash([]byte) string
}

var _ Hasher = (*SHA256Hasher)(nil)

type SHA256Hasher struct{}

func NewSHA256() *SHA256Hasher {
	return &SHA256Hasher{}
}

func (hasher *SHA256Hasher) Hash(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}
