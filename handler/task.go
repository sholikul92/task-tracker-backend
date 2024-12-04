package handler

import (
	"log"
	"net/http"
	"strconv"
	"task-tracker/models"
	"task-tracker/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler interface {
	GetTasks(c *gin.Context)
	GetTaskById(c *gin.Context)
	AddTask(c *gin.Context)
	EditTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	UpdateStatus(c *gin.Context)
}

type taskHandler struct {
	service service.TaskService
}

func NewTaskHandler(s service.TaskService) *taskHandler {
	return &taskHandler{
		service: s,
	}
}

func (h *taskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.service.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, tasks)
}
func (h *taskHandler) GetTaskById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Task ID",
		})
		return
	}

	task, err := h.service.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Id is not found",
		})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *taskHandler) AddTask(c *gin.Context) {
	var taskPayload models.Task
	if err := c.ShouldBindJSON(&taskPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.AddTask(&taskPayload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error add task",
		})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "add task success",
	})
}

func (h *taskHandler) EditTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Task ID",
		})
		return
	}

	var taskPayload models.Task
	if err := c.ShouldBindJSON(&taskPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error decoded json",
		})
		return
	}

	if err := h.service.EditTask(id, &taskPayload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "update task failed",
		})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "update task success",
	})
}
func (h *taskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Task ID",
		})
		return
	}

	if err := h.service.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "delete task failed",
		})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "delete task success",
	})
}

func (h *taskHandler) UpdateStatus(c *gin.Context) {
	idParam := c.Param("id")
	statusStr := c.Query("status")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status, err := strconv.ParseBool(statusStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateStatusTask(id, status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "update status task success"})
}
