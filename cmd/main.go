package main

import (
	"task_management_system/internal/handler"
	"task_management_system/internal/repository"
	"task_management_system/internal/service"
	"task_management_system/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {

	postgresConn := db.InitPostgres()

	taskRepo := repository.NewTaskRepository(postgresConn)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	router := gin.Default()

	taskHandler.RegisterRoutes(router)

	router.Run(":8080")
}
