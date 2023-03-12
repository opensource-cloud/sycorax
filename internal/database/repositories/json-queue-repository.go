package jsonrepositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opensource-cloud/sycorax/internal/database/engine"
	"github.com/opensource-cloud/sycorax/internal/domain/entities"
	"github.com/opensource-cloud/sycorax/internal/domain/repositories"
)

type JsonQueueRepository struct {
	repositories.QueueRepository

	db *jsondb.JsonDB
}

func NewJsonQueueRepository(db *jsondb.JsonDB) *JsonQueueRepository {
	return &JsonQueueRepository{
		db: db,
	}
}

func (r *JsonQueueRepository) makeQueueRefIdKey(refId string) string {
	return fmt.Sprintf("queues:%s", refId)
}

func (r *JsonQueueRepository) FindOneById(refId string) (*entities.Queue, error) {
	key := r.makeQueueRefIdKey(refId)

	queueAsString := r.db.Get(key)
	if queueAsString == "" {
		return nil, entities.QueueDoesNotExists
	}

	var queue *entities.Queue
	err := json.Unmarshal([]byte(queueAsString), &queue)
	if err != nil {
		return nil, err
	}

	return queue, nil
}

func (r *JsonQueueRepository) FindAll() ([]*entities.Queue, error) {
	listOfQueuesAsString, err := r.db.FindManyByIndex("queues")
	if err != nil {
		return nil, err
	}

	var queues []*entities.Queue

	for _, queueAsString := range listOfQueuesAsString {
		var incomingQueue *entities.Queue
		err := json.Unmarshal([]byte(queueAsString), &incomingQueue)
		if err != nil {
			return nil, err
		}
		queues = append(queues, incomingQueue)
	}

	return queues, nil
}

func (r *JsonQueueRepository) Upsert(queue *entities.Queue) error {
	key := r.makeQueueRefIdKey(queue.RefID)

	err := r.db.Set(key, queue.ToJSON())
	if err != nil {
		return errors.New(fmt.Sprintf("Error upserting the queue %s on json db, err %s", queue.Name, err))
	}

	return nil
}

func (r *JsonQueueRepository) Delete(queue *entities.Queue) error {
	key := r.makeQueueRefIdKey(queue.RefID)

	err := r.db.Delete(key)
	if err != nil {
		return errors.New(fmt.Sprintf("Error deleting the queue %s on json db, err %s", queue.Name, err))
	}

	return nil
}
