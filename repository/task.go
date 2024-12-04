package repository

import (
	"task-tracker/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks() ([]models.Task, error)
	GetTaskById(id int) (*models.Task, error)
	AddTask(task *models.Task) error
	EditTask(id int, task *models.Task) error
	DeleteTask(id int) error
	UpdateStatusCompleted(id int, status bool) error
}

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepo {
	return &taskRepo{db}
}

func (t *taskRepo) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := t.db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *taskRepo) GetTaskById(id int) (*models.Task, error) {
	var task models.Task
	err := t.db.Where("id = ?", id).First(&task).Error
	if err != nil {
		return &models.Task{}, err
	}

	return &task, nil
}

func (t *taskRepo) AddTask(task *models.Task) error {
	if err := t.db.Create(&task).Error; err != nil {
		return err
	}
	return nil
}

func (t *taskRepo) EditTask(id int, task *models.Task) error {
	if err := t.db.Where("ID = ?", id).Save(&task).Error; err != nil {
		return err
	}
	return nil
}

func (t *taskRepo) DeleteTask(id int) error {
	if err := t.db.Delete(&models.Task{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (t *taskRepo) UpdateStatusCompleted(id int, status bool) error {
	if err := t.db.Table("tasks").Where("ID = ?", id).Update("completed", status).Error; err != nil {
		return err
	}

	return nil
}
