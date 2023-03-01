package app

import (
	dto "github.com/opensource-cloud/sycorax/domain/dtos"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"path"
)

type (
	yamlQueue struct {
		RefID            string `yaml:"ref_id"`
		Name             string `yaml:"name"`
		QueueType        string `yaml:"type"`
		Driver           string `yaml:"driver"`
		MaxSizeOfMessage int    `yaml:"max_size_of_message"`
	}
	yamlQueueFile struct {
		Queues yaml.Node `yaml:"queues"`
	}
)

func (app *App) LoadYamlFiles() *App {
	log.Println("--------------- [Yaml - Resources] ---------------")

	files, err := os.ReadDir(app.Paths.Yaml)
	if err != nil {
		log.Fatal(err)
	}

	var queues map[string]*yamlQueue = make(map[string]*yamlQueue)

	for _, file := range files {
		fileName := file.Name()
		filePath := path.Join(app.Paths.Yaml, fileName)
		switch fileName {
		case "queues.yaml":
			loadYamlQueueFile(fileName, filePath, queues)
		default:
			log.Printf("No condition for %s in switch statement", fileName)
		}
	}

	countOfQueues := len(queues)
	log.Printf("Count of queues %d", countOfQueues)

	if countOfQueues > 0 {
		for _, queue := range queues {
			createQueueFromYaml(app, queue)
		}
	}

	log.Println("--------------- [Yaml - Resources] ---------------")

	return app
}

// Queues Yaml - Loading and Creating
func loadYamlQueueFile(fileName string, filePath string, queues map[string]*yamlQueue) {
	log.Printf("Reading and parsing %s", fileName)

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file %s, path: %s, error: %s", fileName, filePath, err)
	}

	schema := yamlQueueFile{}
	err = yaml.Unmarshal(file, &schema)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = schema.Queues.Decode(&queues)
	if err != nil {
		log.Fatalf("Error decoding queues from yaml file, err: %v", err)
	}
}
func createQueueFromYaml(app *App, yamlQueue *yamlQueue) {
	queueDTO := dto.CreateQueueDTO{
		Name:  yamlQueue.Name,
		RefID: yamlQueue.RefID,
		Config: &dto.CreateQueueConfigDTO{
			Driver:           yamlQueue.Driver,
			Type:             yamlQueue.QueueType,
			MaxSizeOfMessage: yamlQueue.MaxSizeOfMessage,
		},
	}

	_, err := app.Services.CreateQueue(queueDTO)
	if err != nil {
		log.Fatalf("Error creating queue, detail: %s", err)
	}
}
