package domain

type (
	CreateQueueConfigDTO struct {
		Type             string `json:"type"`
		Driver           string `json:"driver"`
		MaxSizeOfMessage int    `json:"max_size_of_message"`
	}
	CreateQueueDTO struct {
		Name   string                `json:"name"  binding:"required"`
		Config *CreateQueueConfigDTO `json:"config" binding:"required"`
	}
)
