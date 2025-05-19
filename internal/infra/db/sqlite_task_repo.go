package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/otakenz/kova/internal/core/task"
	"github.com/otakenz/kova/internal/infra/db/query"
)

// SqliteTaskRepo is a thread-safe in-memory storage for tasks
type SqliteTaskRepo struct {
	DB *sql.DB
}

func NewTaskRepo(DB *sql.DB) *SqliteTaskRepo {
	return &SqliteTaskRepo{DB: DB}
}

func (s *SqliteTaskRepo) Init(ctx context.Context) error {
	_, err := s.DB.ExecContext(ctx, query.InitTask)
	return err
}

func (s *SqliteTaskRepo) Create(ctx context.Context, t *task.Task) error {
	_, err := s.DB.ExecContext(ctx, query.InsertTask,
		t.ID,
		t.Title,
		t.Description,
		t.Status,
		t.Priority,
		t.EstimateMin,
		t.ActualMin,
		t.AssignedTo,
		t.CreatedAt,
		t.UpdatedAt,
		t.CompletedAt,
	)
	return err
}

func (s *SqliteTaskRepo) List(ctx context.Context) ([]*task.Task, error) {
	rows, err := s.DB.QueryContext(ctx, query.ListTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*task.Task
	for rows.Next() {
		var t task.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.Priority,
			&t.EstimateMin, &t.ActualMin, &t.AssignedTo, &t.CreatedAt, &t.UpdatedAt, &t.CompletedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	return tasks, nil
}

// Get
func (s *SqliteTaskRepo) Get(ctx context.Context, id string) (*task.Task, error) {
	row := s.DB.QueryRowContext(ctx, query.SelectTaskByID, id)

	var t task.Task
	err := row.Scan(
		&t.ID,
		&t.Title,
		&t.Description,
		&t.Status,
		&t.Priority,
		&t.EstimateMin,
		&t.ActualMin,
		&t.AssignedTo,
		&t.CreatedAt,
		&t.UpdatedAt,
		&t.CompletedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("task.Task not found")
	}
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// Update
func (s *SqliteTaskRepo) Update(ctx context.Context, t *task.Task) error {
	_, err := s.DB.ExecContext(ctx, query.UpdateTask,
		t.ID,
		t.Title,
		t.Description,
		t.Status,
		t.Priority,
		t.EstimateMin,
		t.ActualMin,
		t.AssignedTo,
		t.CreatedAt,
		t.UpdatedAt,
		t.CompletedAt,
	)

	return err
}

// Delete
func (s *SqliteTaskRepo) Delete(ctx context.Context, id string) error {
	_, err := s.DB.ExecContext(ctx, query.DeleteTask, id)
	return err
}
