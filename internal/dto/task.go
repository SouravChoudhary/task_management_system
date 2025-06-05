package dto

type CreateTaskRequest struct {
	Title  string `json:"title" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type UpdateTaskRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}
