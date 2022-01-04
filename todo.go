package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {}

type (
	todoModel struct {
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
func (todoModel) assign_table_name() string {
	return "todos"
}

var db *gorm.DB

// initial db
func init() {
	var err error
	var constr string
	constr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", "localhost", "7552", "root", "golang_postgre", "kevin7552")
	fmt.Printf("%s\n", constr)
	db, err = gorm.Open("postgres", constr)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic("fail to connect db")
	}

	db.AutoMigrate(&todoModel{})
}
