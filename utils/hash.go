package utils

import "golang.org/x/crypto/bcrypt"

func HashString(value string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(value), 10)

	return string(password), err
}

func ValidatePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
