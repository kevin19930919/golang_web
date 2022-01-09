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
	}

	GetBookModel struct {
		ID string `json:"BookID"`
	}

	CreateBookModel struct {
		Title string `json:"title"`
	}
)

//CreateBook ... Insert New data
func CreatetBook(book *Book) (err error) {
	book.ID = CreateBookID(book.Title)
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

func GetBook(book *Book, title string) (err error) {
	fmt.Println("get book by title:%s", title)
	if err = database.DB.Where("title = ?", title).First(book).Error; err != nil {
		return err
	}
	return nil
}

func CreateBookID(title string) string {
	t := time.Now()
	return title + t.String()
}
