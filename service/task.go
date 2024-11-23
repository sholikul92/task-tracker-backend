package service

import (
	"task-tracker/models"
	"task-tracker/repository"
)

type TaskService interface {
	GetTasks() []models.Task
	GetTaskById(id int) (*models.Task, error)
	AddTask(task *models.Task) error
	EditTask(id int, task *models.Task) error
	DeleteTask(id int) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *taskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) GetTasks() []models.Task {
	tasks := s.repo.GetTasks()
	return tasks
}

func (s *taskService) GetTaskById(id int) (*models.Task, error) {
	task, err := s.repo.GetTaskById(id)
	if err != nil {
		return &models.Task{}, err
	}

	return task, nil
}

func (s *taskService) AddTask(task *models.Task) error {
	if err := s.repo.AddTask(task); err != nil {
		return err
	}
	return nil
}

func (s *taskService) EditTask(id int, task *models.Task) error {
	if err := s.repo.EditTask(id, task); err != nil {
		return err
	}
	return nil
}

func (s *taskService) DeleteTask(id int) error {
	if err := s.repo.DeleteTask(id); err != nil {
		return err
	}
	return nil
}
