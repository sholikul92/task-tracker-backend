package service

import (
	"task-tracker/models"
	"task-tracker/repository"
)

type TaskService interface {
	GetTasks() ([]models.Task, error)
	GetTaskById(id int) (*models.Task, error)
	AddTask(task *models.Task) error
	EditTask(id int, task *models.Task) error
	DeleteTask(id int) error
}

type taskService struct {
	repo repository.TaskRepository
}

func newTaskService(repo repository.TaskRepository) *taskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) GetTasks() ([]models.Task, error) {
	return nil, nil
}
func (s *taskService) GetTaskById(id int) (*models.Task, error) {
	return nil, nil
}
func (s *taskService) AddTask(task *models.Task) error {
	return nil
}
func (s *taskService) EditTask(id int, task *models.Task) error {
	return nil
}
func (s *taskService) DeleteTask(id int) error {
	return nil
}
