package service

import (
	"errors"
	"user-service/internal/model"
	"user-service/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (us *UserService) CreateUser(name, email string) (*model.User, error) {
	// Business logic can go here (e.g., validation)
	if name == "" || email == "" {
		return nil, errors.New("name and email cannot be empty")
	}

	return us.repo.CreateUser(name, email)
}

func (us *UserService) GetUserByID(id int) (*model.User, error) {
	return us.repo.GetUserByID(id)
}

func (us *UserService) ListUsers() (*map[int]*model.User, error) {
	return us.repo.ListUsers()
}

func (us *UserService) UpdateUser(id int, name string) (*model.User, error) {
	return us.repo.UpdateUser(id, name)
}
