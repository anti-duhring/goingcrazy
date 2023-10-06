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
		nome ILIKE CONCAT('%', ?::text, '%') OR
		apelido ILIKE CONCAT('%', ?::text, '%') OR
		(
			stack IS NOT NULL AND
			jsonb_typeof(stack) = 'array' AND
			EXISTS (
				SELECT 1
				FROM jsonb_array_elements(stack) AS element
				WHERE element::text ILIKE CONCAT('%', ?::text, '%')
			)
		)	
	`, searchTerm, searchTerm, searchTerm).Find(&people).Error; err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucessWithoutMessage(c, http.StatusOK, people)
}
