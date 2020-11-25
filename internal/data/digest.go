package data

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// GenerateDigest uses SHA256 to return a digest
func GenerateDigest(secret string, message string) string {

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
