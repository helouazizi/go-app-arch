package repository

import (
	"fmt"
	"user-service/internal/model"
)

type InMemoryUserRepository struct {
	//Mutx   sync.Mutex
	users map[int]*model.User
	//NextId int
}

func NewInMemoryUserRepo() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		make(map[int]*model.User),
	}
}

type UserRepository interface {
	CreateUser(name, email string) (*model.User, error)
	GetUserByID(id int) (*model.User, error)
	UpdateUser(id int, name string) (*model.User, error)
	ListUsers() (*map[int]*model.User, error)
}

func (repo *InMemoryUserRepository) CreateUser(name, email string) (*model.User, error) {
	user := &model.User{
		ID:    int(len(repo.users) + 1),
		Name:  name,
		Email: email,
	}
	repo.users[int(user.ID)] = user
	return user, nil
}

func (repo *InMemoryUserRepository) GetUserByID(id int) (*model.User, error) {
	/*for _, user := range repo.users {
		if user.ID == id {
			return user, nil
		}
	}*/
	user := repo.users[id]
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (repo *InMemoryUserRepository) UpdateUser(id int, name string) (*model.User, error) {
	/*var user *model.User
	for _, usr := range repo.users {
		if usr.ID == id {
			user = usr
		}
	}*/
	user := repo.users[id]
	user.Name = name
	return user, nil
}

func (repo *InMemoryUserRepository) ListUsers() (*map[int]*model.User, error) {
	return &repo.users, nil
}
