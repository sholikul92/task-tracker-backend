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
	UpdateStatusTask(id int, status bool) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *taskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) GetTasks() ([]models.Task, error) {
	tasks, err := s.repo.GetTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
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

func (s *taskService) UpdateStatusTask(id int, status bool) error {
	if _, err := s.repo.GetTaskById(id); err != nil {
		return err
	}

	if err := s.repo.UpdateStatusCompleted(id, status); err != nil {
		return err
	}
	return nil
}
