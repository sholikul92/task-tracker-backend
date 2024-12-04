package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title     string    `json:"title" binding:"required"`
	Priority  string    `json:"priority" binding:"required"`
	Deadline  time.Time `json:"deadline" binding:"required"`
	Completed bool      `json:"completed"`
}

type DBConfig struct {
	HOST     string
	PORT     int
	USER     string
	PASSWORD string
	DBNAME   string
}
