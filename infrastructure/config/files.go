package config

import (
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type (
	YamlQueue struct {
		Name             string `yaml:"name"`
		Type             string `yaml:"type"`
		Driver           string `yaml:"driver"`
		MaxSizeOfMessage int    `yaml:"max_size_of_message"`
	}
	YamlQueueFile struct {
		Queues yaml.Node `yaml:"queues"`
	}
)

func (app *App) LoadResourcesFolder() *App {
	log.Printf("Running")

	log.Printf(app.Vars.Resources.Path)

	files, err := os.ReadDir(app.Paths.Yaml)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()
		filePath := path.Join(app.Paths.Yaml, fileName)
		file, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("Error reading file %s, path: %s, error: %s", fileName, filePath, err)
		}

		log.Printf("File: \n %s \n", string(file))

		schema := YamlQueueFile{}
		err = yaml.Unmarshal(file, &schema)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		var queues map[string]*YamlQueue
		err = schema.Queues.Decode(&queues)

		for _, queue := range queues {
			log.Printf("Queue %v", queue)
		}
	}

	return app
}
