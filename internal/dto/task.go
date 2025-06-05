package dto

import (
	"errors"
	"task_management_system/internal/constants"
)

type CreateTaskRequest struct {
	Title  string `json:"title" binding:"required"`
	Status string `json:"status" binding:"required"`
}

func (c *CreateTaskRequest) Validate() error {

	if !constants.IsValidTaskStatus(c.Status) {
		return errors.New("invalid status: must be 'Pending' or 'Completed'")
	}
	return nil
}

type UpdateTaskRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

func (c *UpdateTaskRequest) Validate() error {

	if !constants.IsValidTaskStatus(c.Status) {
		return errors.New("invalid status: must be 'Pending' or 'Completed'")
	}
	return nil
}
