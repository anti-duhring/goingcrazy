package router

import (
	"github.com/anti-duhring/goingcrazy/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	handler.InitializeHandler()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/pessoas", handler.CreatePersonHandler)
		v1.GET("/pessoas", handler.GetPersonHandler)
		v1.DELETE("/pessoas", handler.DeletePersonHandler)
	}
}
