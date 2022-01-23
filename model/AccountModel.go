package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golang_web/config"
)

type (
	Account struct {
		ID       uint     `gorm:"primaryKey;autoIncrement:true;"`
		Name     string   `json:"name"`
		Email    string   `json:"email" gorm:"unique;not null`
		Password string   `json:"password" gorm:"unique;not null`
		Booklist Booklist `gorm:"foreignkey:AccountEmail;references:Email"`
	}

	GetAccountModel struct {
		Email string `json:"email"`
	}

	CreateAccountModel struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginInfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

//CreateAccount ... Insert New data
func CreatetAccountBooklist(account *Account, booklist *Booklist) (err error) {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		fmt.Println("create account:%s", account)
		if err = tx.Create(account).Error; err != nil {
			return err
		}

		if err = tx.Model(account).Association("Booklist").Append(booklist).Error; err != nil {
			return err
		}

		return nil
	})
}

func GetAllAccounts(account *[]Account) (err error) {
	if err = database.DB.Find(account).Error; err != nil {
		return err
	}
	return nil
}

func GetAccount(account *Account, email string) (err error) {
	fmt.Println("get account by email:%s", email)
	if err = database.DB.Where("email = ?", email).First(account).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func CheckAccountValid(account *Account, email string, password string) (err error) {
	if err = database.DB.Preload("Booklist").Where("email = ?", email).Where("password = ?", password).Find(account).Error; err != nil {
		return err
	}

	return nil
}
