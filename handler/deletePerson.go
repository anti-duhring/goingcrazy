package handler

import (
	"fmt"
	"net/http"

	"github.com/anti-duhring/goingcrazy/schema"
	"github.com/gin-gonic/gin"
)

func DeletePersonHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	person := schema.Person{}

	if err := db.Model(&person).Where("id = ?", id).First(&person).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("person with id '%s' not found", id))
		return
	}

	if err := db.Delete(&person).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting person with id '%s'", id))
		return
	}

	sendSuccess(c, http.StatusOK, "delete-person", person)

}
