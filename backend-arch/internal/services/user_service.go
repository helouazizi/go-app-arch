package services

import (
	"user-service/internal/models"
	"user-service/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.repo.CreateUser(user)
}

// Implement other methods...
