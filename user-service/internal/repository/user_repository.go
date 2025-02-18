package repository

import "user-service/internal/model"

type UserRepository interface {
	CreateUser(name, email string) (*model.User, error)
	GetUserByID(id int64) (*model.User, error)
	ListUsers() ([]*model.User, error)
}
