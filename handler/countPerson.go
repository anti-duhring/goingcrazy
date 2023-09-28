package handler

import (
	"net/http"

	"github.com/anti-duhring/goingcrazy/schema"
	"github.com/gin-gonic/gin"
)

func CountPersonHandler(c *gin.Context) {

	var count int64

	db.Model(&schema.Person{}).Count(&count)

	sendSuccess(c, http.StatusOK, "count-person", count)
}
