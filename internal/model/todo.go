package model

import "github.com/google/uuid"

type Todo struct {
	ID   uuid.UUID `json:"id" validate:"omitempty,uuid"`
	Name string    `json:"name" validate:"required"`
	Done bool      `json:"done"`
}
