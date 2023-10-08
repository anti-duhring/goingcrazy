package handler

import (
	"net/http"

	"github.com/anti-duhring/goingcrazy/schema"
	"github.com/gin-gonic/gin"
)

func SearchPersonHandler(c *gin.Context) {
	searchTerm := c.Query("t")

	if searchTerm == "" {
		sendError(c, http.StatusBadRequest, "Search term is required")
		return
	}

	people := []schema.Person{}
	maxPersons := 50

	if err := db.Model(&schema.Person{}).Limit(maxPersons).Where(`
		search_index ILIKE CONCAT('%', ?::text, '%')
	`, searchTerm).Find(&people).Error; err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucessWithoutMessage(c, http.StatusOK, people)
}
