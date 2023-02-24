package server

import (
	"github.com/gin-gonic/gin"
	dtos "github.com/opensource-cloud/sycorax/domain/dtos"
	entities "github.com/opensource-cloud/sycorax/domain/entities"
	errors "github.com/opensource-cloud/sycorax/infrastructure/server/errors"
	"log"
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
		traceId, _ := c.Get("traceId")
		log.Print("Bad Request Trace: ", traceId, ", Error: ", err.Error())
		c.IndentedJSON(http.StatusBadRequest, errors.NewInvalidSchemaError(err))
		return
	}

	c.IndentedJSON(http.StatusCreated, queue)
}
