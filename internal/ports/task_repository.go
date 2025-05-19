package ports

import (
	"context"

	"github.com/otakenz/kova/internal/core/task"
)

type TaskRepository interface {
	Create(ctx context.Context, t *task.Task) error
	List(ctx context.Context) ([]*task.Task, error)
	Get(ctx context.Context, id string) (*task.Task, error)
	Update(ctx context.Context, t *task.Task) error
	Delete(ctx context.Context, id string) error
}
