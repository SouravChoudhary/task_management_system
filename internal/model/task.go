package model

import (
	"task_management_system/internal/dto"
	"time"
)

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (t *Task) ApplyUpdate(req dto.UpdateTaskRequest) {
	if req.Title != "" {
		t.Title = req.Title
	}
	if req.Status != "" {
		t.Status = req.Status
	}
}
