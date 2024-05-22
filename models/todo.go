package models

import (
	"time"
)

// TodoList represents a to-do list.
type TodoList struct {
	ID             uint `gorm:"primarykey"`
	UserID         uint `gorm:"index"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	Name           string
	CompletionRate float32
	User           User `gorm:"foreignkey:UserID"`
}

// TodoItem represents an item in a to-do list.
type TodoItem struct {
	ID         uint `gorm:"primarykey"`
	TodoListID uint `gorm:"index"`
	UserID     uint `gorm:"index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	Content    string
	Completed  bool
	TodoList   TodoList `gorm:"foreignkey:TodoListID"`
}
