package encryption

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func HashPassword2(password string) (string, error) {
	h := hmac.New(sha256.New, []byte(password))
	_, err := h.Write([]byte(password))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
