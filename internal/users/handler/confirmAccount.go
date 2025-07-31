package handler

import (
	"net/http"

	"github.com/DiegoAlfaro1/gin-terraform/internal/config"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) ConfirmAccount(c *gin.Context) {
	var input config.UserConfirmation

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Confirm account with Cognito
	err := h.cognitoClient.ConfirmAccount(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to confirm account"})
		return
	}

	//TODO: create cognito method to get a user info, then use the method create from cognito instead of create from email

	user, err := h.userService.CreateUserFromEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user in local DB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account confirmed successfully",
		"user":    user,
	})
}
