package models

import "time"

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Priority  string    `json:"priority" binding:"required"`
	Deadline  time.Time `json:"deadline" binding:"required"`
	Completed bool      `json:"completed"`
}
