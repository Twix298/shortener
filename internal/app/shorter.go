package app

import (
	"crypto/sha256"
	"encoding/hex"
)

func shortURL(url string) string {
	hash := sha256.New()
	hash.Write([]byte(url))
	return hex.EncodeToString(hash.Sum(nil))
}
