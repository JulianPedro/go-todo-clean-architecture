package main

import (
	"log"

	"github.com/JulianPedro/go-todo-clean-architecture/internal/todos"
	"github.com/JulianPedro/go-todo-clean-architecture/internal/todos/controller/api"
	"github.com/JulianPedro/go-todo-clean-architecture/internal/todos/repository"
	"github.com/JulianPedro/go-todo-clean-architecture/internal/todos/service"
)

var (
	todoRepository todos.TodoRepository = repository.NewMemoryRepository()
	todoService    todos.Service        = service.NewService(todoRepository)
	todoController todos.Controller     = api.NewController(todoService)
	todoRouter     todos.Router         = api.NewChiRouter()
)

func main() {
	log.Println("Starting application \"go-todo-clean-architecture\"...")
	todoRouter.GET("/todos", todoController.GetAll)
	todoRouter.POST("/todos", todoController.Create)
	todoRouter.PUT("/todos", todoController.Update)
	todoRouter.DELETE("/todos", todoController.Delete)
	todoRouter.SERVE(8080)
}
