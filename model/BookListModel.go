package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golang_web/config"
)

type (
	Booklist struct {
		ID           uint    `gorm:"primaryKey;autoIncrement:true;"`
		AccountEmail string  `gorm:"not null"`
		Account      Account `gorm:"foreignkey:AccountEmail;references:AccountEmail"`
		Books        []*Book `gorm:"many2many:booklist_book;"`
	}

	CreateBooklistInfo struct {
		BookIDs      []string `json:"book_ids"`
		AccountEmail string   `json:"account_email"`
	}

	DeleteBooklistInfo struct {
		ID string `uri:"id"`
	}
)

func CreateBooklist(booklist *Booklist, books []*Book, account *Account) (err error) {
	booklist.Account = *account
	// check book is not ordered
	booklist.Books = books

	return database.DB.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Preload("Books").Find(booklist).Error; err != nil {
			return err
		}

		if err := tx.Create(booklist).Error; err != nil {
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

}

func GetBooklist(booklist *Booklist, id string) (err error) {
	fmt.Println("get booklist by id:%s", id)
	if err = database.DB.Where("id = ?", id).First(booklist).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBooklist(booklist *Booklist) (err error) {

	return database.DB.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// clear association
		if err := tx.Model(booklist).Association("Books").Clear().Error; err != nil {
			return err
		}
		// delete booklist
		if err := tx.Delete(booklist).Error; err != nil {
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

}
