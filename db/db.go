package db

import (
	"errors"
	"task-tracker/models"
)

type Data []models.Task

var Tasks Data

func (d *Data) GetTasks() ([]models.Task, error) {
	return Tasks, nil
}

func (d *Data) GetTaskById(id int) (*models.Task, error) {
	for _, task := range Tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("Task not found")
}

func (d *Data) AddTask(task *models.Task) error {
	task.ID += 1

	Tasks = append(Tasks, *task)
	return nil
}

func (d *Data) EditTask(id int, task *models.Task) error {
	taskFound, err := d.GetTaskById(id)
	if err != nil {
		return err
	}

	*taskFound = *task
	return nil
}

func (d *Data) DeleteTask(id int) error {
	var index int
	for i, task := range Tasks {
		if task.ID == id {
			index = i
			break
		}
	}
	if index == 0 {
		return errors.New("id not found")
	}

	Tasks = append(Tasks[:index], Tasks[index+1:]...)
	return nil
}
