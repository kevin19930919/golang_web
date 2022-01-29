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
