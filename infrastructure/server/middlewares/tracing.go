package server

import (
	"github.com/gin-gonic/gin"
	"github.com/opensource-cloud/sycorax/core"
)

func TracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetHeader("X-Trace-Id")
		if traceId == "" {
			traceId = core.NewUUID()
		}

		c.Header("X-Trace-Id", traceId)
		c.Set("traceId", traceId)
	}
}