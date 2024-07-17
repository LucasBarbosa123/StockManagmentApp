package hashing_utilities

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	hashedStr := hash.Sum(nil)

	return hex.EncodeToString(hashedStr)
}
