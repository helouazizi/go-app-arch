package repository

import (
	"fmt"
	"user-service/internal/model"
)

type InMemoryUserRepository struct {
	users []*model.User
}
type UserRepository interface {
	CreateUser(name, email string) (*model.User, error)
	GetUserByID(id int64) (*model.User, error)
	ListUsers() ([]*model.User, error)
}

func (repo *InMemoryUserRepository) CreateUser(name, email string) (*model.User, error) {
	user := &model.User{
		ID:    int64(len(repo.users) + 1),
		Name:  name,
		Email: email,
	}
	repo.users = append(repo.users, user)
	return user, nil
}

func (repo *InMemoryUserRepository) GetUserByID(id int64) (*model.User, error) {
	for _, user := range repo.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (repo *InMemoryUserRepository) ListUsers() ([]*model.User, error) {
	return repo.users, nil
}
