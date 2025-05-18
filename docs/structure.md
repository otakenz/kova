```graphql
kova/
├── cmd/                    # Entry points
│   ├── server/             # Starts REST API backend
│   ├── tui/                # (Optional) Starts TUI frontend
│   └── cli/                # (Optional) CLI command entry point
│
├── internal/               # Internal business logic (non-exported)
│   ├── app/                # Orchestrates use cases (service layer)
│   ├── core/               # Domain logic: models + pure functions
│   │   ├── task/						# Task models, ticket states, logic
│   │   ├── project/				# Project management (multi-project support)
│   │   ├── user/						# User profile and personal metrics
│   │   └── tracker/        # Time tracking & estimation
│   ├── infra/              # Infrastructure adapters
│   │   ├── db/             # SQLite setup and SQL repositories
│   │   └── logger/         # Logging setup
│   └── ports/              # Interfaces between app/core and infra/api
│
├── api/                    # HTTP interface (OpenAPI-friendly)
│   ├── v1/                 # API version 1 routes + handlers
│   └── middleware/         # Request logging, auth, etc. (Functions that run
															before your main handler)
│
├── pkg/                    # Reusable libs to be exported
├── ui/                     # TUI frontend code (if built-in)
├── web/                    # Web frontend (if embedded or developed later)
│
├── config/                 # Config files (env, YAML, etc.)
├── docs/                   # Developer docs, architecture, API schema
├── scripts/                # Dev/start/test helpers
├── test/                   # Integration or acceptance tests
│
├── go.mod
├── go.sum
├── README.md
└── LICENSE
```

| Layer         | Role                                                          |
| ------------- | ------------------------------------------------------------- |
| `core/`       | Domain logic: pure Go structs, methods, no dependencies       |
| `app/`        | Use cases: orchestrates logic (e.g., create task, track time) |
| `ports/`      | Interfaces (e.g., `TaskRepository`, `TrackerClock`)           |
| `infra/`      | Implements ports (e.g., SQLite, loggers)                      |
| `api/`        | HTTP layer, converts JSON ↔ domain models                    |
| `cmd/`        | Binaries: `server`, `cli`, or `tui`                           |
| `ui/`, `web/` | Optional frontends                                            |
