package config

import (
	"log"
	"os"
)

func (app *App) LoadResourcesFolder() {
	log.Printf("Running")

	log.Printf(app.Vars.Resources.Path)

	files, err := os.ReadDir("./resources")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		log.Println(file.Name(), file.IsDir())
	}
}
