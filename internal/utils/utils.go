package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Page    int         `json:"page,omitempty"`
	Limit   int         `json:"limit,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{Success: true, Data: data})
}

func SuccessWithPagination(c *gin.Context, data interface{}, page, limit int) {
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    data,
		Page:    page,
		Limit:   limit,
	})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, APIResponse{Success: true, Data: data})
}

func Error(c *gin.Context, status int, msg string) {
	c.JSON(status, APIResponse{Success: false, Error: msg})
}

func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
