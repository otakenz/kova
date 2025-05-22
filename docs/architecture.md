## System Design

```css
[ Client (CLI / TUI / Browser UI) ]
             │  HTTP Requests
             ▼
        [ API Layer ]
   - Receives and validates HTTP requests
   - Calls core logic
   - Returns JSON responses
             │
             ▼
       [ Core Logic ]
   - Task & Project workflows
   - Time tracking and estimation logic
   - Validation & rules enforcement
             │
             ▼
       [ Persistence ]
   - SQLite DB access
   - Data models CRUD operations
```

## Data Flow (Tasks)

```css
[ HTTP Handler Layer ]
         ↓
[   Service Layer   ]  ← Business logic lives here
         ↓
[ Repository Layer ]  ← Interface to DB (SQLite, Postgres, mock, etc.)
         ↓
[   Infrastructure  ]  ← Actual database implementation

```

## Data Models

### Task

```go
type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      Status     `json:"status"`
	Priority    Priority   `json:"priority"`
	EstimateMin int        `json:"estimate_min"`
	ActualMin   int        `json:"actual_min"`
	AssignedTo  string     `json:"assigned_to"`
	StartedAt   *time.Time `json:"started_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CompletedAt time.Time `json:"completed_at"`
}
```

### Status

```go
type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in_progress"
	Done       Status = "done"
	Aborted    Status = "aborted"
)
```
