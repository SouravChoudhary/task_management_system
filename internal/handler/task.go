package handler

import (
	"task_management_system/internal/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService service.TaskService
}

// Constructor
func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

// Route registration
func (h *TaskHandler) RegisterRoutes(router *gin.Engine) {
	taskRoutes := router.Group("/tasks")
	{
		taskRoutes.POST("", h.Create)
		taskRoutes.GET("", h.List)
		taskRoutes.GET("/:id", h.Get)
		taskRoutes.PUT("/:id", h.Update)
		taskRoutes.DELETE("/:id", h.Delete)
	}
}

// Handlers

func (h *TaskHandler) Create(c *gin.Context) {

}

func (h *TaskHandler) List(c *gin.Context) {

}

func (h *TaskHandler) Get(c *gin.Context) {

}

func (h *TaskHandler) Update(c *gin.Context) {

}

func (h *TaskHandler) Delete(c *gin.Context) {

}
