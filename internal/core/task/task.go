package task

import "time"

// Status represents task progress
type Status string

// Priority represents task priority
type Priority string

const (
	Todo       Status = "todo"
	InProgress Status = "in_progress"
	Done       Status = "done"
	Aborted    Status = "aborted"
)

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)

// TODO: split out DTOs (json tags) to separate API DTO

// Task represents a single ticket/task
type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      Status     `json:"status"`
	Priority    Priority   `json:"priority"`
	EstimateMin int        `json:"estimate_min"`
	ActualMin   int        `json:"actual_min"`
	AssignedTo  *string    `json:"assigned_to"`
	StartedAt   *time.Time `json:"started_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

func (t *Task) RemainingMin() int {
	return t.EstimateMin - t.ActualMin
}

func (t *Task) RemainingNow() int {
	if t.Status != InProgress || t.StartedAt.IsZero() {
		return t.RemainingMin()
	}
	elapsed := int(time.Since(*t.StartedAt).Minutes())
	return t.RemainingMin() - elapsed
}

func (t *Task) IsOverdue() bool {
	if t.Status != InProgress {
		return false
	}
	elapsed := time.Since(*t.StartedAt).Minutes()
	return elapsed > float64(t.EstimateMin)
}
