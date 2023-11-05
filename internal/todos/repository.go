package todos

import (
	"github.com/JulianPedro/go-todo-clean-architecture/internal/model"
	"github.com/google/uuid"
)

type TodoRepository interface {
	Create(todo *model.Todo) error
	Update(id uuid.UUID, todo *model.Todo) error
	Delete(id uuid.UUID) error
	GetAll() ([]*model.Todo, error)
}
