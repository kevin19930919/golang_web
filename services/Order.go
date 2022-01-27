package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golang_web/config"
	"golang_web/model"
	"time"
)

type (
	CreateOrderByListInfo struct {
		BooklistID string        `json:"list_id"`
		Orders     []OrderDetail `json:"orders"`
	}

	ReturneOrderInfo struct {
		ID int32 `json:"order_id"`
	}

	OrderDetail struct {
		BookID  string
		EndDate time.Time
	}
	Order struct {
		OrderID int32
	}
)

func CreateOderRemoveList(order *model.Order, book *model.Book, account *model.Account, booklist *model.Booklist) (err error) {
	local, _ := time.LoadLocation("Local")
	order.StartDate = time.Now().In(local)

	return database.DB.Transaction(func(tx *gorm.DB) error {
		//create order
		if err := tx.Create(order).Error; err != nil {
			fmt.Println("fail to create order")
			return err
		}
		if err = tx.Model(order).Association("Account").Append(*account).Error; err != nil {
			fmt.Println("fail to associate account")
			return err
		}
		if err = tx.Model(order).Association("Book").Append(*book).Error; err != nil {
			fmt.Println("fail to associate book")
			return err
		}
		if err := tx.Model(book).Update("status", 1).Error; err != nil {
			fmt.Println("fail to remove book in book list")
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
}

// func GetAllOrders(order *[]Order) (err error) {
// 	if err = database.DB.Find(order).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func GetOrderByAccount(order *model.Order, account *model.Account) (err error) {
	fmt.Println("get order by account:%s", account)
	if err = database.DB.Model(order).Related(account).Error; err != nil {
		return err
	}
	return nil
}

func (this *Order) GetOrderByID(order *Order) (err error) {
	if err = database.DB.Where("id = ?", this.OrderID).First(order).Error; err != nil {
		return err
	}
	return nil
}

func (this *Order) UpdateOderStatusBookReturned(order *model.Order, returned bool, book_status int32) (err error) {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// preload book table
		if err := tx.Preload("Books").Find(order).Error; err != nil {
			return err
		}
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Model(order).Update("returned", returned).Error; err != nil {
			return err
		}

		if err := tx.Model(order.Book).Update("status", book_status).Error; err != nil {
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
}

func (this *Order) ReturnBooks(order *model.Order) (err error) {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// preload book table
		if err := tx.Preload("Books").Find(order).Error; err != nil {
			return err
		}
		if err := tx.Model(order).Update("returned", true).Error; err != nil {
			return err
		}

		if err := tx.Model(order.Book).Update("status", 0).Error; err != nil {
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
}
