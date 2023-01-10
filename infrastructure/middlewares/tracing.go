package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/we-are-uranus/sycorax/core"
)

func TracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetHeader("X-Trace-Id")
		if traceId == "" {
			traceId = core.NewUUID()
		}
		c.Header("X-Trace-Id", traceId)
	}
}
