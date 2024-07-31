package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"tostrike/controller"
	"tostrike/db"
	"tostrike/repository"
	"tostrike/usecase"
)

func main() {

	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	TaskRepository := repository.NewTaskRepository(dbConnection)
	TaskUseCase := usecase.NewTaskUseCase(TaskRepository)
	TaskController := controller.NewTaskController(TaskUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/tasks", TaskController.GetTasks)
	server.GET("/task/:id", TaskController.GetTaskById)
	server.POST("/task", TaskController.CreateTask)

	server.Run(":8080")
}
