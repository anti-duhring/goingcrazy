package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

func sendError(c *gin.Context, code int, message string) {
	c.Header("Content-Type", "application/json")
	c.JSON(code, gin.H{"error": message, "status": code})
}

func sendSuccess(c *gin.Context, statusCode int, op string, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(statusCode, gin.H{
		"message": fmt.Sprintf("operation from handler: %s sucessfull", op),
		"data":    data,
	})
}

func sendWithoutJSON(c *gin.Context, statusCode int) {
	c.Header("Content-Type", "application/json")
	c.JSON(statusCode, nil)
}

type PersonResponse struct {
	ID         uuid.UUID      `json:"id"`
	Apelido    string         `json:"apelido"`
	Nome       string         `json:"nome"`
	Nascimento datatypes.Date `json:"nascimento"`
	Stack      datatypes.JSON `json:"stack"`
}
