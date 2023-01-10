package middlewares

import "github.com/gin-gonic/gin"

func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
	}
}
