package model

type (
	Booklist struct {
		ID           int     `gorm:"primaryKey;autoIncrement:true;"`
		AccountEmail string  `gorm:"unique;not null"`
		Books        []*Book `gorm:"many2many:booklist_book;"`
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
