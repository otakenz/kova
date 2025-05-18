package v1

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/otakenz/kova/internal/core/task"
)

type TaskHandler struct {
	Store *task.Store
}

func NewTaskHandler(store *task.Store) *TaskHandler {
	return &TaskHandler{Store: store}
}

// CreateTask Method
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var t task.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Store.Create(&t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

// ListTask Method
func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Store.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

// GetTask by id Method
func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	task, err := h.Store.Get(id)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

// UpdateTask by id Method
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var input task.Task
	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	input.ID = id
	if err := h.Store.Update(&input); err != nil {
		http.Error(w, "update failed", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(input)
}

// DeleteTask by id Method
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.Store.Delete(id); err != nil {
		http.Error(w, "delete failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
