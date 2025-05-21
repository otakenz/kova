package api

import (
	"net/http"

	"github.com/otakenz/kova/api/middleware"
	v1 "github.com/otakenz/kova/api/v1"
	"github.com/otakenz/kova/internal/app"

	"github.com/go-chi/chi/v5"
)

func NewRouter(taskService *app.TaskService) http.Handler {
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.Logging)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	// Mount versioned API
	r.Route("/api", func(r chi.Router) {
		r.Mount("/v1", v1.Routes(taskService))
	})

	return r
}
