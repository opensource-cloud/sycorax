package app

import (
	"github.com/gin-gonic/gin"
	"github.com/opensource-cloud/sycorax/core"
	dtos "github.com/opensource-cloud/sycorax/domain/dtos"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// StartHttpServer starts the routes server using app config
func (app *App) StartHttpServer() {
	r := gin.New()

	// Formatting request log
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return ""
	}))

	// Setting the recovery mode
	r.Use(gin.Recovery())

	// Setting all middlewares
	r.Use(headersMiddleware())
	r.Use(tracingMiddleware())

	// Loading all routes
	// TODO: Abstract this into another file?

	// Queues
	r.POST("/queues", postCreateQueue)
	r.GET("/queues", findManyQueues)

	// Set all configs for development mode
	if app.IsDEV {
		log.Print("Setting configs to routes development mode")
		err := r.SetTrustedProxies([]string{"0.0.0.0"})
		if err != nil {
			log.Fatalf("Error setting localhost as trusted proxy, err: %s", err)
		}
	}

	// Set all configs for production mode
	if app.IsPROD {
		log.Print("Setting configs to routes production mode")
		gin.SetMode(gin.ReleaseMode)
	}

	// TODO: Rethink this set manually
	port := ":6789"

	log.Printf("Setting up routes server on port %s", port)
	err := r.Run(port)
	if err != nil {
		log.Fatalf("Error running the routes server: %s", err)
	}

	log.Printf("Server is running on port %s", port)
}

// TODO: Abstract all this middlewares functions into another file
func headersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
	}
}
func tracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		traceId := c.GetHeader("X-Trace-Id")

		if traceId == "" {
			traceId = core.NewUUID()
		}

		log.WithFields(log.Fields{
			"trace_id":       traceId,
			"request_url":    c.Request.URL.Path,
			"request_method": c.Request.Method,
			"user_agent":     c.Request.UserAgent(),
		}).Info("A new request received")

		c.Header("X-Trace-Id", traceId)
		c.Set("traceId", traceId)

		c.Next()

		latency := time.Since(start)
		log.WithFields(log.Fields{
			"trace_id":             traceId,
			"response_latency":     latency.String(),
			"response_status_code": c.Writer.Status(),
		}).Info("Request ended")
	}
}

// TODO: Abstract all this handlers functions into another file

// Queues

// POST /queues
func postCreateQueue(c *gin.Context) {
	var dto dtos.CreateQueueDTO

	err := c.BindJSON(&dto)
	if err != nil {
		body := NewInvalidSchemaError(err)
		body.ParseErrorsToFields(c, dto)
		c.IndentedJSON(http.StatusBadRequest, body)
		return
	}

	queue, err := app.Services.CreateQueue(dto)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, NewInvalidSchemaError(err))
		return
	}

	c.IndentedJSON(http.StatusCreated, queue)
}

// GET /queues
func findManyQueues(c *gin.Context) {
	queues, err := app.Services.FindManyQueues()
	if err != nil {
		body := NewInternalServerError(err)
		c.IndentedJSON(http.StatusInternalServerError, body)
		return
	}

	items := make([]interface{}, len(queues))

	for i, q := range queues {
		items[i] = q
	}

	pagination := NewPagination(items)

	c.IndentedJSON(http.StatusOK, pagination)

	return
}
