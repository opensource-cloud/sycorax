package main

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/we-are-uranus/sycorax/infrastructure/handlers"
	"github.com/we-are-uranus/sycorax/infrastructure/middlewares"
)

func main() {
	r := gin.Default()

	r.Use(middlewares.HeadersMiddleware())
	r.Use(middlewares.TracingMiddleware())

	r.POST("/queues", handlers.PostCreateQueue)

	err := r.Run()

	if err != nil {
		panic("Was not possible to start the service")
	}
}
