package domain

import "github.com/google/uuid"

type Users []User

type User struct {
	ID       uuid.UUID
	Name     string
	Username string
	Password string
}
