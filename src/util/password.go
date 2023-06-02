package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// Generate hashed password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Unable to generated hashed password: %w", err)
	}
	return string(hashedPassword), nil
}

// Check if password are matched
func checkPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
