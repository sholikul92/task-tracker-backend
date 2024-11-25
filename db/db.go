package db

import (
	"errors"
	"fmt"
	"task-tracker/models"
)

type Data []models.Task

var Tasks Data

func (d *Data) GetTasks() []models.Task {
	return Tasks
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
	id := len(Tasks) + 1
	task.ID = id

	Tasks = append(Tasks, *task)
	return nil
}

func (d *Data) EditTask(id int, task *models.Task) error {
	taskIndex := d.findTaskIndex(id)

	if taskIndex == -1 {
		return fmt.Errorf("Task with id %d not found", id)
	}

	task.ID = Tasks[taskIndex].ID
	Tasks[taskIndex] = *task
	return nil
}

func (d *Data) DeleteTask(id int) error {
	taskIndex := d.findTaskIndex(id)

	if taskIndex == -1 {
		return errors.New("id not found")
	}

	Tasks = append(Tasks[:taskIndex], Tasks[taskIndex+1:]...)
	return nil
}

func (d *Data) findTaskIndex(id int) int {
	var indexTask int
	for i, task := range Tasks {
		if task.ID == id {
			indexTask = i
			break
		}
		indexTask = -1
	}

	return indexTask
}
