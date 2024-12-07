package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassowrd(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Failed to hash password: %w", err) 
	}
	return string(hashedPassword), err
}

func CheckPassowrd(password string, hashedPasword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPasword), []byte(password))
}