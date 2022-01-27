package model

import (
	"time"
)

type (
	Order struct {
		ID           uint      `gorm:"primaryKey;autoIncrement:true;"`
		StartDate    time.Time `json:"start_date" gorm:"not null"`
		EndDate      time.Time `json:"end_date" gorm:"not null"`
		State        bool      `json:"returned" gorm:"not null;default:false"`
		AccountEmail string    `gorm:"not null"`
		Account      Account   `gorm:"foreignkey:AccountEmail;references:AccountEmail"`
		BookID       string    `gorm:"not null"`
		Book         Book      `gorm:"foreignkey:BookID;references:BookID"`
	}

	CreateOrderInfo struct {
		EndDate      time.Time `json:"end_date"`
		BookID       string    `json:"book_id"`
		AccountEmail string    `json:"account_email"`
	}

	CreateOrderByListInfo struct {
		BooklistID string `json:"list_id"`
	}

	ReturneOrderInfo struct {
		ID int32 `json:"order_id"`
	}
)

// //CreateOrder ... Insert New data
// func CreateOrder(order *Order, books []*Book, account *Account) (err error) {
// 	order.Account = *account
// 	// check book is not ordered
// 	order.Books = books

// 	local, _ := time.LoadLocation("Local")
// 	order.CreateTime = time.Now().In(local)

// 	return database.DB.Transaction(func(tx *gorm.DB) error {
// 		// do some database operations in the transaction (use 'tx' from this point, not 'db')
// 		if err := tx.Preload("Books").Find(order).Error; err != nil {
// 			return err
// 		}

// 		if err := tx.Create(order).Error; err != nil {
// 			return err
// 		}
// 		if err := tx.Model(order.Books).Update("status", 1).Error; err != nil {
// 			return err
// 		}
// 		// return nil will commit the whole transaction
// 		return nil
// 	})

// }

// func CreateOderRemoveList(order *Order, booklist *Booklist) (err error) {
// 	var account Account
// 	if err := GetAccount(&account, booklist.AccountEmail); err != nil {
// 		return err
// 	}
// 	order.Account = account
// 	// check book is not ordered
// 	order.Books = booklist.Books

// 	local, _ := time.LoadLocation("Local")
// 	order.CreateTime = time.Now().In(local)

// 	return database.DB.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Create(order).Error; err != nil {
// 			fmt.Println("fail to create order")
// 			return err
// 		}
// 		if err := tx.Model(order.Books).Update("status", 1).Error; err != nil {
// 			fmt.Println("fail to update book records")
// 			return err
// 		}
// 		// remove booklist association
// 		if err = tx.Model(booklist).Association("Books").Clear().Error; err != nil {
// 			fmt.Println("fail to remove book in book list")
// 			return err
// 		}
// 		// return nil will commit the whole transaction
// 		return nil
// 	})

// }

// func GetAllOrders(order *[]Order) (err error) {
// 	if err = database.DB.Find(order).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetOrderByAccount(order *Order, account *Account) (err error) {
// 	fmt.Println("get order by account:%s", account)
// 	if err = database.DB.Model(order).Related(account).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetOrderByID(order *Order, id int32) (err error) {
// 	if err = database.DB.Where("id = ?", id).First(order).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func UpdateOderStatusBookReturned(order *Order, returned bool, book_status int32) (err error) {
// 	return database.DB.Transaction(func(tx *gorm.DB) error {
// 		// preload book table
// 		if err := tx.Preload("Books").Find(order).Error; err != nil {
// 			return err
// 		}
// 		// do some database operations in the transaction (use 'tx' from this point, not 'db')
// 		if err := tx.Model(order).Update("returned", returned).Error; err != nil {
// 			return err
// 		}

// 		if err := tx.Model(order.Books).Update("status", book_status).Error; err != nil {
// 			return err
// 		}
// 		// return nil will commit the whole transaction
// 		return nil
// 	})
// }

// func ReturnBooks(order *Order) (err error) {
// 	return database.DB.Transaction(func(tx *gorm.DB) error {
// 		// preload book table
// 		if err := tx.Preload("Books").Find(order).Error; err != nil {
// 			return err
// 		}
// 		if err := tx.Model(order).Update("returned", true).Error; err != nil {
// 			return err
// 		}

// 		if err := tx.Model(order.Books).Update("status", 0).Error; err != nil {
// 			return err
// 		}
// 		// return nil will commit the whole transaction
// 		return nil
// 	})
// }
