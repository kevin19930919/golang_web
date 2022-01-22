package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golang_web/config"
)

type (
	Booklist struct {
		ID           int     `gorm:"primaryKey;autoIncrement:true;"`
		AccountEmail string  `gorm:"unique;not null"`
		Books        []*Book `gorm:"many2many:booklist_book;"`
	}

	CreateBooklistInfo struct {
		BookID       string `json:"book_id"`
		AccountEmail string `json:"account_email"`
	}

	DeleteBooklistInfo struct {
		ID     string `uri:"list_id" binding:"required"`
		BookID string `uri:"book_id" binding:"required"`
	}

	GetBooklistInfo struct {
		AccountEmail string `uri:"account_email" binding:"required"`
	}
)

func CreateBooklist(booklist *Booklist, books []*Book, account *Account) (err error) {
	// check book is not ordered
	booklist.Books = books

	return database.DB.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// if err := tx.Preload("Books").Find(booklist).Error; err != nil {
		// 	fmt.Print("related book by book list fail")
		// 	return err
		// }

		if err := tx.Create(booklist).Error; err != nil {
			return err
		}
		if err := tx.Model(&account).Association("Booklist").Append(booklist).Error; err != nil {
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

}

func GetBooklistByID(booklist *Booklist, id int) (err error) {
	fmt.Println("get booklist by id", id)
	if err = database.DB.Where("id = ?", id).First(booklist).Error; err != nil {
		fmt.Println("get booklist by id fail", err)
		return err
	}
	return nil
}

func GetBooklistByAccount(booklist *Booklist, account *Account) (err error) {

	// err = database.DB.Model(account).Related(booklist, "Booklist").Error
	err = database.DB.Preload("Booklist").Find(account).Error
	// gorm take not found record as an error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Not found any record by email:", account.Email)
		return nil

	} else if err != nil {

		fmt.Println("get booklist by email fail", err)
		return err
	}
	if err = database.DB.Preload("Books").Find(&(account.Booklist)).Error; err != nil {
		return err
	}
	*booklist = account.Booklist
	return nil
}

func DeleteBookFromBooklist(booklist *Booklist, book *Book) (err error) {

	return database.DB.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// clear association

		if err := tx.Model(booklist).Association("Books").Delete(book).Error; err != nil {
			fmt.Println("fail to remove book from booklist")
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

}
