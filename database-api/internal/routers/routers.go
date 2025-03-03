package routers

import (
	"net/http"

	"user-service/internal/app"
)

func SetupRoutes(h *app.Application) {
	// this route for user
	http.HandleFunc("POST /users", h.UserHandler.CreateUser)
	// http.HandleFunc("/users/update", h.UserHandler.UpdateUser)
	// http.HandleFunc("/users", h.UserHandler.ListUsers)

	// this routs for posts
	// http.HandleFunc("/users", h.UserHandler.CreateUser)
	// http.HandleFunc("/users/update", h.UserHandler.UpdateUser)
	// http.HandleFunc("/users", h.UserHandler.ListUsers)

	// this routs for authofication
	// http.HandleFunc("/users", h.UserHandler.CreateUser)
	// http.HandleFunc("/users/update", h.UserHandler.UpdateUser)
	// http.HandleFunc("/users", h.UserHandler.ListUsers)
}
