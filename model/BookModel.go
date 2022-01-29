package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golang_web/config"
)

type (
	Book struct {
		ID        string      `gorm:"primaryKey;"`
		Title     string      `json:"title"`
		Desc      string      `json:"desc"`
		Status    int32       `json:"status" gorm:"default:0"`
		Orders    []*Order    `gorm:"many2many:order_book;"`
		Booklists []*Booklist `gorm:"many2many:booklist_book;"`
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

func GetBookByID(book *Book, id string) (err error) {
	fmt.Println("get book by id:", id)
	err = database.DB.Where("id = ?", id).First(book).Error

	if err == gorm.ErrRecordNotFound {
		fmt.Println("Not found any book by id:", id)
		return nil

	} else if err != nil {
		fmt.Println("get book by id fail", err)
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
