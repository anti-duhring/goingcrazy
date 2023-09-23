package handler

import (
	"fmt"
	"net/http"

	"github.com/anti-duhring/goingcrazy/schema"
	"github.com/gin-gonic/gin"
)

func GetPersonHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	person := schema.Person{}

	if err := db.Model(&schema.Person{}).Where("id = ?", id).First(&person).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("person with id '%s' not found", id))
		return
	}

	sendSuccess(c, "get-person", person)
}
