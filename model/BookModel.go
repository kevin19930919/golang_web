package model

import (
	"fmt"
	"golang_web/config"
)

type (
	Book struct {
		ID     string   `gorm:"primaryKey;"`
		Title  string   `json:"title"`
		Desc   string   `jsno:"desc"`
		Status int32    `json:"status" gorm:"default:0"`
		Orders []*Order `gorm:"many2many:order_book;"`
	}

	CreateBookModel struct {
		ID    string `json:"id"`
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
	if err = database.DB.Omit("Orders").Create(book).Error; err != nil {
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

func UpdateBookStatus(book *Book, status int32) (err error) {
	// status:
	//    0:avaliable
	//    1:ordered
	//    2:preserved
	if err = database.DB.Model(book).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
