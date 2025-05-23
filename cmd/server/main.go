package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/otakenz/kova/api"
	"github.com/otakenz/kova/internal/app/task"
	"github.com/otakenz/kova/internal/infra/db"
	"github.com/otakenz/kova/pkg/logger"
)

func main() {
	// Initialize logger
	if err := logger.Init(); err != nil {
		log.Fatal("failed to initialize logger:", err)
	}
	defer logger.Sync()

	logger.Sugar.Infow("Starting Kova server...")

	// Initialize SQLite database
	DB, err := db.New("kova.db")
	if err != nil {
		log.Fatal("failed to open DB:", err)
	}

	// Initialize Repository
	ctx := context.Background()
	taskRepo := db.NewTaskRepo(DB)
	if err := taskRepo.Init(ctx); err != nil {
		log.Fatal("failed to init task store:", err)
	}

	// Initialize Service
	taskService := task.NewTaskService(taskRepo)

	// Initialize API router
	router := api.NewRouter(taskService)

	// Initialize HTTP server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Channel to listen for interrupt of terminate signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start an anonymous function as a new goroutine
	go func() {
		log.Println("Starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Make current goroutine wait until a signal is received from stop channel
	<-stop
	log.Println("Shutting down server...")

	// Gracefully shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Failed to gracefully shutdown: %v", err)
	}

	log.Println("Server stopped")
}
