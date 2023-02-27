package app

import (
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
