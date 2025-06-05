package repository

import (
	"context"
	"database/sql"
	"task_management_system/internal/model"
)

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) Create(ctx context.Context, task model.Task) error {

	return nil
}

func (r *taskRepository) GetByID(ctx context.Context, id string) (model.Task, error) {
	return model.Task{}, nil
}

func (r *taskRepository) List(ctx context.Context, status string, limit, offset int) ([]model.Task, error) {

	return nil, nil
}

func (r *taskRepository) Update(ctx context.Context, task model.Task) error {
	return nil
}

func (r *taskRepository) Delete(ctx context.Context, id string) error {
	return nil
}
