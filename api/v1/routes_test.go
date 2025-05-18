package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/otakenz/kova/internal/core/task"
	"github.com/otakenz/kova/internal/infra/db"
)

func setupTaskStore() *task.Store {
	// TODO: mock the db implementing db with interface?
	db, err := db.New(":memory:")
	if err != nil {
		log.Fatal("failed to open DB:", err)
	}

	taskStore := task.NewStore(db)
	if err := taskStore.Init(); err != nil {
		log.Fatal("failed to init task store:", err)
	}
	return taskStore
}

func TestPingRoute(t *testing.T) {

	taskStore := setupTaskStore()
	handler := Routes(taskStore)

	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {

		t.Errorf("expected status %d, got %d", http.StatusOK, status)
	}

	expected := `{"message":"pong"}`
	if rr.Body.String() != expected {
		t.Errorf("expected body %q, got %q", expected, rr.Body.String())
	}
}

func TestCreateTaskRoute(t *testing.T) {
	db, err := db.New(":memory:")
	if err != nil {
		log.Fatal("failed to open DB:", err)
	}

	taskStore := task.NewStore(db)
	if err := taskStore.Init(); err != nil {
		t.Fatal(err)
	}

	handler := Routes(taskStore)

	body := strings.NewReader(`{"title":"Test Task","status":"todo"}`)
	req, err := http.NewRequest("POST", "/tasks", body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, status)
	}

	var resp task.Task
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	if resp.Title != "Test Task" {
		t.Errorf("expected title %q, got %q", "Test Task", resp.Title)
	}

	if resp.ID == "" {
		t.Errorf("expected non-empty id")
	}

	if resp.Status != "todo" {
		t.Errorf("expected status %q, got %q", "todo", resp.Status)
	}
}

func TestListTasksRoute(t *testing.T) {
	db, err := db.New(":memory:")
	if err != nil {
		log.Fatal("failed to open DB:", err)
	}

	taskStore := task.NewStore(db)
	if err := taskStore.Init(); err != nil {
		t.Fatal(err)
	}

	handler := Routes(taskStore)

	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, status)
	}

	expected := "[]"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("expected body %s, got %s", expected, rr.Body.String())
	}
}
