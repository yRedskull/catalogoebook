package hash_app

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashString(input string) string {
	hash := sha256.Sum256([]byte(input)) // Calcula o hash SHA-256
	return hex.EncodeToString(hash[:])  // Converte para string hexadecimal
}
