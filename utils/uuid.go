package utils

import "github.com/google/uuid"

// GenerateUUID generates and returns a universally unique identifier (UUID) as a string.
func GenerateUUID() string {
	_uuid := uuid.New()

	return _uuid.String()
}
