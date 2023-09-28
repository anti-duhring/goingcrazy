package handler

import (
	"fmt"
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

	person := schema.Person{}

	if err := db.Model(&schema.Person{}).Where("id = ?", id).First(&person).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("person with id '%s' not found", id))
		return
	}

	personResponse := PersonResponse{
		ID:         person.ID,
		Apelido:    person.Apelido,
		Nome:       person.Nome,
		Nascimento: person.Nascimento,
		Stack:      person.Stack,
	}

	sendSucessWithoutMessage(c, http.StatusOK, personResponse)
}
