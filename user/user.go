package user

import (
	"github.com/google/uuid"
)

// User represents a signed up user to the service
type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
}

