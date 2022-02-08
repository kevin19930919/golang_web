package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golang_web/config"
	"golang_web/model"
)

type (
	Booklist struct {
		BooklistModel model.Booklist
	}

	InsertBooklistInfo struct {
		ID     string `json:"list_id" binding:"required"`
		BookID string `json:"book_id" binding:"required"`
	}

	DeleteBooklistInfo struct {
		ID     string `uri:"list_id" binding:"required"`
		BookID string `uri:"book_id" binding:"required"`
	}

	GetBooklistInfo struct {
		AccountEmail string `uri:"account_email" binding:"required"`
	}
)

func (this Booklist) CheckBookNotInBooklist(book_id string) (status bool) {
	for i := 0; i < len(this.BooklistModel.Books); {
		if this.BooklistModel.Books[i].ID == book_id {
			fmt.Println("book was added in list before")
			return false
		}
	}
	return true
}

func (this Booklist) InsertBookInBooklist(book *model.Book) (err error) {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Model(&this.BooklistModel).Association("Books").Append(book).Error; err != nil {
			fmt.Println("fail to insert book in book list")
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
}

// factory method
func GetBooklistByID(id int) (booklistmodel Booklist, err error) {
	var booklist model.Booklist

	if err = database.DB.Preload("Books").Where("id = ?", id).First(&booklist).Error; err != nil {
		fmt.Println("get booklist by id fail", err)
		return booklistmodel, err
	}
	booklistmodel.BooklistModel = booklist
	return booklistmodel, nil
}

func GetBooklistByAccount(booklist *model.Booklist, account *model.Account) (err error) {

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

func (this Booklist) DeleteBookFromBooklist(book *model.Book) (err error) {

	return database.DB.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// clear association

		if err := tx.Model(&this.BooklistModel).Association("Books").Delete(book).Error; err != nil {
			fmt.Println("fail to remove book from booklist")
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

}
