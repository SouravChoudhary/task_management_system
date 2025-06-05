package service

import (
	"context"
	"database/sql"
	"errors"
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

var ErrTaskNotFound = errors.New("task not found")

func (s *taskService) CreateTask(ctx context.Context, req dto.CreateTaskRequest) (model.Task, error) {

	task := model.Task{
		Title:  req.Title,
		Status: req.Status,
	}

	err := s.repo.Create(ctx, task)
	return task, err
}

func (s *taskService) GetTask(ctx context.Context, id string) (model.Task, error) {
	task, err := s.repo.GetByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Task{}, ErrTaskNotFound
	}
	return task, err
}

func (s *taskService) ListTasks(ctx context.Context, status string, limit, offset int) ([]model.Task, error) {
	return s.repo.List(ctx, status, limit, offset)
}

func (s *taskService) UpdateTask(ctx context.Context, id string, req dto.UpdateTaskRequest) error {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrTaskNotFound
		}
		return err
	}

	task.ApplyUpdate(req)
	return s.repo.Update(ctx, task)
}

func (s *taskService) DeleteTask(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
