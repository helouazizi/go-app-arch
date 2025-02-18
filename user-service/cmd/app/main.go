package main

import (
	"fmt"
	"log"
	"net/http"
	"user-service/internal/api"
	"user-service/internal/repository"
	"user-service/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	// Repository implementation (stub for now)
	repo := &repository.InMemoryUserRepository{}
	userService := service.NewUserService(repo)
	userHandler := api.NewUserHandler(userService)

	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userHandler.ListUsers).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUser).Methods("GET")

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
