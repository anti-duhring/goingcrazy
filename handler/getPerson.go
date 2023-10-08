package handler

import (
	"net/http"

	"github.com/anti-duhring/goingcrazy/schema"
	"github.com/gin-gonic/gin"
)

func GetPersonHandler(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "pathParam").Error())
		return
	}

	_, personExistsErr := cache.GetPerson(c, id)

	if personExistsErr == nil {
		sendWithoutJSON(c, http.StatusOK)
		return
	}

	person := schema.Person{}

	if err := db.Model(&schema.Person{}).Select("id").Where("id = ?", id).First(&person).Error; err != nil {
		sendWithoutJSON(c, http.StatusNotFound)
		return
	}

	// personResponse := PersonResponse{
	// 	ID:         person.ID,
	// 	Apelido:    person.Apelido,
	// 	Nome:       person.Nome,
	// 	Nascimento: person.Nascimento,
	// 	Stack:      person.Stack,
	// }

	sendWithoutJSON(c, http.StatusOK)
}
