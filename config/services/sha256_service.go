package services

import (
	"crypto/sha256"
	"fmt"
)

// Function for Crypto password
func SHAR256Encoder(s string) string {
	str := sha256.Sum256([]byte(s))

	return fmt.Sprintf("%x", str)
}