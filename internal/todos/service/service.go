package service

import (
	"errors"

	"github.com/JulianPedro/go-todo-clean-architecture/internal/model"
	"github.com/JulianPedro/go-todo-clean-architecture/internal/todos"
	"github.com/google/uuid"
)

type Service struct {
	repository todos.TodoRepository
}

func NewService(repository todos.TodoRepository) todos.Service {
	return &Service{repository: repository}
}

func (s *Service) Create(todo *model.Todo) error {
	todo.ID = uuid.New()
	if todo.Name == "" {
		return errors.New("todo name cannot be empty")
	}
	return s.repository.Create(todo)
}

func (s *Service) Update(id uuid.UUID, todo *model.Todo) error {
	return s.repository.Update(id, todo)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repository.Delete(id)
}

func (s *Service) GetAll() ([]*model.Todo, error) {
	return s.repository.GetAll()
}
