package main

import (
	"fmt"
	"log"
	"net/http"
	"user-service/internal/handlers"
	"user-service/internal/repository"
	"user-service/internal/service"
)

// AppHandlers holds all application handlers
type AppHandlers struct {
	UserHandler *handlers.UserHandler
	//PostHandler *api.PostHandler
	//AuthHandler *api.AuthHandler
}

// setupRoutes initializes all HTTP routes
func setupRoutes(h *AppHandlers) {
	http.HandleFunc("/users", h.UserHandler.CreateUser)
	http.HandleFunc("/users/update", h.UserHandler.UpdateUser)
	http.HandleFunc("/users", h.UserHandler.ListUsers)
}

// initializeApp initializes repositories, services, and handlers for all features
func initializeApp() *AppHandlers {
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

	return &AppHandlers{
		UserHandler: userHandler,
		//PostHandler: postHandler,
		//AuthHandler: authHandler,
	}
}

func main() {
	handlers := initializeApp()
	setupRoutes(handlers)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
