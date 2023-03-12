package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/opensource-cloud/sycorax/internal/api/v1"
	"github.com/opensource-cloud/sycorax/internal/config"
	log "github.com/sirupsen/logrus"
)

type (
	Server struct {
		config *config.Config
	}
)

// NewServer Return a new ApiServer struct
func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

// Start Init the HTTP Server
func (s *Server) Start() error {
	r := gin.New()

	// Formatting request log
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return ""
	}))

	// Setting the recovery mode
	r.Use(gin.Recovery())
	r.Use(RegisterMiddlewares())

	// Set all configs for development mode
	if s.config.IsDEV {
		log.Print("Setting configs to routes development mode")
		err := r.SetTrustedProxies([]string{"0.0.0.0"})
		if err != nil {
			log.Printf("Error setting localhost as trusted proxy, err: %s", err)
			return err
		}
	}

	// Set all configs for production mode
	if s.config.IsPROD {
		log.Print("Setting configs to routes production mode")
		gin.SetMode(gin.ReleaseMode)
	}

	v1Group := r.Group("/v1")
	v1.RegisterV1RoutesGroups(v1Group, s.config)

	port := ":6789"

	log.Printf("Setting up routes server on port %s", port)
	err := r.Run(port)
	if err != nil {
		log.Printf("Error running the routes server: %s", err)
		return err
	}

	log.Printf("Server is running on port %s", port)

	return nil
}
