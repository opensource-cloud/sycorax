package v1

import "C"
import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/opensource-cloud/sycorax/internal/config"
	"github.com/opensource-cloud/sycorax/internal/domain/dtos"
	"github.com/opensource-cloud/sycorax/internal/domain/entities"
	"net/http"
)

func RegisterQueueV1Routes(rg *gin.RouterGroup, c *config.Config) {
	queues := rg.Group("/queues")

	providers := c.Providers()
	useCases := providers.UseCases

	queues.GET("/", func(context *gin.Context) {
		findManyQueues(context, useCases)
	})

	queues.GET("/:ref_id", func(context *gin.Context) {
		findQueueByRefId(context, useCases)
	})

	queues.PUT("/:ref_id", func(context *gin.Context) {
		upsertQueue(context, useCases)
	})

	queues.DELETE("/:ref_id", func(context *gin.Context) {
		deleteQueueByRefId(context, useCases)
	})
}

// PUT /queues/:ref_id
// BODY dtos.UpsertQueueDTO
// HTTP 201  *entities.Queue
// HTTP 400 v1.SycoraxError
// HTTP 422 v1.SycoraxError
// HTTP 500 v1.SycoraxError
func upsertQueue(c *gin.Context, uc *config.UseCases) {
	refID := c.Param("ref_id")
	if refID == "" {
		body := NewInvalidSchemaError(errors.New("invalid reference id, please check /docs"))
		c.IndentedJSON(http.StatusBadRequest, body)
		return
	}

	var dto dtos.UpsertQueueDTO

	err := c.BindJSON(&dto)
	if err != nil {
		body := NewInvalidSchemaError(err)
		body.ParseErrorsToFields(c, dto)
		c.IndentedJSON(http.StatusBadRequest, body)
		return
	}

	queue, err := uc.Queues.UpsertQueue(refID, dto)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, NewUnprocessableEntity(err))
		return
	}

	c.IndentedJSON(http.StatusCreated, queue)
}

// GET /queues
// QUERY ?pageSize=<int> || Optional (default=25)
// QUERY ?page=<int> || Optional (use page or cursor)
// QUERY ?cursor<string> || Optional (use cursor or page)
// HTTP 200 dtos.Pagination<[]*entities.Queue>
// HTTP 500 v1.SycoraxError
func findManyQueues(c *gin.Context, uc *config.UseCases) {
	pagination, err := uc.Queues.FindAllPaginated()
	if err != nil {
		body := NewInternalServerError(err)
		c.IndentedJSON(http.StatusInternalServerError, body)
		return
	}
	c.IndentedJSON(http.StatusOK, pagination)
}

// GET /queues/:ref_id
// HTTP 200 *entities.Queue
// HTTP 400 v1.SycoraxError
// HTTP 404 v1.SycoraxError
// HTTP 500 v1.SycoraxError
func findQueueByRefId(c *gin.Context, uc *config.UseCases) {
	refID := c.Param("ref_id")
	if refID == "" {
		body := NewInvalidSchemaError(errors.New("invalid reference id, please check /docs"))
		c.IndentedJSON(http.StatusBadRequest, body)
		return
	}

	queue, err := uc.Queues.FindOneByRefId(refID)
	if err != nil {
		switch err {
		case entities.QueueDoesNotExists:
			body := NewSycoraxError("Queue does not exists", "NOT_FOUND", err)
			body.AddField(&FieldError{
				Field:   "ref_id",
				Value:   refID,
				Message: "Wrong reference identifier",
			})
			c.IndentedJSON(http.StatusNotFound, body)
			return
		default:
			c.IndentedJSON(http.StatusInternalServerError, NewInternalServerError(err))
			return
		}
	}

	c.IndentedJSON(http.StatusOK, queue)
}

// DELETE /queues/:ref_id
// HTTP 204 void
// HTTP 400 v1.SycoraxError
// HTTP 404 v1.SycoraxError
// HTTP 422 v1.SycoraxError
// HTTP 500 v1.SycoraxError
func deleteQueueByRefId(c *gin.Context, uc *config.UseCases) {
	refID := c.Param("ref_id")
	if refID == "" {
		body := NewInvalidSchemaError(errors.New("invalid reference id, please check /docs"))
		c.IndentedJSON(http.StatusBadRequest, body)
		return
	}

	err := uc.Queues.DeleteQueueByRefId(refID)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, NewUnprocessableEntity(err))
		return
	}

	c.IndentedJSON(http.StatusNoContent, "")
}
