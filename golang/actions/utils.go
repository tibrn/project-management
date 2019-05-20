package actions

import (
	"encoding/base64"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(unique string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(unique), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(hash)
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
