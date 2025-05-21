package app

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/otakenz/kova/internal/core/task"
	"github.com/otakenz/kova/internal/ports"
)

type TaskService struct {
	TaskRepo ports.TaskRepository
}

func NewTaskService(TaskRepo ports.TaskRepository) *TaskService {
	return &TaskService{TaskRepo: TaskRepo}
}

func (s *TaskService) CreateTask(ctx context.Context, t *task.Task) (*task.Task, error) {
	t.ID = uuid.NewString()
	t.ActualMin = 0
	t.CreatedAt = time.Now()
	t.UpdatedAt = t.CreatedAt

	if err := t.Validate(); err != nil {
		return nil, err
	}

	err := s.TaskRepo.Create(ctx, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *TaskService) ListTasks(ctx context.Context) ([]*task.Task, error) {
	return s.TaskRepo.List(ctx)
}

func (s *TaskService) GetTask(ctx context.Context, id string) (*task.Task, error) {
	return s.TaskRepo.Get(ctx, id)
}

func (s *TaskService) UpdateTask(ctx context.Context, t *task.Task, trigger *task.Trigger) (*task.Task, error) {
	if trigger != nil {
		// Create FSM starting from current task status
		fsm := task.NewStateMachine(t.Status)
		// Fire the trigger to attempt transition
		if err := fsm.Fire(string(*trigger)); err != nil {
			return nil, fmt.Errorf("invalid state transition: %w", err)
		}
		// Update task status with FSM's new state
		t.Status = task.Status(fsm.MustState().(string))
	}

	t.UpdatedAt = time.Now()

	if err := t.ValidateTitle(t.Title); err != nil {
		return nil, err
	}

	err := s.TaskRepo.Update(ctx, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	return s.TaskRepo.Delete(ctx, id)
}
