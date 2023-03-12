package config

import (
	"github.com/opensource-cloud/sycorax/internal/domain/dtos"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"path"
)

type (
	yamlQueueFile struct {
		Queues yaml.Node `yaml:"queues"`
	}
)

func (c *Config) LoadYamlFiles() {
	log.Println("--------------- [Yaml - Resources] ---------------")

	files, err := os.ReadDir(c.Paths.Yaml)
	if err != nil {
		log.Fatal(err)
	}

	var queues map[string]*dtos.YamlQueueDTO = make(map[string]*dtos.YamlQueueDTO)

	for _, file := range files {
		fileName := file.Name()
		filePath := path.Join(c.Paths.Yaml, fileName)
		switch fileName {
		case "queues.yaml":
			loadYamlQueueFile(fileName, filePath, queues)
		default:
			log.Printf("No condition for %s in switch statement", fileName)
		}
	}

	countOfQueues := len(queues)
	log.Printf("Count of queues %d", countOfQueues)

	providers := c.Providers()
	useCases := providers.UseCases

	if countOfQueues > 0 {
		for _, queue := range queues {
			createQueueFromYaml(queue, useCases)
		}
	}

	log.Println("--------------- [Yaml - Resources] ---------------")
}

func loadYamlQueueFile(fileName string, filePath string, queues map[string]*dtos.YamlQueueDTO) {
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
func createQueueFromYaml(dto *dtos.YamlQueueDTO, uc *UseCases) {
	_, err := uc.Queues.UpsertQueue(dto.RefID, dtos.UpsertQueueDTO{
		Name: dto.Name,
		Config: &dtos.QueueConfigDTO{
			Driver:           dto.Driver,
			Type:             dto.QueueType,
			MaxSizeOfMessage: dto.MaxSizeOfMessage,
		},
	})
	if err != nil {
		log.Fatalf("Error creating queue, detail: %s", err)
	}
}
