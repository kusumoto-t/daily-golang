package repository

import (
	"testClean/domain"
)

type InMemoryUserRepository struct {
	users map[string]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {

	users := make(map[string]*domain.User)

	// 初期データを追加
	users["user1"] = &domain.User{ID: "user1", Name: "Alice"}
	users["user2"] = &domain.User{ID: "user2", Name: "Bob"}
	return &InMemoryUserRepository{
		users: users,
	}
}

func (repo InMemoryUserRepository) FindByID(id string) (*domain.User, error) {
	user, ok := repo.users[id]

	if !ok {
		return nil, domain.ErrUserNotFound
	}

	return user, nil
}
