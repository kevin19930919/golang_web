package model

import (
	"fmt"
	"golang_web/config"
)

type (
	TodoModel struct {
		ID        uint   `gorm:"primaryKey;autoIncrement:true;"`
		Title     string `json:"title"`
		Completed int    `json:"completed`
	}

	GetTodoModel struct {
		Title string `json:"title"`
	}

	CreateTodoModel struct {
		Title     string `json:"title"`
		Completed int    `json:"completed"`
	}
)

// assign table name
func (TodoModel) assign_table_name() string {
	return "todos"
}

//CreateUser ... Insert New data
func CreatetTodo(todo *TodoModel) (err error) {
	fmt.Println("create content:%s", todo)
	if err = database.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTodos(todo *[]TodoModel) (err error) {
	if err = database.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodo(todo *TodoModel, title string) (err error) {
	fmt.Println("get todo by title:%s", title)
	if err = database.DB.Where("title = ?", title).First(todo).Error; err != nil {
		return err
	}
	return nil
}
