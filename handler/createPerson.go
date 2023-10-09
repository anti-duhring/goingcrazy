package handler

import (
	"errors"
	"net/http"

	"github.com/anti-duhring/goingcrazy/config"
	"github.com/anti-duhring/goingcrazy/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreatePersonHandler(c *gin.Context) {
	person := schema.Person{}

	if err := c.BindJSON(&person); err != nil {
		logger.Errorf("error binding json: %v", err.Error())

		sendWithoutJSON(c, http.StatusBadRequest)
		return
	}

	if err := person.Validate(c); err != nil {
		logger.Errorf("error validating request: %v", err.Error())

		sendWithoutJSON(c, http.StatusUnprocessableEntity)
		return
	}

	_, personExistsErr := cache.GetPerson(c, person.Apelido)

	if personExistsErr == nil {
		sendWithoutJSON(c, http.StatusUnprocessableEntity)
		return
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
