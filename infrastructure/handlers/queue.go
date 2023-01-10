package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/we-are-uranus/sycorax/domain/dtos"
	"github.com/we-are-uranus/sycorax/domain/entities"
	"github.com/we-are-uranus/sycorax/infrastructure/errors"
	"net/http"
)

// PostCreateQueue its a handler for POST /queues router
func PostCreateQueue(c *gin.Context) {
	var dto dtos.CreateQueueDTO

	err := c.BindJSON(&dto)
	if err != nil {
		body := errors.NewInvalidSchemaError(err)
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
