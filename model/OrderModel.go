package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golang_web/config"
	"time"
)

type (
	Order struct {
		ID           uint      `gorm:"primaryKey;autoIncrement:true;"`
		CreateTime   time.Time `json:"create_time" gorm:"not null"`
		RemainTime   float64   `json:"remain_time" gorm:"not null"`
		Returned     bool      `json:"returned" gorm:"not null;default:false"`
		AccountEmail string    `gorm:"not null"`
		Account      Account   `gorm:"foreignkey:AccountEmail;references:AccountEmail"`
		Books        []*Book   `gorm:"many2many:order_book;"`
	}

	CreateOrderInfo struct {
		RemainTime   float64  `json:"remain_time"`
		BookIDs      []string `json:"book_ids"`
		AccountEmail string   `json:"account_email"`
	}
)

//CreateOrder ... Insert New data
func CreateOrder(order *Order, books []*Book, account *Account) (err error) {
	order.Account = *account
	// check book is not ordered
	order.Books = books

	local, _ := time.LoadLocation("Local")
	order.CreateTime = time.Now().In(local)

	return database.DB.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Preload("Books").Find(order).Error; err != nil {
			return err
		}

		if err := tx.Create(order).Error; err != nil {
			return err
		}
		if err := tx.Model(order.Books).Update("status", 1).Error; err != nil {
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

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

func GetOrderByID(order *Order, id int32) (err error) {
	if err = database.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

func UpdateOderStatusBookReturned(order *Order, returned bool, book_status int32) (err error) {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// preload book table
		if err := tx.Preload("Books").Find(order).Error; err != nil {
			return err
		}
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Model(order).Update("returned", returned).Error; err != nil {
			return err
		}

		if err := tx.Model(order.Books).Update("status", book_status).Error; err != nil {
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
}
