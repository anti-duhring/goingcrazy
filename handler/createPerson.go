package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePersonHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Post person",
	})
}
