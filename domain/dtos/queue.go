package domain

import "encoding/json"

type QueueConfig struct {
	MaxCountOfRetry          json.Number `json:"max_count_of_retry" binding:"required"`
	RetryDelayInMilliseconds json.Number `json:"retry_delay_in_milliseconds" binding:"required"`
}

type CreateQueueDTO struct {
	Name   string       `json:"name" binding:"required"`
	Config *QueueConfig `json:"yaml" binding:"required"`
}
