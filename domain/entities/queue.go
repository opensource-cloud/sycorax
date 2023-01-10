package entities

import (
	"github.com/we-are-uranus/sycorax/core"
	"github.com/we-are-uranus/sycorax/domain/dtos"
	"time"
)

type Queue struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewQueue(dto dtos.CreateQueueDTO) (*Queue, error) {
	queue := &Queue{
		Id:        core.NewUUID(),
		Name:      dto.Name,
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
