package domain

import "errors"

type User struct {
	ID   string
	Name string
}

type UserRepository interface {
	FindByID(id string) (*User, error)
}

var (
	ErrUserNotFound = errors.New("user not found")
)