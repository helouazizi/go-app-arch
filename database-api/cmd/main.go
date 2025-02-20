package main

import (
	"fmt"
	"log"
	"net/http"

	"user-service/internal/app"
	"user-service/internal/routers"
	"user-service/pkg/config"
	"user-service/pkg/logger"
)

func main() {
	logger, err := logger.Create_Logger()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Close()
	configurations := config.LoadConfig()
	application := app.NewApp(configurations)
	routers.SetupRoutes(application)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
