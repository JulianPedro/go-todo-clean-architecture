package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JulianPedro/go-todo-clean-architecture/internal/todos"
	"github.com/go-chi/chi"
)

var (
	chiDispatcher *chi.Mux = chi.NewRouter()
)

type ChiRouter struct{}

func NewChiRouter() todos.Router {
	return &ChiRouter{}
}

func (c *ChiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (c *ChiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (c *ChiRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Put(uri, f)
}

func (c *ChiRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Delete(uri, f)
}

func (c *ChiRouter) SERVE(port int) {
	log.Println("Server running on port", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), chiDispatcher)
}
