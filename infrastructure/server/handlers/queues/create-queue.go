package server

import (
	"github.com/gin-gonic/gin"
	dtos "github.com/we-are-uranus/sycorax/domain/dtos"
	entities "github.com/we-are-uranus/sycorax/domain/entities"
	errors "github.com/we-are-uranus/sycorax/infrastructure/server/errors"
	"net/http"
)

// PostCreateQueue its a handler for POST /queues router
func PostCreateQueue(c *gin.Context) {
	var dto dtos.CreateQueueDTO

	err := c.BindJSON(&dto)
	if err != nil {
		body := errors.NewInvalidSchemaError(err)
		body.ParseErrorsToFields(c, dto)
		c.IndentedJSON(http.StatusBadRequest, body)
		return
	}

	queue, err := entities.NewQueue(dto)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.NewInvalidSchemaError(err))
		return
	}

	c.IndentedJSON(http.StatusCreated, queue)
}
