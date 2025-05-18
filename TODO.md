# TODO.md — Kova MVP Roadmap

## Phase 1: Core Backend Setup

### 🔧 Project Initialization

- [/] Initialize Go module: `go mod init github.com/YOUR_USERNAME/kova`
- [/] Create project folder structure
- [ ] Add core dependencies:
  - [/] `chi` for routing
  - [ ] `sqlite3` driver
  - [/] `uuid` generator
  - [ ] Logging package (e.g., `slog`, `zap`, or `zerolog`)
  - [ ] `godotenv` for config (optional)

---

## 📦 Domain Logic: Tasks

### `internal/core/task/`

- [/] Define `Task` struct:
  - `ID`, `Title`, `Status`, `EstimateMinutes`, `StartTime`, `EndTime`
- [/] Create task status enum: `Todo`, `InProgress`, `Done`, `Aborted`
- [ ] Implement core logic:
  - [ ] Validate state transitions
  - [ ] Calculate remaining time
  - [ ] Detect overdue tasks

---

## 🧠 Application Logic

### `internal/app/task/`

- [ ] Implement use cases:
  - [ ] `StartTask(taskID, estimateMins)`
  - [ ] `CompleteTask(taskID)`
  - [ ] `GetActiveTasks(userID)`

---

## 🧩 Persistence Layer

### `internal/infra/db/`

- [/] Define SQLite schema (`tasks`, `users`, `projects`)
- [/] DB initialization logic
- [ ] Create `TaskRepository` interface:
  - [ ] `Create(task)`
  - [ ] `Update(task)`
  - [ ] `ListByUser(userID)`
  - [ ] `GetActive(userID)`

---

## 🌐 API (REST)

### `api/v1/`

- [ ] Define API routes:
  - [/] `POST   /v1/tasks` → Create new task
  - [ ] `POST   /v1/tasks/:id/start` → Start task with time estimate
  - [ ] `POST   /v1/tasks/:id/complete` → Mark task complete
  - [ ] `GET    /v1/tasks/active` → Fetch running tasks and countdown
- [ ] Setup middleware:
  - [ ] Logging
  - [ ] CORS
  - [ ] Request ID (optional)

---

## 🧪 Testing & Tooling

- [ ] Unit tests for `core/task/` logic
- [ ] Add `.env` or `config.yaml` support
- [ ] Script: `scripts/dev.sh` to start server

---

## 🌱 Optional Polish

- [ ] Add `README.md` with project goals and architecture
- [ ] Add `LICENSE` (MIT)
- [ ] Scaffold OpenAPI (Swagger) definition

---

> Once backend is functional, begin planning TUI or CLI frontend.
