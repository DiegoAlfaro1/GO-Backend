package handler

import (
	"net/http"

	"github.com/DiegoAlfaro1/gin-terraform/internal/config"
	"github.com/DiegoAlfaro1/gin-terraform/internal/users/service"
	"github.com/DiegoAlfaro1/gin-terraform/internal/util"
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
		users.POST("/testEndpoint", h.TestFunction)
	}
}

func (h *UserHandler) TestFunction(c *gin.Context) {

	user, err := h.cognitoClient.GetUserFromEmail("john.doe@example.com")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error getting user from cognito"})
		return
	}

	userAtributes := util.ExtractAttributes(user)

	c.JSON(http.StatusOK, gin.H{"Email": userAtributes["email"], "name":userAtributes["name"], "birthdate": userAtributes["birthdate"], "ID": userAtributes["custom:custom_id"]  })
}