package usecases

import (
	"github.com/opensource-cloud/sycorax/internal/domain/dtos"
	"github.com/opensource-cloud/sycorax/internal/domain/entities"
	"github.com/opensource-cloud/sycorax/internal/domain/repositories"
)

// QueueUseCases
// The main repository is repository.JsonQueueRepository
type QueueUseCases struct {
	repository repositories.QueueRepository
}

// NewQueueUseCases Returns a pointer of QueueUseCases struct
func NewQueueUseCases(repository repositories.QueueRepository) *QueueUseCases {
	return &QueueUseCases{
		repository: repository,
	}
}

// UpsertQueue receives a queue as DTO and build entities.Queue struct and saves in the database
func (uc *QueueUseCases) UpsertQueue(refID string, dto dtos.UpsertQueueDTO) (*entities.Queue, error) {
	queue, err := entities.NewQueue(refID, dto)
	if err != nil {
		return nil, err
	}

	err = uc.repository.Upsert(queue)
	if err != nil {
		return nil, err
	}

	return queue, nil
}

// DeleteQueueByRefId Find and Delete a queue by refId, if not found will return an error
func (uc *QueueUseCases) DeleteQueueByRefId(refId string) error {
	queue, err := uc.repository.FindOneById(refId)
	if err != nil {
		return err
	}
	return uc.repository.Delete(queue)
}

// FindOneByRefId Returns a entities.Queue or error
func (uc *QueueUseCases) FindOneByRefId(refId string) (*entities.Queue, error) {
	return uc.repository.FindOneById(refId)
}

// FindAllPaginated Returns a paginated list of entities.Queue, if it doesn't exist equal or greater than 1 will return an empty list of objects
func (uc *QueueUseCases) FindAllPaginated() (*dtos.Pagination, error) {
	queues, err := uc.repository.FindAll()
	if err != nil {
		return nil, err
	}

	items := make([]interface{}, len(queues))
	for i, q := range queues {
		items[i] = q
	}

	pagination := dtos.NewPagination(items)

	return pagination, nil
}
