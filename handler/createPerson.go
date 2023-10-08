package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/anti-duhring/goingcrazy/config"
	"github.com/anti-duhring/goingcrazy/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

func CreatePersonHandler(c *gin.Context) {
	request := CreatePersonRequest{}

	if err := c.BindJSON(&request); err != nil {
		logger.Errorf("error binding json: %v", err.Error())

		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(c); err != nil {
		logger.Errorf("error validating request: %v", err.Error())

		sendError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	_, personExistsErr := cache.GetPerson(c, request.Apelido)

	if personExistsErr == nil {
		sendWithoutJSON(c, http.StatusUnprocessableEntity)
		return
	}

	stackJSON, err := json.Marshal(request.Stack)

	if err != nil {
		logger.Errorf("error marshalling stack: %v", err.Error())

		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	person := schema.Person{
		Apelido:     request.Apelido,
		Nome:        request.Nome,
		Nascimento:  datatypes.Date(request.Nascimento.Time),
		Stack:       stackJSON,
		SearchIndex: fmt.Sprintf("%s %s %s", request.Apelido, request.Nome, stackJSON),
	}
	person.ID = uuid.New()

	if err := worker.Create(c, person); err != nil {
		if errors.Is(err, config.ErrNicknameAlreadyExists) {
			sendWithoutJSON(c, http.StatusUnprocessableEntity)
			return
		}

		logger.Errorf("error creating person: %v", err.Error())

		sendWithoutJSON(c, http.StatusUnprocessableEntity)

		return
	}

	addLocationToHeader(c, person.ID)
	sendWithoutJSON(c, http.StatusCreated)
}

func addLocationToHeader(c *gin.Context, id uuid.UUID) {
	c.Header("Location", "/pessoas/"+id.String())
}
