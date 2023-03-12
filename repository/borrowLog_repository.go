package repository

import (
	"fmt"

	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type BorrowLogRepository interface {
	CreateBorrowLog(borrowLog model.BorrowLog) error
}

type borrowLogConnection struct {
	connection *gorm.DB
}

func NewBorrowlogRepository(db *gorm.DB) BorrowLogRepository {
	return &borrowLogConnection{
		connection: db,
	}
}

func (db *borrowLogConnection) CreateBorrowLog(borrowLog model.BorrowLog) error {
	if err := db.connection.Save(&borrowLog).Error; err != nil {
		fmt.Println("-------Here have error in save borrow-----", err)
		return err
	}
	return nil
}
