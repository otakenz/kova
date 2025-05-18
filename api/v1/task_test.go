package v1

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/otakenz/kova/internal/core/task"
	"github.com/otakenz/kova/internal/infra/db"
)

func setupTestHandler() *TaskHandler {
	db, err := db.New(":memory:")
	if err != nil {
		log.Fatal("failed to open DB:", err)
	}
	store := task.NewStore(db)
	taskStore := task.NewStore(db)
	if err := taskStore.Init(); err != nil {
		log.Fatal("failed to init task store:", err)
	}
	return NewTaskHandler(store)
}

func TestCreateTask(t *testing.T) {
	handler := setupTestHandler()
	r := chi.NewRouter()
	r.Post("/tasks", handler.CreateTask)

	body := map[string]interface{}{
		"title":  "New Task",
		"status": "todo",
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/tasks", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, resp.Code)
	}

	var tsk task.Task
	if err := json.NewDecoder(resp.Body).Decode(&tsk); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if tsk.Title != "New Task" {
		t.Errorf("expected title %q, got %q", "New Task", tsk.Title)
	}
}

func TestListTasks(t *testing.T) {
	handler := setupTestHandler()
	handler.Store.Create(&task.Task{Title: "Task 1", Status: "todo"})

	r := chi.NewRouter()
	r.Get("/tasks", handler.ListTasks)
	req := httptest.NewRequest("GET", "/tasks", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", resp.Code)
	}

	var tasks []task.Task
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		t.Fatalf("failed to decode list: %v", err)
	}

	if len(tasks) == 0 {
		t.Errorf("expected at least one task")
	}
}

func TestGetTask(t *testing.T) {
	h := setupTestHandler()
	tsk := &task.Task{Title: "Find me", Status: "todo"}
	h.Store.Create(tsk)

	r := chi.NewRouter()
	r.Get("/tasks/{id}", h.GetTask)
	req := httptest.NewRequest("GET", "/tasks/"+tsk.ID, nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.Code)
	}

	var out task.Task
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		t.Fatalf("failed to decode: %v", err)
	}
	if out.ID != tsk.ID {
		t.Errorf("expected ID %q, got %q", tsk.ID, out.ID)
	}
}

func TestDeleteTask(t *testing.T) {
	h := setupTestHandler()
	tsk := &task.Task{Title: "Delete me", Status: "todo"}
	h.Store.Create(tsk)

	r := chi.NewRouter()
	r.Delete("/tasks/{id}", h.DeleteTask)
	req := httptest.NewRequest("DELETE", "/tasks/"+tsk.ID, nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusNoContent {
		t.Errorf("expected 204 No Content, got %d", resp.Code)
	}

	// Verify task is actually deleted
	if _, err := h.Store.Get(tsk.ID); err == nil {
		t.Errorf("task was not deleted")
	}
}
