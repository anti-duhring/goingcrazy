package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()

	initializeRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
