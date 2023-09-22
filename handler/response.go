package handler

import "github.com/gin-gonic/gin"

func sendError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message, "status": code})
}
