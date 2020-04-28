package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	errorToken = errors.New("Error generating token")
)

func GenerateToken(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", errorToken
	}

	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
