package app

import (
	"database/sql"
	"log"

	"user-service/internal/database"
	"user-service/internal/handlers"
	"user-service/internal/repository"
	"user-service/internal/services"
	"user-service/pkg/config"
	"user-service/pkg/logger"

	_ "github.com/mattn/go-sqlite3"
)

type Application struct {
	DB          *sql.DB
	UserHandler *handlers.UserHandler
}

func NewApp(config *config.Configuration) *Application {
	// Connect to SQLite3
	db, err := sql.Open("sqlite3", config.DB_PATH)
	if err != nil {
		logger.LogWithDetails(err)
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// Run database migrations to craete tables
	database.Migrate(db)

	// Initialize repository & service
	userRepo := repository.NewSQLiteUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	return &Application{
		DB:          db,
		UserHandler: userHandler,
	}
}
