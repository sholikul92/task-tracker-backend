package db

import (
	"errors"
	"fmt"
	"task-tracker/models"
)

type Data struct {
	TaskList []models.Task
}

func (d *Data) GetTasks() []models.Task {
	if d.TaskList == nil {
		d.TaskList = []models.Task{}
	}

	return d.TaskList
}

func (d *Data) GetTaskById(id int) (*models.Task, error) {
	for _, task := range d.TaskList {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("Task not found")
}

func (d *Data) AddTask(task *models.Task) error {
	id := len(d.TaskList) + 1
	task.ID = id

	d.TaskList = append(d.TaskList, *task)
	return nil
}

func (d *Data) EditTask(id int, task *models.Task) error {
	taskIndex := d.findTaskIndex(id)

	if taskIndex == -1 {
		return fmt.Errorf("Task with id %d not found", id)
	}

	task.ID = d.TaskList[taskIndex].ID
	d.TaskList[taskIndex] = *task
	return nil
}

func (d *Data) DeleteTask(id int) error {
	taskIndex := d.findTaskIndex(id)

	if taskIndex == -1 {
		return errors.New("id not found")
	}

	d.TaskList = append(d.TaskList[:taskIndex], d.TaskList[taskIndex+1:]...)
	return nil
}

func (d *Data) findTaskIndex(id int) int {
	for i, task := range d.TaskList {
		if task.ID == id {
			return i
		}
	}

	return -1
}
