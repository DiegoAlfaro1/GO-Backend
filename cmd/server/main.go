package main

import (
	"os"

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

    // Step 3: Initialize Cognito client
    appClientID := os.Getenv("COGNITO_CLIENT_ID")
    cognitoClient := config.NewCognitoClient(appClientID)

    // Step 4: Set up Gin
    r := gin.Default()

    // Step 5: Initialize layers
    userRepo := repository.NewUserRepository()
    userService := service.NewUserService(userRepo)
    userHandler := handler.NewUSerHandler(userService, cognitoClient)

    // Step 6: Register routes
    userHandler.RegisterRoutes(r)

    // Step 7: Run the server
    r.Run(":8080")
}