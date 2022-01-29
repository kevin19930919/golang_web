package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golang_web/config"
	"golang_web/model"
)

type (
	CreateBookModel struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		Desc  string `jsno:"desc"`
	}

	GetBookModel struct {
		ID string `json:"BookID"`
	}

	Book struct {
		BookID string `json:"BookID"`
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

func (this *Book) CheckBookAvaliable() (status bool, err error) {
	var book model.Book
	err = database.DB.Where("id = ?", this.BookID).Where("status = ?", 1).First(&book).Error

	if err == gorm.ErrRecordNotFound {
		fmt.Println("book is avalible")
		return true, nil

	} else if err != nil {
		fmt.Println("get book by id fail", err)
		return false, err
	}
	return false, nil
}
