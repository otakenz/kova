package task

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const taskColumns = "id, title, description, status, estimate_min, actual_min, assigned_to, created_at, updated_at, completed_at"

// Store is a thread-safe in-memory storage for tasks
type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Init() error {
	query := `
    CREATE TABLE IF NOT EXISTS tasks (
        id TEXT PRIMARY KEY,
        title TEXT NOT NULL,
        description TEXT,
        status TEXT,
        estimate_min INTEGER,
        actual_min INTEGER,
        assigned_to TEXT,
        created_at DATETIME,
        updated_at DATETIME,
        completed_at DATETIME
    );`
	_, err := s.db.Exec(query)
	return err
}

func (s *Store) Create(t *Task) error {
	t.ID = uuid.NewString()
	t.ActualMin = 0
	t.CreatedAt = time.Now()
	t.UpdatedAt = t.CreatedAt

	_, err := s.db.Exec(
		`
        INSERT INTO tasks (id, title, description, status, estimate_min, actual_min, assigned_to, created_at, updated_at, completed_at)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		t.ID,
		t.Title,
		t.Description,
		t.Status,
		t.EstimateMin,
		t.ActualMin,
		t.AssignedTo,
		t.CreatedAt,
		t.UpdatedAt,
		t.CompletedAt,
	)
	return err
}

func (s *Store) List() ([]*Task, error) {
	rows, err := s.db.Query(
		`SELECT id, title, description, status, estimate_min, actual_min, assigned_to, created_at, updated_at, completed_at FROM tasks`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status,
			&t.EstimateMin, &t.ActualMin, &t.AssignedTo, &t.CreatedAt, &t.UpdatedAt, &t.CompletedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	return tasks, nil
}

// Get
func (s *Store) Get(id string) (*Task, error) {
	row := s.db.QueryRow(
		`SELECT id, title, description, status, estimate_min, actual_min, assigned_to, created_at,
							updated_at, completed_at FROM tasks WHERE id = ?`, id)

	var task Task
	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Status,
		&task.EstimateMin,
		&task.ActualMin,
		&task.AssignedTo,
		&task.CreatedAt,
		&task.UpdatedAt,
		&task.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("task not found")
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// Update
func (s *Store) Update(t *Task) error {
	_, err := s.db.Exec(
		`UPDATE tasks set title = ?, description = ?, status = ?, estimate_min = ?, actual_min = ?,
								assigned_to = ?, created_at = ?, updated_at = ?, completed_at = ? WHERE id = ?`,
		t.ID,
		t.Title,
		t.Description,
		t.Status,
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
func (s *Store) Delete(id string) error {
	_, err := s.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
