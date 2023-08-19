package handler

import (
	"github.com/gin-gonic/gin"
)

func CreatePersonHandler(c *gin.Context) {
	request := CreatePersonRequest{}

	c.BindJSON(&request)

	logger.Infof("Request received: %+v", request)
}
