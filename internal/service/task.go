package service

import (
	"context"
	"task_management_system/internal/dto"
	"task_management_system/internal/model"
	"task_management_system/internal/repository"
)

type TaskService interface {
	CreateTask(ctx context.Context, req dto.CreateTaskRequest) (model.Task, error)
	GetTask(ctx context.Context, id string) (model.Task, error)
	ListTasks(ctx context.Context, status string, limit, offset int) ([]model.Task, error)
	UpdateTask(ctx context.Context, id string, req dto.UpdateTaskRequest) error
	DeleteTask(ctx context.Context, id string) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo}
}

func (s *taskService) CreateTask(ctx context.Context, req dto.CreateTaskRequest) (model.Task, error) {

	return model.Task{}, nil
}

func (s *taskService) GetTask(ctx context.Context, id string) (model.Task, error) {
	return model.Task{}, nil
}

func (s *taskService) ListTasks(ctx context.Context, status string, limit, offset int) ([]model.Task, error) {
	return nil, nil
}

func (s *taskService) UpdateTask(ctx context.Context, id string, req dto.UpdateTaskRequest) error {
	return nil
}

func (s *taskService) DeleteTask(ctx context.Context, id string) error {
	return nil
}
