package repository

import (
	"task-tracker/db"
	"task-tracker/models"
)

type TaskRepository interface {
	GetTasks() ([]models.Task, error)
	GetTaskById(id int) (*models.Task, error)
	AddTask(task *models.Task) error
	EditTask(id int, task *models.Task) error
	DeleteTask(id int) error
}

type taskRepo struct {
	db db.Data
}

func newTaskRepository(db db.Data) *taskRepo {
	return &taskRepo{
		db: db,
	}
}

func (t *taskRepo) GetTasks() ([]models.Task, error) {
	return nil, nil
}

func (t *taskRepo) GetTaskById(id int) (*models.Task, error) {
	return nil, nil
}

func (t *taskRepo) AddTask(task *models.Task) error {
	return nil
}

func (t *taskRepo) EditTask(id int, task *models.Task) error {
	return nil
}

func (t *taskRepo) DeleteTask(id int) error {
	return nil
}
