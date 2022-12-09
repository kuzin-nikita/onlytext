package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	FullName string
	Email    string
	//Posts
}
