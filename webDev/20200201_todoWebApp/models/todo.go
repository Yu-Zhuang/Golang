package models

import (
	"todoWebApp/dbc"

	"gorm.io/gorm"
)

// Todo model
type Todo struct {
	gorm.Model
	Content   string `json:"content" form:"content"`
	UserRefer uint
}

// CreateTodo 新增人
func CreateTodo(todo *Todo) (err error) {
	result := dbc.Db.Create(&todo)
	return result.Error
}
