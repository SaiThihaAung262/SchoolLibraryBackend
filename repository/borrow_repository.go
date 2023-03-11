package repository

import (
	"fmt"

	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type BorrowRepository interface {
	CreateBorrow(borrow model.Borrow) error
}

type borrowConnection struct {
	connection *gorm.DB
}

func NewBorrowRepository(db *gorm.DB) BorrowRepository {
	return &borrowConnection{
		connection: db,
	}
}

func (db *borrowConnection) CreateBorrow(borrow model.Borrow) error {
	if err := db.connection.Save(&borrow).Error; err != nil {
		fmt.Println("-------Here have error in save borrow-----", err)
		return err
	}
	return nil
}
