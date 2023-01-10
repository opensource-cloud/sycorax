package dtos

type CreateQueueDTO struct {
	Name string `json:"name" binding:"required"`
}
