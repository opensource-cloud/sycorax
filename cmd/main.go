package main

import (
	"github.com/opensource-cloud/sycorax/internal/api"
	"github.com/opensource-cloud/sycorax/internal/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	log.Println("------------------------------ [Sycorax] ------------------------------")

	log.Print("Loading config")
	c := config.NewConfig()

	go c.LoadYamlFiles()

	server := api.NewServer(c)

	err := server.Start()
	if err != nil {
		panic(err)
	}

	log.Println("------------------------------ [Sycorax] ------------------------------")

}
