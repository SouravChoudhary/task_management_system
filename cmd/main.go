package main

import (
	"log"
	"task_management_system/config"
	"task_management_system/internal/handler"
	"task_management_system/internal/repository"
	"task_management_system/internal/service"
	"task_management_system/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config.yaml: %v", err)
	}

	postgresConn, err := db.InitPostgres(&cfg.DB)
	if err != nil {
		log.Fatalf("Unable to start service due to config error: %v", err)
	}

	taskRepo := repository.NewTaskRepository(postgresConn)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	router := gin.Default()

	taskHandler.RegisterRoutes(router)

	router.Run(":8080")
}
