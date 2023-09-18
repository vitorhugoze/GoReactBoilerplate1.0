package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func CreateHash(value string) ([]byte, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return hash, nil
}

func CompareHash(value, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))

	return err == nil
}
