package handler

import (
	"net/http"

	"github.com/anti-duhring/goingcrazy/schema"
	"github.com/gin-gonic/gin"
)

func SearchPersonHandler(c *gin.Context) {
	searchTerm := c.Query("t")

	if searchTerm == "" {
		sendWithoutJSON(c, http.StatusBadRequest)
		return
	}

	people := []schema.Person{}
	maxPersons := 50

	if err := db.Model(&schema.Person{}).Select("id").Limit(maxPersons).Where(`
		search_index ILIKE CONCAT('%', ?::text, '%')
	`, searchTerm).Find(&people).Error; err != nil {
		sendWithoutJSON(c, http.StatusOK)
		return
	}

	sendWithoutJSON(c, http.StatusOK)
}
