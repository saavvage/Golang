package models

import "gorm.io/gorm"

// Task represent model of objects
type Task struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
