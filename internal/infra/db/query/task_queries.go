package query

const InitTask = `
	CREATE TABLE IF NOT EXISTS tasks (
        id TEXT PRIMARY KEY,
        title TEXT NOT NULL,
        description TEXT,
        status TEXT,
        priority TEXT,
        estimate_min INTEGER,
        actual_min INTEGER,
        assigned_to TEXT,
        created_at DATETIME,
        updated_at DATETIME,
        completed_at DATETIME
    );
	`

const InsertTask = `
INSERT INTO tasks (id, title, description, status, priority, estimate_min, actual_min, assigned_to, created_at, updated_at, completed_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

const SelectTaskByID = `
SELECT id, title, description, status, priority, estimate_min, actual_min, assigned_to, created_at, updated_at, completed_at
FROM tasks
WHERE id = ?
`

const ListTasks = `
SELECT id, title, description, status, priority, estimate_min, actual_min, assigned_to, created_at, updated_at, completed_at
FROM tasks
ORDER BY created_at DESC
`

const UpdateTask = `
UPDATE tasks
SET title = ?, description = ?, status = ?, priority = ?, estimate_min = ?, actual_min = ?, assigned_to = ?, updated_at = ?, completed_at = ?
WHERE id = ?
`

const DeleteTask = `
DELETE FROM tasks
WHERE id = ?
`
