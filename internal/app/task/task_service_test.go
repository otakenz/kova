package task_test

import (
	"context"
	"errors"
	"testing"
	"time"

	apptask "github.com/otakenz/kova/internal/app/task"
	"github.com/otakenz/kova/internal/core/task"
)

type mockTaskRepo struct {
	CreateFn func(ctx context.Context, t *task.Task) error
	ListFn   func(ctx context.Context) ([]*task.Task, error)
	GetFn    func(ctx context.Context, id string) (*task.Task, error)
	UpdateFn func(ctx context.Context, t *task.Task) error
	DeleteFn func(ctx context.Context, id string) error
}

func (m *mockTaskRepo) Create(ctx context.Context, t *task.Task) error {
	return m.CreateFn(ctx, t)
}

func (m *mockTaskRepo) List(ctx context.Context) ([]*task.Task, error) {
	return m.ListFn(ctx)
}

func (m *mockTaskRepo) Get(ctx context.Context, id string) (*task.Task, error) {
	return m.GetFn(ctx, id)
}

func (m *mockTaskRepo) Update(ctx context.Context, t *task.Task) error {
	return m.UpdateFn(ctx, t)
}

func (m *mockTaskRepo) Delete(ctx context.Context, id string) error {
	return m.DeleteFn(ctx, id)
}

func TestTaskService_CreateTask(t *testing.T) {
	mockRepo := &mockTaskRepo{}
	service := apptask.NewTaskService(mockRepo)

	validTask := &task.Task{
		Title:       "Test Task",
		Status:      task.Todo,
		Priority:    task.Medium,
		EstimateMin: 30,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Happy path: Create succeeds
	mockRepo.CreateFn = func(ctx context.Context, t *task.Task) error {
		return nil
	}

	createdTask, err := service.CreateTask(context.Background(), validTask)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if createdTask.ID == "" {
		t.Error("expected ID to be set")
	}
	if createdTask.ActualMin != 0 {
		t.Errorf("expected ActualMin to be 0, got %d", createdTask.ActualMin)
	}

	// Validation failure: Missing title
	badTask := &task.Task{
		Status:    task.Todo,
		Priority:  task.Low,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = service.CreateTask(context.Background(), badTask)
	if err == nil {
		t.Error("expected validation error but got none")
	}

	// Repo failure
	mockRepo.CreateFn = func(ctx context.Context, t *task.Task) error {
		return errors.New("db error")
	}
	_, err = service.CreateTask(context.Background(), validTask)
	if err == nil || err.Error() != "db error" {
		t.Errorf("expected db error, got %v", err)
	}
}

func TestTaskService_UpdateTask(t *testing.T) {
	mockRepo := &mockTaskRepo{}
	service := apptask.NewTaskService(mockRepo)

	taskToUpdate := &task.Task{
		ID:          "123",
		Title:       "Update Me",
		Status:      task.Todo,
		Priority:    task.High,
		CreatedAt:   time.Now().Add(-time.Hour),
		UpdatedAt:   time.Now().Add(-time.Hour),
		EstimateMin: 60,
	}

	mockRepo.UpdateFn = func(ctx context.Context, t *task.Task) error {
		return nil
	}

	updatedTask, err := service.UpdateTask(context.Background(), taskToUpdate, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if updatedTask.UpdatedAt.Before(taskToUpdate.CreatedAt) {
		t.Error("UpdatedAt should be updated to current time")
	}

	// Validation failure (empty title)
	taskToUpdate.Title = ""
	_, err = service.UpdateTask(context.Background(), taskToUpdate, nil)
	if err == nil {
		t.Error("expected validation error due to empty title")
	}
	taskToUpdate.Title = "Update Me" // restore

	// Repo failure
	mockRepo.UpdateFn = func(ctx context.Context, t *task.Task) error {
		return errors.New("update failed")
	}
	_, err = service.UpdateTask(context.Background(), taskToUpdate, nil)
	if err == nil || err.Error() != "update failed" {
		t.Errorf("expected update failed error, got %v", err)
	}
}
