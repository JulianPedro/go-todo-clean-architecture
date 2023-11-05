package repository

import (
	"errors"

	"github.com/JulianPedro/go-todo-clean-architecture/internal/model"
	"github.com/JulianPedro/go-todo-clean-architecture/internal/todos"
	"github.com/google/uuid"
)

var (
	todosData []*model.Todo = []*model.Todo{}
)

type MemoryRepository struct{}

func NewMemoryRepository() todos.TodoRepository {
	return &MemoryRepository{}
}

func (m *MemoryRepository) Create(todo *model.Todo) error {
	todosData = append(todosData, todo)
	return nil
}

func (m *MemoryRepository) Update(id uuid.UUID, todo *model.Todo) error {
	for i, t := range todosData {
		if t.ID == id {
			todo.ID = id
			if todo.Name == "" {
				todo.Name = t.Name
			}
			todosData[i] = todo
			return nil
		}
	}
	return errors.New("todo not found")
}

func (m *MemoryRepository) Delete(id uuid.UUID) error {
	for i, t := range todosData {
		if t.ID == id {
			todosData = append(todosData[:i], todosData[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}

func (m *MemoryRepository) GetAll() ([]*model.Todo, error) {
	return todosData, nil
}
