package internal

import (
	"crypto/sha256"
)

func CryptoSha256(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	return h.Sum([]byte("nice"))
}
