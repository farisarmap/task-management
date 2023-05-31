package helper

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	result, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	var hashPassword string
	if err != nil {
		fmt.Println("failed to hash password", err)
		return hashPassword, err
	}

	hashPassword = string(result[:])

	return hashPassword, nil
}
