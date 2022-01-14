package model

import (
	"fmt"
	"golang_web/config"
	"time"
)

type (
	Order struct {
		ID           uint      `gorm:"primaryKey;autoIncrement:true;"`
		CreateTime   time.Time `json:"create_time"`
		RemainTime   float64   `json:"remain_time"`
		AccountEmail string
		BookID       string  `gorm:"unique;not null"`
		Account      Account `gorm:"foreignkey:AccountEmail;references:AccountEmail"`
		Book         Book    `gorm:"foreignkey:BookID;references:BookID"`
	}

	CreateOrder struct {
		RemainTime   float64 `json:"remain_time"`
		BookTitle    string  `json:"book_title"`
		AccountEmail string  `json:"account_email"`
	}
)

//CreateBook ... Insert New data
func CreatetOrder(order *Order, book *Book, account *Account) (err error) {
	order.Account = *account
	order.Book = *book
	order.CreateTime = time.Now()
	if err = database.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func GetAllOrders(order *[]Order) (err error) {
	if err = database.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderByAccount(order *Order, account *Account) (err error) {
	fmt.Println("get order by account:%s", account)
	if err = database.DB.Model(account).Related(order).Error; err != nil {
		return err
	}
	return nil
}
