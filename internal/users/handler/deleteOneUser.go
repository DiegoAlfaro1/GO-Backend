package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) DeleteOne(c *gin.Context) {
	idParam := c.Param("id")

	err := h.userService.DeleteOne(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}