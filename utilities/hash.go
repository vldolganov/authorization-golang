package utilities

import (
	"log"

	"github.com/alexedwards/argon2id"
)

func HashPassword(password string) string {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	return hash
}

func CheckPasswordHash(password, hash string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		log.Fatal(err)
	}
	return match
}
