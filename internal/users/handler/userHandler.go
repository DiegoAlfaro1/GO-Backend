package handler

import (
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