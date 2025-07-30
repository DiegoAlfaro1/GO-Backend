package main

import (
	"github.com/DiegoAlfaro1/gin-terraform/internal/config"
	"github.com/DiegoAlfaro1/gin-terraform/internal/users/handler"
	"github.com/DiegoAlfaro1/gin-terraform/internal/users/repository"
	"github.com/DiegoAlfaro1/gin-terraform/internal/users/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Step 1: Load .env file
	config.LoadEnv() // loads environment variables from .env file

	// Step 2: Initialize Database connection
	config.InitDatabaseConnection() 

	// Step 3: Set up Gin
	r := gin.Default()

	// Step 4: Initialize layers
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUSerHandler(userService)

	// Step 5: Register routes
	userHandler.RegisterRoutes(r)

	// Step 6: Run the server
	r.Run(":8080")
}
