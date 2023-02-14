package main

import (
	"github.com/opensource-cloud/sycorax/infrastructure/config"
	server "github.com/opensource-cloud/sycorax/infrastructure/server"
)

func main() {
	app := config.GetApp()

	app.LoadResourcesFolder()

	server.StartHttpServer(app)
}
