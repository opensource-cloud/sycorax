package main

import (
	"github.com/opensource-cloud/sycorax/infrastructure/config"
	"github.com/opensource-cloud/sycorax/infrastructure/server"
)

func main() {
	app := config.GetApp()

	app.LoadResourcesFolder()

	server.StartHttpServer(app)
}
