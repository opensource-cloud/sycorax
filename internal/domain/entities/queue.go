package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opensource-cloud/sycorax/internal/core"
	"github.com/opensource-cloud/sycorax/internal/domain/dtos"
	"time"
)

const (
	FIFO         = "FIFO"
	MemoryDriver = "MEMORY_DRIVER"
)

var (
	QueueDoesNotExists = errors.New("queue does not exists")
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

func NewQueue(refID string, dto dtos.UpsertQueueDTO) (*Queue, error) {
	queue := &Queue{
		Id:    core.NewUUID(),
		RefID: refID,
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
		return errors.New("queue.name must not contain more than 255 characters")
	}

	if q.Config.Type != FIFO {
		return errors.New("queue.config type not allowed, check the documentation")
	}

	if q.Config.Driver != MemoryDriver {
		return errors.New("queue.config.driver not allowed, check the documentation")
	}

	if q.Config.MaxSizeOfMessage > 5 {
		return errors.New("queue.config.max_size_of_message must not be greater than 5 megabytes")
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

func (q *Queue) Update(dto dtos.UpsertQueueDTO) error {
	q.Name = dto.Name
	q.Config = &QueueConfig{
		Type:             dto.Config.Type,
		Driver:           dto.Config.Driver,
		MaxSizeOfMessage: dto.Config.MaxSizeOfMessage,
	}

	err := q.isValid()
	if err != nil {
		return err
	}

	return nil
}
