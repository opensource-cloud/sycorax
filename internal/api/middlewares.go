package api

import (
	"github.com/gin-gonic/gin"
	"github.com/opensource-cloud/sycorax/internal/core"
	log "github.com/sirupsen/logrus"
	"time"
)

func RegisterMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

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
