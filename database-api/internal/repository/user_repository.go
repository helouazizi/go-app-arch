package repository

import (
	"database/sql"

	"user-service/internal/models"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	// GetUserByID(id int) (model.User, error)
	// UpdateUser(user model.User) (model.User, error)
	// ListUsers() ([]model.User, error)
}

type SQLiteUserRepository struct {
	db *sql.DB
}

func NewSQLiteUserRepository(db *sql.DB) *SQLiteUserRepository {
	return &SQLiteUserRepository{db: db}
}

func (r *SQLiteUserRepository) CreateUser(user models.User) (models.User, error) {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	r.db.Exec(query, user.Name, user.Email)
	return user, nil
}

// Implement other methods...
