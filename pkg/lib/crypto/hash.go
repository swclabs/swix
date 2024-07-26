package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashOf SHA-256
func HashOf(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
