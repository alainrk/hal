package utils

import (
	"github.com/google/uuid"
)

// GenerateID generates a unique ID
func GenerateID() string {
	return uuid.New().String()
}
