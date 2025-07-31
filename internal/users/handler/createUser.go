package handler

import (
	"net/http"

	"github.com/DiegoAlfaro1/gin-terraform/internal/config"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) Create(c *gin.Context) {
	var input config.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Register user with Cognito first
	err := h.cognitoClient.SignUp(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register with Cognito"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered with Cognito. Please confirm your account with the code sent to your email.",
		"email":   input.Email,
	})
}