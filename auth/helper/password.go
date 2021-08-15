package helper

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error encrypting password: %s", err.Error())
		return "", errors.New("internal server error")
	}

	return string(hash), nil
}

func ComparePassword(password, hash []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		return errors.New("password does not match")
	}

	return nil
}
