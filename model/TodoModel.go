package model

import (
	"github.com/jinzhu/gorm"
	"golang_web/config"
)

type (
	TodoModel struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed`
	}

	fmtTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)

// assign table name
func (TodoModel) assign_table_name() string {
	return "todos"
}

//CreateUser ... Insert New data
func CreatetTodo(todo *TodoModel) (err error) {
	if err = database.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}
