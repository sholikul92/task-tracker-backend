package main

import (
	"fmt"
	"log"
	"task-tracker/db"
	"task-tracker/handler"
	"task-tracker/repository"
	"task-tracker/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupRouter()

	fmt.Println("Server runing in http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	repo := repository.NewTaskRepository(db.Data{})
	service := service.NewTaskService(repo)
	handler := handler.NewTaskHandler(service)

	task := r.Group("task")
	{
		task.POST("/add", handler.AddTask)
		task.GET("/", handler.GetTasks)
		task.GET("/:id", handler.GetTaskById)
		task.PUT("/edit/:id", handler.EditTask)
		task.DELETE("/delete/:id", handler.DeleteTask)
	}

	return r
}
