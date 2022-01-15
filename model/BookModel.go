package model

import (
	"fmt"
	"golang_web/config"
	"time"
)

type (
	Book struct {
		ID    string `gorm:"primaryKey;"`
		Title string `json:"title"`
		Desc  string `jsno:"desc"`
	}

	GetBookModel struct {
		ID string `json:"BookID"`
	}
)

//CreateBook ... Insert New data
func CreatetBook(book *Book) (err error) {
	fmt.Println("create book:%s", book)
	if err = database.DB.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func GetAllBooks(book *[]Book) (err error) {
	if err = database.DB.Find(book).Error; err != nil {
		return err
	}
	return nil
}

func GetBook(book *Book, id string) (err error) {
	fmt.Println("get book by id:%s", id)
	if err = database.DB.Where("id = ?", id).First(book).Error; err != nil {
		return err
	}
	return nil
}

func CreateBookID(title string) string {
	t := time.Now()
	return title + t.String()
}
