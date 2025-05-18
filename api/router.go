package api

import (
	"net/http"

	"github.com/otakenz/kova/api/middleware"
	v1 "github.com/otakenz/kova/api/v1"
	"github.com/otakenz/kova/internal/core/task"

	"github.com/go-chi/chi/v5"
)

func NewRouter(taskStore *task.Store) http.Handler {
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	// Mount versioned API
	r.Route("/api", func(r chi.Router) {
		r.Mount("/v1", v1.Routes(taskStore))
	})

	return r
}
