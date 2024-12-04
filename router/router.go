package router

import (
	"task-tracker/handler"
	"task-tracker/repository"
	"task-tracker/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	repo := repository.NewTaskRepository(db)
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
