package helper

import (
	"github.com/alexedwards/argon2id"
)

func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	return hash, err
}

func CheckPassword(hashedPassword, password string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, hashedPassword)
	return err == nil && match
}