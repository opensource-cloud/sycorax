package main

import (
	"github.com/opensource-cloud/sycorax/app"
	"log"
)

func main() {
	log.Print("Creating app from scratch")
	application := app.GetApp()
	log.Print("Application created")

	log.Print("Loading all yaml files inside resources folder")
	application.LoadYamlFiles()

	log.Print("Starting application routes server")
	application.StartHttpServer()
}
