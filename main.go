package main

import (
	"github.com/opensource-cloud/sycorax/app"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	log.Println("------------------------------ [Sycorax] ------------------------------")

	log.Print("Creating app from scratch")
	application := app.GetApp()
	log.Print("Application created")

	log.Print("Loading all yaml files inside resources folder")
	application.LoadYamlFiles()

	log.Println("------------------------------ [Sycorax] ------------------------------")

	application.StartHttpServer()
}
