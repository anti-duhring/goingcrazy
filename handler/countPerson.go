package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CountPersonHandler(c *gin.Context) {

	var count int64

	db.Table("people").Count(&count)

	c.Header("Content-Type", "text/plain")
	c.String(http.StatusOK, "%d", count)
}
