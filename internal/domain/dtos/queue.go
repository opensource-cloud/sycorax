package dtos

type (
	QueueConfigDTO struct {
		Type             string `json:"type"`
		Driver           string `json:"driver"`
		MaxSizeOfMessage int    `json:"max_size_of_message"`
	}
	UpsertQueueDTO struct {
		Name   string          `json:"name"  binding:"required"`
		Config *QueueConfigDTO `json:"config" binding:"required"`
	}
	YamlQueueDTO struct {
		RefID            string `yaml:"ref_id"`
		Name             string `yaml:"name"`
		QueueType        string `yaml:"type"`
		Driver           string `yaml:"driver"`
		MaxSizeOfMessage int    `yaml:"max_size_of_message"`
	}
)
