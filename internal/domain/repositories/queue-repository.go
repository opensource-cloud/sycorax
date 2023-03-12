package repositories

import "github.com/opensource-cloud/sycorax/internal/domain/entities"

type QueueRepository interface {
	FindOneById(queueId string) (*entities.Queue, error)
	FindAll() ([]*entities.Queue, error)
	Upsert(queue *entities.Queue) error
	Delete(queue *entities.Queue) error
}
