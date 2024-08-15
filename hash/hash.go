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

func (_ *SHA256Hasher) Hash(data []byte) string {
	h := sha256.New()
	h.Write(data)
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}
