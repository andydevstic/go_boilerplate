package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashFromPassword(password []byte) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func IsPasswordValid(hash []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)

	if err != nil {
		err = fmt.Errorf("compare password: %w", err)
	}

	return err == nil
}
