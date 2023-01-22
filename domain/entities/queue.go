package domain

import (
	"github.com/we-are-uranus/sycorax/core"
	dtos "github.com/we-are-uranus/sycorax/domain/dtos"
	"time"
)

type (
	QueuePublishConfig struct {
		Driver           string `json:"driver"`
		MaxSizeOfMessage int    `json:"max_size_of_message"`
	}
	QueueDeliveryConfig struct {
		Driver     string `json:"driver"`
		RawMessage bool   `json:"raw_message"`
	}
	QueueConfig struct {
		Type     string               `json:"type"`
		Publish  *QueuePublishConfig  `json:"publish"`
		Delivery *QueueDeliveryConfig `json:"delivery"`
	}
	Queue struct {
		Id        string       `json:"id"`
		Name      string       `json:"name"`
		Config    *QueueConfig `json:"config"`
		CreatedAt time.Time    `json:"created_at"`
		UpdatedAt time.Time    `json:"updated_at"`
	}
)

func NewQueue(dto dtos.CreateQueueDTO) (*Queue, error) {
	queue := &Queue{
		Id:   core.NewUUID(),
		Name: dto.Name,
		Config: &QueueConfig{
			Type: "FIFO",
			Publish: &QueuePublishConfig{
				Driver:           "MEMORY",
				MaxSizeOfMessage: 1, // MB
			},
			Delivery: &QueueDeliveryConfig{
				Driver:     "HTTP_POLLING",
				RawMessage: true,
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := queue.isValid()

	if err != nil {
		return nil, err
	}

	return queue, nil
}

func (q *Queue) isValid() error {
	return nil
}
