package task

import "time"

// Status represents task progress
type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in_progress"
	Done       Status = "done"
	Aborted    Status = "aborted"
)

// Task represents a single ticket/task
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	EstimateMin int       `json:"estimate_min"`
	ActualMin   int       `json:"actual_min"`
	AssignedTo  string    `json:"assigned_to"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CompletedAt time.Time `json:"completed_at"`
}
