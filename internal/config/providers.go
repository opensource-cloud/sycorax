package config

import (
	"github.com/opensource-cloud/sycorax/internal/database/repositories"
	"github.com/opensource-cloud/sycorax/internal/domain/usecases"
)

type (
	Repositories struct {
		QueueRepository *jsonrepositories.JsonQueueRepository
	}
	UseCases struct {
		Queues *usecases.QueueUseCases
	}
	Providers struct {
		UseCases     *UseCases
		Repositories *Repositories
	}
)

func (c *Config) Providers() *Providers {
	return &Providers{
		UseCases:     c.useCases(),
		Repositories: c.repositories(),
	}
}

// Private pointer receivers methods

func (c *Config) useCases() *UseCases {
	repositories := c.repositories()

	queueRepository := repositories.QueueRepository

	return &UseCases{
		Queues: usecases.NewQueueUseCases(queueRepository),
	}
}

func (c *Config) repositories() *Repositories {
	return &Repositories{
		QueueRepository: jsonrepositories.NewJsonQueueRepository(c.DB),
	}
}
