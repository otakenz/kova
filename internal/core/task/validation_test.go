package task_test

import (
	"testing"
	"time"

	"github.com/otakenz/kova/internal/core/task"
)

func TestTask_Validate(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name    string
		task    task.Task
		wantErr bool
	}{
		{
			name: "valid task",
			task: task.Task{
				ID:          "1",
				Title:       "Sample Task",
				Description: "Testing",
				Status:      task.Todo,
				Priority:    task.Medium,
				EstimateMin: 60,
				ActualMin:   0,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			wantErr: false,
		},
		{
			name: "missing title",
			task: task.Task{
				Status:    task.Todo,
				Priority:  task.Low,
				CreatedAt: now,
				UpdatedAt: now,
			},
			wantErr: true,
		},
		{
			name: "title must be less than 255 characters",
			task: task.Task{
				Title:     string(make([]byte, 256)),
				Status:    task.Todo,
				Priority:  task.Low,
				CreatedAt: now,
				UpdatedAt: now,
			},
			wantErr: true,
		},
		{
			name: "done without completed_at",
			task: task.Task{
				Title:     "Finish report",
				Status:    task.Done,
				Priority:  task.High,
				CreatedAt: now,
				UpdatedAt: now,
			},
			wantErr: true,
		},
		{
			name: "completed_at with wrong status",
			task: task.Task{
				Title:       "Invalid",
				Status:      task.Todo,
				Priority:    task.Medium,
				CreatedAt:   now,
				UpdatedAt:   now,
				CompletedAt: &now,
			},
			wantErr: true,
		},
		{
			name: "negative estimate",
			task: task.Task{
				Title:       "Negative estimate",
				Status:      task.Todo,
				Priority:    task.Low,
				EstimateMin: -5,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			wantErr: true,
		},
		{
			name: "updated before created",
			task: task.Task{
				Title:     "Time travel",
				Status:    task.Todo,
				Priority:  task.Medium,
				CreatedAt: now,
				UpdatedAt: now.Add(-time.Minute),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.task.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
