package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

//GetHash returns
func GetHash(dataString string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(dataString))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
