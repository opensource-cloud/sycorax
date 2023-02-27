package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opensource-cloud/sycorax/core"
	dtos "github.com/opensource-cloud/sycorax/domain/dtos"
	"time"
)

const (
	FIFO         = "FIFO"
	MemoryDriver = "MEMORY_DRIVER"
)

type (
	QueueConfig struct {
		Type             string `json:"type"`
		Driver           string `json:"driver"`
		MaxSizeOfMessage int    `json:"max_size_of_message"`
	}
	Queue struct {
		Id        string       `json:"id"`
		RefID     string       `json:"ref_id"`
		Name      string       `json:"name"`
		Config    *QueueConfig `json:"config"`
		CreatedAt time.Time    `json:"created_at"`
		UpdatedAt time.Time    `json:"updated_at"`
	}
)

func NewQueue(dto dtos.CreateQueueDTO) (*Queue, error) {
	queue := &Queue{
		Id:    core.NewUUID(),
		RefID: dto.RefID,
		Name:  dto.Name,
		Config: &QueueConfig{
			Type:             dto.Config.Type,
			Driver:           dto.Config.Driver,
			MaxSizeOfMessage: dto.Config.MaxSizeOfMessage,
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
	if len(q.Name) >= 255 {
		return errors.New("queue name must not contain more than 255 characters")
	}

	if q.Config.Type != FIFO {
		return errors.New("queue app type not allowed, check the documentation")
	}

	if q.Config.Driver != MemoryDriver {
		return errors.New("queue app driver not allowed, check the documentation")
	}

	if q.Config.MaxSizeOfMessage > 5 {
		return errors.New("queue max size of message must not be greater than 5 megabytes")
	}

	return nil
}

func (q *Queue) ToJSON() string {
	queueAsJson, err := json.Marshal(q)
	if err != nil {
		panic(fmt.Sprintf("Error parsing queue %s, err: %s", q.Name, err))
	}
	return string(queueAsJson)
}
