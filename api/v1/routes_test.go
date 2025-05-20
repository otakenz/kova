package v1

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
//
// 	"github.com/otakenz/kova/internal/core/task"
// )

// type MockTaskRepo struct {
// 	tasks map[string]*task.Task
// }
//
// func NewMockTaskRepo() *MockTaskRepo {
// 	return &MockTaskRepo{tasks: make(map[string]*task.Task)}
// }
//
// func TestCreateTaskRoute(t *testing.T) {
// 	repo := NewMockTaskRepo()
// 	handler := Routes(repo)
//
// 	body := strings.NewReader(`{"title":"Test Task","status":"todo","priority":"low"}`)
// 	req, err := http.NewRequest("POST", "/tasks", body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	req.Header.Set("Content-Type", "application/json")
// 	rr := httptest.NewRecorder()
// 	handler.ServeHTTP(rr, req)
//
// 	if status := rr.Code; status != http.StatusCreated {
// 		t.Errorf("expected status %d, got %d", http.StatusCreated, status)
// 	}
//
// 	var resp task.Task
// 	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
// 		t.Errorf("failed to decode response: %v", err)
// 	}
// 	if resp.Title != "Test Task" {
// 		t.Errorf("expected title %q, got %q", "Test Task", resp.Title)
// 	}
// 	if resp.ID == "" {
// 		t.Errorf("expected non-empty id")
// 	}
// 	if resp.Status != "todo" {
// 		t.Errorf("expected status %q, got %q", "todo", resp.Status)
// 	}
// 	if resp.Priority != "low" {
// 		t.Errorf("expected priority %q, got %q", "low", resp.Priority)
// 	}
// }
//
// func TestListTasksRoute(t *testing.T) {
// 	taskStore := setupTaskStore()
// 	handler := Routes(taskStore)
//
// 	req, err := http.NewRequest("GET", "/tasks", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	rr := httptest.NewRecorder()
// 	handler.ServeHTTP(rr, req)
//
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("expected status %d, got %d", http.StatusOK, status)
// 	}
//
// 	expected := "[]"
// 	if strings.TrimSpace(rr.Body.String()) != expected {
// 		t.Errorf("expected body %s, got %s", expected, rr.Body.String())
// 	}
// }
