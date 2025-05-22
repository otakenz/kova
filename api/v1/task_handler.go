package v1

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	apptask "github.com/otakenz/kova/internal/app/task"
	"github.com/otakenz/kova/internal/core/task"
	"github.com/otakenz/kova/pkg/logger"
)

type TaskHandler struct {
	TaskService *apptask.TaskService
}

func NewTaskHandler(TaskService *apptask.TaskService) *TaskHandler {
	return &TaskHandler{TaskService: TaskService}
}

// CreateTask Method
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var t task.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		logger.Sugar.Errorw("failed to decode task", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := h.TaskService.CreateTask(ctx, &t)
	if err != nil {
		logger.Sugar.Errorw("failed to create task", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Sugar.Infow("task created", "id", created.ID, "title", created.Title)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

// ListTask Method
func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := h.TaskService.ListTasks(ctx)
	if err != nil {
		logger.Sugar.Errorw("failed to list tasks", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Sugar.Infow("tasks listed", "count", len(tasks))
	json.NewEncoder(w).Encode(tasks)
}

// GetTask by id Method
func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	task, err := h.TaskService.GetTask(ctx, id)
	if err != nil {
		logger.Sugar.Errorw("failed to get task", "id", id, "error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	logger.Sugar.Infow("task retrieved", "id", task.ID, "title", task.Title)
	json.NewEncoder(w).Encode(task)
}

// UpdateTask by id Method
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	// Decode the task payload from body
	var t task.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		logger.Sugar.Errorw("failed to decode task", "error", err)
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	triggerStr := r.URL.Query().Get("trigger")
	trigger, err := task.ParseTrigger(triggerStr)
	if err != nil {
		logger.Sugar.Errorw("invalid trigger", "trigger", trigger, "error", err)
		http.Error(w, "invalid trigger", http.StatusBadRequest)
		return
	}

	t.ID = id
	updated, err := h.TaskService.UpdateTask(ctx, &t, &trigger)
	if err != nil {
		logger.Sugar.Errorw("failed to update task", "id", id, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Sugar.Infow("task updated", "id", updated.ID, "title", updated.Title)
	json.NewEncoder(w).Encode(updated)
}

// DeleteTask by id Method
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if err := h.TaskService.DeleteTask(ctx, id); err != nil {
		logger.Sugar.Errorw("failed to delete task", "id", id, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Sugar.Infow("task deleted", "id", id)
	w.WriteHeader(http.StatusNoContent)
}
