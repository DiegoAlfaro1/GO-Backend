package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) GetOne(c *gin.Context) {
	idParam := c.Param("id")
	user, err := h.userService.GetOneUser(idParam)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}