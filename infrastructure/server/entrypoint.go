package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/we-are-uranus/sycorax/infrastructure/config"
	middlewares "github.com/we-are-uranus/sycorax/infrastructure/server/middlewares"
	routes "github.com/we-are-uranus/sycorax/infrastructure/server/routes"
)

func StartHttpServer() {
	envs := config.GetEnvVars()

	r := gin.Default()

	r.Use(middlewares.HeadersMiddleware())
	r.Use(middlewares.TracingMiddleware())
	r.Use(middlewares.LoggerMiddleware())

	// Load all server routes using a pointer reference
	routes.LoadRestRoutes(r)

	// Set all trusted proxies for development environment
	if envs.IsDev {
		r.SetTrustedProxies([]string{"0.0.0.0"})
	}

	err := r.Run(fmt.Sprintf(":%s", envs.HTTP.Port))

	if err != nil {
		panic("Was not possible to start the service")
	}
}
