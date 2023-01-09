package utils

import "github.com/google/uuid"

func GenerateUUID() string {
	_uuid := uuid.New()

	return _uuid.String()
}
