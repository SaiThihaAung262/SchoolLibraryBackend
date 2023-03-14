package repository

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type BorrowLogRepository interface {
	CreateBorrowLog(borrowLog model.BorrowLog) error
	GetBorrowCountByBookUUIDAndDate(req *dto.ReqBorrowCountByBookUUIDAndDateDto) (uint64, error)
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

func (db *borrowLogConnection) GetBorrowCountByBookUUIDAndDate(req *dto.ReqBorrowCountByBookUUIDAndDateDto) (uint64, error) {
	var total uint64

	filter := " where id != 0"

	if req.BookUUID != "" {
		filter += fmt.Sprintf(" AND book_uuid = '%s'", req.BookUUID)
	}

	if req.StartDate != "" && req.EndDate != "" {
		filter += fmt.Sprintf(" AND created_at BETWEEN '%s' AND '%s'", req.StartDate, req.EndDate)
	}

	countQuery := fmt.Sprintf("SELECT COUNT(1) AS 'borrow_count' FROM borrow_logs %s GROUP BY book_title", filter)
	if err := db.connection.Raw(countQuery).Scan(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}
