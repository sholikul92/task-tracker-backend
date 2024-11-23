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
	tasks, err := t.db.GetTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *taskRepo) GetTaskById(id int) (*models.Task, error) {
	task, err := t.db.GetTaskById(id)
	if err != nil {
		return &models.Task{}, nil
	}

	return task, nil
}

func (t *taskRepo) AddTask(task *models.Task) error {
	if err := t.db.AddTask(task); err != nil {
		return err
	}
	return nil
}

func (t *taskRepo) EditTask(id int, task *models.Task) error {
	if err := t.db.EditTask(id, task); err != nil {
		return err
	}
	return nil
}

func (t *taskRepo) DeleteTask(id int) error {
	if err := t.db.DeleteTask(id); err != nil {
		return err
	}
	return nil
}
