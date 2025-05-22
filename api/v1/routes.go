package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/otakenz/kova/internal/app/task"
)

func Routes(taskService *task.TaskService) http.Handler {
	r := chi.NewRouter()
	registerTaskRoutes(r, taskService)
	return r
}

func registerTaskRoutes(r chi.Router, taskService *task.TaskService) {
	taskHandler := NewTaskHandler(taskService)
	// Group task-related endpoints under "/tasks".
	// POST "/"  -> Create a new task
	// GET  "/"  -> List all tasks
	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", taskHandler.CreateTask)
		r.Get("/", taskHandler.ListTasks)
		r.Get("/{id}", taskHandler.GetTask)
		r.Put("/{id}", taskHandler.UpdateTask)
		r.Delete("/{id}", taskHandler.DeleteTask)
		r.Post("/{id}/start", taskHandler.StartTask)
	})
}
