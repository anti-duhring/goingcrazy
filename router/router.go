package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()

	initializeRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("API_PORT")
	go r.Run(":" + port)
}
