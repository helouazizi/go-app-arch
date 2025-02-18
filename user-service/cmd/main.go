package main

import (
	"fmt"
	"log"
	"net/http"
	"user-service/internal/app"
	"user-service/internal/routers"
	"user-service/pkg/config"
)

func main() {
	configurations := config.LoadConfig()
	application := app.NewApp(configurations)
	routers.SetupRoutes(application)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
