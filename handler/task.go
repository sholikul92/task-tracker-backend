package handler

import (
	"task-tracker/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler interface {
	GetTasks(c *gin.Context)
	GetTaskById(c *gin.Context)
	AddTask(c *gin.Context)
	EditTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

type taskHandler struct {
	service service.TaskService
}

func newTaskHandler(s service.TaskService) *taskHandler {
	return &taskHandler{
		service: s,
	}
}

func (h *taskHandler) GetTasks(c *gin.Context) {

}
func (h *taskHandler) GetTaskById(c *gin.Context) {

}
func (h *taskHandler) AddTask(c *gin.Context) {

}
func (h *taskHandler) EditTask(c *gin.Context) {

}
func (h *taskHandler) DeleteTask(c *gin.Context) {

}
