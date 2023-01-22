package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		// after request
		latency := time.Since(t).Milliseconds()
		status := c.Writer.Status()

		log.Printf("path %s | latency %d | status %d", c.Request.URL.String(), latency, status)
	}
}
