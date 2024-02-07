package domain

type User struct {
	ID   int
	Name string
	Age  int
}

type UserRepository interface {
	FindByID(id string) (*User, error)
	Save(user *User) error
}