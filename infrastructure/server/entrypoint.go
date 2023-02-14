package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/opensource-cloud/sycorax/infrastructure/config"
	middlewares "github.com/opensource-cloud/sycorax/infrastructure/server/middlewares"
	routes "github.com/opensource-cloud/sycorax/infrastructure/server/routes"
)

func StartHttpServer(app *config.App) {
	r := gin.Default()

	r.Use(middlewares.HeadersMiddleware())
	r.Use(middlewares.TracingMiddleware())
	r.Use(middlewares.LoggerMiddleware())

	// Load all server routes using a pointer reference
	routes.LoadRestRoutes(r)

	// Set all trusted proxies for development environment
	if app.IsDEV {
		err := r.SetTrustedProxies([]string{"0.0.0.0"})
		if err != nil {
			panic("Error setting localhost as trusted proxy")
		}
	}

	port := fmt.Sprintf(":%s", app.Vars.Http.Port)
	err := r.Run()

	if err != nil {
		panic("Error running the server")
	}

	log.Print(fmt.Printf("Server is running on port %s", port))
}
