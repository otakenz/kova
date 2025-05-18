package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/otakenz/kova/internal/core/task"
)

func Routes(taskStore *task.Store) http.Handler {
	r := chi.NewRouter()
	registerPingRoute(r)
	registerTaskRoutes(r, taskStore)
	return r
}

func registerPingRoute(r chi.Router) {
	// Register a GET route on "/ping"
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"pong"}`))
	})
}

func registerTaskRoutes(r chi.Router, taskStore *task.Store) {
	taskHandler := NewTaskHandler(taskStore)
	// Group task-related endpoints under "/tasks".
	// POST "/"  -> Create a new task
	// GET  "/"  -> List all tasks
	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", taskHandler.CreateTask)
		r.Get("/", taskHandler.ListTasks)
		r.Get("/{id}", taskHandler.GetTask)
		r.Put("/{id}", taskHandler.UpdateTask)
		r.Delete("/{id}", taskHandler.DeleteTask)
	})
}
