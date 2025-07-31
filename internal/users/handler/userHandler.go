package handler

import (
	"net/http"

	"github.com/DiegoAlfaro1/gin-terraform/internal/config"
	"github.com/DiegoAlfaro1/gin-terraform/internal/users/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{
	userService service.UserService
	cognitoClient config.CognitoInterface
}

func NewUSerHandler(userService service.UserService,cognitoClient config.CognitoInterface) *UserHandler {
	return &UserHandler{	
		userService: userService,
		cognitoClient: cognitoClient,
	
	}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine){
	users := r.Group("/users")
	{
		users.GET("/", h.GetAll )
		users.POST("/", h.Create)
		users.POST("/confirm", h.ConfirmAccount)
		users.POST("/login", h.Login)
		users.GET("/:id", h.GetOne)
		users.DELETE("/:id", h.DeleteOne)
	}
}

func (h *UserHandler) GetAll(c *gin.Context){
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetOne(c *gin.Context) {
	idParam := c.Param("id")
	user, err := h.userService.GetOneUser(idParam)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

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
		"email": input.Email,
	})
}

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
		"user": user,
	})
}

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
		"token": token,
	})
}

func (h *UserHandler) DeleteOne(c *gin.Context) {
	idParam := c.Param("id")

	err := h.userService.DeleteOne(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}