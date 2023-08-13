package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllUsersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": "",
		"error":   nil,
	})
}
