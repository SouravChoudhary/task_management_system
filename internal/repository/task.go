package repository

import (
	"context"
	"task_management_system/internal/model"
)

type TaskRepository interface {
	Create(ctx context.Context, task model.Task) error
	GetByID(ctx context.Context, id string) (model.Task, error)
	List(ctx context.Context, status string, limit, offset int) ([]model.Task, error)
	Update(ctx context.Context, task model.Task) error
	Delete(ctx context.Context, id string) error
}
