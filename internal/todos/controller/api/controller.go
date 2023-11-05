package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JulianPedro/go-todo-clean-architecture/internal/model"
	"github.com/JulianPedro/go-todo-clean-architecture/internal/todos"
	"github.com/google/uuid"
)

type TodosController struct {
	todosService todos.Service
}

func NewController(service todos.Service) todos.Controller {
	return &TodosController{todosService: service}
}

func (t *TodosController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todo := &model.Todo{}
	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "invalid request payload"}`))
		return
	}
	err := t.todosService.Create(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func (t *TodosController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todo := &model.Todo{}
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "id param is required"}`))
		return
	}
	idUUID, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "invalid id param"}`))
		return
	}
	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "invalid request payload"}`))
		return
	}
	err = t.todosService.Update(idUUID, todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error updating todo"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

func (t *TodosController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "id param is required"}`))
		return
	}
	idUUID, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "invalid id param"}`))
		return
	}
	err = t.todosService.Delete(idUUID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error deleting todo"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "todo deleted successfully"}`))
}

func (t *TodosController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todos, err := t.todosService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error getting todos"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}
