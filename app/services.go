package app

import (
	"encoding/json"
	"errors"
	"fmt"
	dtos "github.com/opensource-cloud/sycorax/domain/dtos"
	domain "github.com/opensource-cloud/sycorax/domain/entities"
	"log"
)

type Services struct {
	db   *JsonDB
	vars *Vars
}

// NewServices returns an instance of Service that contains a bunch of functions
func NewServices(db *JsonDB) *Services {
	return &Services{
		db: db,
	}
}

// CreateQueue Does all the necessary logic to create a queue and save into the db (used by http and file resources)
func (s *Services) CreateQueue(dto dtos.CreateQueueDTO) (*domain.Queue, error) {
	log.Printf("Creating queue %s", dto.Name)

	db := s.db

	refIdKey := db.MakeQueueKey(dto.RefID)

	if db.Has(refIdKey) {
		log.Printf("Queue %s already exists in db.", refIdKey)
		err := db.Delete(refIdKey)
		if err != nil {
			log.Printf("Error deleting %s on database", refIdKey)
			return nil, err
		}
	}

	queue, err := domain.NewQueue(dto)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error creating %s, err: %s", dto.Name, err))
	}

	err = db.Set(refIdKey, queue.ToJSON())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error setting the queue %s on json db, err %s", queue.Name, err))
	}

	log.Printf("Queue %s - %s created successfully", queue.Name, queue.RefID)

	return queue, nil
}

// FindManyQueues Get all queues on DB by index.
func (s Services) FindManyQueues() ([]*domain.Queue, error) {
	log.Print("Finding all queues in db")

	db := s.db

	queuesAsString, err := db.FindManyByIndex("queues")
	if err != nil {
		log.Print("Error getting all queues")
		return nil, err
	}

	var queues []*domain.Queue

	log.Print("Mapping all queues to struct")

	for _, queueAsString := range queuesAsString {
		var incomingQueue *domain.Queue
		err := json.Unmarshal([]byte(queueAsString), &incomingQueue)
		if err != nil {
			log.Printf("Error parsing queue %s", queuesAsString)
			return nil, err
		}
		queues = append(queues, incomingQueue)
	}

	log.Print("All queues loaded, returning")

	return queues, nil
}
