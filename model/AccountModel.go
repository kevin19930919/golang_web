package model

import (
	"fmt"
	"golang_web/config"
)

type (
	Account struct {
		ID       uint   `gorm:"primaryKey;autoIncrement:true;"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	GetAccountModel struct {
		Email string `json:"email"`
	}

	CreateAccountModel struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

//CreateAccount ... Insert New data
func CreatetAccount(account *Account) (err error) {
	fmt.Println("create account:%s", account)
	if err = database.DB.Create(account).Error; err != nil {
		return err
	}
	return nil
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
		return err
	}
	return nil
}
