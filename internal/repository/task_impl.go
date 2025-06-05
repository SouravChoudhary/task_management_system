package repository

import (
	"context"
	"database/sql"
	"task_management_system/internal/model"
	"time"

	"github.com/google/uuid"
)

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) Create(ctx context.Context, task model.Task) error {
	task.ID = uuid.New().String()
	task.CreatedAt = time.Now()
	_, err := r.db.ExecContext(ctx, "INSERT INTO tasks (id, title, status, created_at) VALUES ($1, $2, $3, $4)", task.ID, task.Title, task.Status, task.CreatedAt)
	return err
}

func (r *taskRepository) GetByID(ctx context.Context, id string) (model.Task, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, title, status, created_at FROM tasks WHERE id=$1", id)
	var task model.Task
	err := row.Scan(&task.ID, &task.Title, &task.Status, &task.CreatedAt)
	return task, err
}

func (r *taskRepository) List(ctx context.Context, status string, limit, offset int) ([]model.Task, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if status != "" {
		rows, err = r.db.QueryContext(ctx, "SELECT id, title, status, created_at FROM tasks WHERE status=$1 LIMIT $2 OFFSET $3", status, limit, offset)
	} else {
		rows, err = r.db.QueryContext(ctx, "SELECT id, title, status, created_at FROM tasks LIMIT $1 OFFSET $2", limit, offset)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Status, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *taskRepository) Update(ctx context.Context, task model.Task) error {
	_, err := r.db.ExecContext(ctx, "UPDATE tasks SET title=$1, status=$2 WHERE id=$3", task.Title, task.Status, task.ID)
	return err
}

func (r *taskRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM tasks WHERE id=$1", id)
	return err
}
