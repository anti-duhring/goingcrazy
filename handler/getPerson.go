package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPersonHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get person",
	})
}
