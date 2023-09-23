package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, code int, message string) {
	c.Header("Content-Type", "application/json")
	c.JSON(code, gin.H{"error": message, "status": code})
}

func sendSuccess(c *gin.Context, op string, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("operation from handler: %s sucessfull", op),
		"data":    data,
	})
}
