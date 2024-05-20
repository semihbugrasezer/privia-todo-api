package models

import (
	"time"
)

type TodoList struct {
	ID             uint       `json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
	Name           string     `json:"name"`
	CompletionRate float32    `json:"completion_rate"`
	UserID         uint       `json:"user_id"`
	TodoItems      []TodoItem `json:"todo_items"`
}
type TodoItem struct {
	ID         uint       `json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	TodoListID uint       `json:"todo_list_id"`
	Content    string     `json:"content"`
	Completed  bool       `json:"completed"`
}
