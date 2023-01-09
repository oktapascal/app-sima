package utils

import "golang.org/x/crypto/bcrypt"

// HashString generates a hashed version of the given string using bcrypt.
// The hashed value can be used to securely store sensitive information, such as passwords.
func HashString(value string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(value), 10)

	return string(password), err
}

// ValidatePassword compares a plaintext password to a hashed password using bcrypt.
// It returns true if the plaintext password is valid, false otherwise.
func ValidatePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
