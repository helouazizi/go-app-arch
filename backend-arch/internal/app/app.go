package app

import (
	"database/sql"
	"log"
	"user-service/internal/handlers"
	"user-service/internal/repository"
	"user-service/internal/service"
	"user-service/pkg/config"
	"user-service/pkg/logger"

	_ "github.com/mattn/go-sqlite3"
)

// AppHandlers holds all application handlers
type Application struct {
	DB          *sql.DB
	UserHandler *handlers.UserHandler
	//PostHandler *api.PostHandler
	//AuthHandler *api.AuthHandler
}

func NewApp(config *config.Configuration) *Application {
	// Connect to SQLite3
	db, err := sql.Open("sqlite3", config.DB_PATH)
	if err != nil {
		logger.LogWithDetails(err)
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// in this case lets create our db tables
	// Create tables if not exists
	//createTables(db)

	// Repositories
	userRepo := repository.NewInMemoryUserRepo()
	//postRepo := repository.NewInMemoryPostRepository()

	// Services
	userService := service.NewUserService(userRepo)
	//postService := service.NewPostService(postRepo)
	//authService := service.NewAuthService(userRepo) // Auth may depend on users

	// Handlers
	userHandler := handlers.NewUserHandler(userService)
	//postHandler := handlers.NewPostHandler(postService)
	//authHandler := handlers.NewAuthHandler(authService)

	return &Application{
		DB:          db,
		UserHandler: userHandler,
		//PostHandler: postHandler,
		//AuthHandler: authHandler,
	}
}
