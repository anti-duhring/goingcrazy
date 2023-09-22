package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anti-duhring/goingcrazy/schema"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

func CreatePersonHandler(c *gin.Context) {
	request := CreatePersonRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())

		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	stackJSON, err := json.Marshal(request.Stack)

	if err != nil {
		logger.Errorf("error marshalling stack: %v", err.Error())

		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	person := schema.Person{
		Apelido:    request.Apelido,
		Nome:       request.Nome,
		Nascimento: datatypes.Date(request.Nascimento),
		Stack:      stackJSON,
	}

	if err := db.Create(&person).Error; err != nil {
		logger.Errorf("error creating perso: %v", err.Error())

		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}
}
