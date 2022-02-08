package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golang_web/config"
	"golang_web/model"
	"strconv"
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
		BookID string `json:"book_id" binding:"required"`
		Days   string `json:"end_date" binding:"required"`
	}

	GetOrderInfo struct {
		AccountEmail string `uri:"account_email" binding:"required"`
	}

	Order struct {
		OrderID int32 `json:"order_id"`
	}
)

func CreateOderRemoveList(order *model.Order, days string, book *model.Book, account *model.Account, booklist *model.Booklist) (err error) {
	local, _ := time.LoadLocation("Local")
	order.StartDate = time.Now().In(local)
	IntDays, _ := strconv.Atoi(days)
	order.EndDate = order.StartDate.AddDate(0, 0, IntDays)

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
			fmt.Println("fail to update book status")
			return err
		}
		if err := tx.Model(booklist).Association("Books").Delete(book).Error; err != nil {
			fmt.Println("fail to remove book in book list")
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
}

func GetUnReturnOrderByAccount(orders *[]model.Order, account *model.Account) (err error) {
	fmt.Println("get order by account:%s", account)
	if err = database.DB.Preload("Book").Preload("Account").Where("account_email = ?", account.Email).Where("state = ?", false).Find(&orders).Error; err != nil {
		return err
	}
	return nil
}

func (this *Order) GetOrderByID(order *model.Order) (err error) {
	if err := database.DB.Where("id = ?", this.OrderID).First(&order).Error; err != nil {
		fmt.Println("fail to get order model by id", err)
		return err
	}
	return nil
}

func (this *Order) ReturnBooks() (err error) {
	var order model.Order
	if err := this.GetOrderByID(&order); err != nil {
		return err
	}
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// preload book table
		if err := tx.Preload("Book").Find(&order).Error; err != nil {
			fmt.Println("fail to preload Book with Order:", err)
			return err
		}

		if err := tx.Model(&order).Update("state", true).Error; err != nil {
			fmt.Println("fail to update order status:", err)
			return err
		}
		if err := tx.Model(&order.Book).Update("status", 0).Error; err != nil {
			fmt.Println("fail to update Book status:", err)
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
}
