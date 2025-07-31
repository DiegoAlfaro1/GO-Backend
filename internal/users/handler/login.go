package handler

import (
	"net/http"

	"github.com/DiegoAlfaro1/gin-terraform/internal/config"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) Login(c *gin.Context) {
	var input config.UserLogin

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Authenticate with Cognito
	token, err := h.cognitoClient.SignIn(&input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}