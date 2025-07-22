package handler

import (
	"net/http"

	"github.com/DiegoAlfaro1/gin-terraform/internal/users/model"
	"github.com/DiegoAlfaro1/gin-terraform/internal/users/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{
	userService service.UserService
}

func NewUSerHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine){
	users := r.Group("/users")
	{
		users.GET("/", h.GetAll )
		users.GET("/:id", h.GetOne)
		users.POST("/", h.Create)
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

func (h *UserHandler) Create(c *gin.Context){
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	createdUser, err := h.userService.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Message": "User created successfully", "User": createdUser}) 
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