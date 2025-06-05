package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"task_management_system/internal/constants"
	"task_management_system/internal/dto"
	"task_management_system/internal/service"
	"task_management_system/internal/utils"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

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

func (h *TaskHandler) Create(c *gin.Context) {
	var req dto.CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusUnprocessableEntity, "Invalid input: "+err.Error())
		return
	}

	if err := req.Validate(); err != nil {
		utils.Error(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	task, err := h.taskService.CreateTask(c.Request.Context(), req)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Created(c, task)
}

func (h *TaskHandler) List(c *gin.Context) {
	const maxLimit = 100
	const defaultLimit = 10

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		utils.Error(c, 400, "invalid page number")
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", fmt.Sprint(defaultLimit)))
	if err != nil || limit < 1 {
		utils.Error(c, 400, "invalid limit")
		return
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	offset := (page - 1) * limit

	status := c.Query("status")
	if status != "" && !constants.IsValidTaskStatus(status) {
		utils.Error(c, 400, "invalid status filter")
		return
	}

	tasks, err := h.taskService.ListTasks(c.Request.Context(), status, limit, offset)
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}

	utils.SuccessWithPagination(c, tasks, page, limit)
}

func (h *TaskHandler) Get(c *gin.Context) {
	id := c.Param("id")

	if !utils.IsValidUUID(id) {
		utils.Error(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	task, err := h.taskService.GetTask(c.Request.Context(), id)
	if err != nil {
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}

	utils.Success(c, task)
}

func (h *TaskHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateTaskRequest

	if !utils.IsValidUUID(id) {
		utils.Error(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := req.Validate(); err != nil {
		utils.Error(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := h.taskService.UpdateTask(c.Request.Context(), id, req); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "updated"})
}

func (h *TaskHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if !utils.IsValidUUID(id) {
		utils.Error(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	if err := h.taskService.DeleteTask(c.Request.Context(), id); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "deleted"})
}
