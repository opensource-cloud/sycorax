package server

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/we-are-uranus/sycorax/infrastructure/server/handlers/queues"
)

func LoadRestRoutes(r *gin.Engine) {
	r.POST("/queues", handlers.PostCreateQueue)
}
