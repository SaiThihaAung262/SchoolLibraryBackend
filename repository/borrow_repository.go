package repository

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type BorrowRepository interface {
	CreateBorrow(borrow model.Borrow) (*model.Borrow, error)
	GetBorrowHistory(req *dto.BorrowHistoryRequest) ([]model.Borrow, int64, error)
	UpdateBorrowStatus(borrow model.Borrow) (*model.Borrow, error)
	IsAlreadyBorrowThisBook(userUUID string, bookUUID string) (tx *gorm.DB)
	GetBorrowingAndExpireData(req *dto.BorrowHistoryRequest) ([]model.Borrow, int64, error)
}

type borrowConnection struct {
	connection *gorm.DB
}

func NewBorrowRepository(db *gorm.DB) BorrowRepository {
	return &borrowConnection{
		connection: db,
	}
}

func (db *borrowConnection) CreateBorrow(borrow model.Borrow) (*model.Borrow, error) {
	if err := db.connection.Save(&borrow).Error; err != nil {
		fmt.Println("-------Here have error in save borrow-----", err)
		return nil, err
	}
	return &borrow, nil
}
func (db *borrowConnection) GetBorrowingAndExpireData(req *dto.BorrowHistoryRequest) ([]model.Borrow, int64, error) {
	var borrowHistory []model.Borrow
	var total int64

	var offset uint64
	var pageSize uint64
	if req.Page != 0 {
		offset = (req.Page - 1) * req.PageSize
	} else {
		offset = 0
	}

	if req.PageSize != 0 {
		pageSize = req.PageSize
	} else {
		pageSize = 10

	}
	filter := " where deleted_at IS NULL"

	if req.UserUUID != "" {
		filter += fmt.Sprintf(" and user_uuid = '%s'", req.UserUUID)
	}

	filter += fmt.Sprintf(" and status = %d or status = %d", 1, 3)

	sql := fmt.Sprintf("select * from borrows %s order by created_at desc limit %v offset %v", filter, pageSize, offset)
	res := db.connection.Raw(sql).Scan(&borrowHistory)

	countQuery := fmt.Sprintf("select count(1) from borrows %s", filter)

	fmt.Println("Hre is count query", countQuery)
	if err := db.connection.Raw(countQuery).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	if res.Error == nil {
		return borrowHistory, total, nil
	}

	return nil, 0, nil

}

func (db *borrowConnection) GetBorrowHistory(req *dto.BorrowHistoryRequest) ([]model.Borrow, int64, error) {
	var borrowHistory []model.Borrow
	var total int64

	var offset uint64
	var pageSize uint64
	if req.Page != 0 {
		offset = (req.Page - 1) * req.PageSize
	} else {
		offset = 0
	}

	if req.PageSize != 0 {
		pageSize = req.PageSize
	} else {
		pageSize = 10

	}
	filter := " where deleted_at IS NULL"

	if req.ID != 0 {
		filter += fmt.Sprintf(" and id = %d", req.ID)
	}

	if req.Status != 0 {
		filter += fmt.Sprintf(" and status = %d", req.Status)
	}

	if req.Type != 0 {
		filter += fmt.Sprintf(" and type = %d", req.Type)
	}

	if req.UserUUID != "" {
		filter += fmt.Sprintf(" and user_uuid = '%s'", req.UserUUID)
	}

	if req.BookUUID != "" {
		filter += fmt.Sprintf(" and book_uuid = '%s'", req.BookUUID)
	}

	sql := fmt.Sprintf("select * from borrows %s order by created_at desc limit %v offset %v", filter, pageSize, offset)
	res := db.connection.Raw(sql).Scan(&borrowHistory)

	countQuery := fmt.Sprintf("select count(1) from borrows %s", filter)
	if err := db.connection.Raw(countQuery).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	if res.Error == nil {
		return borrowHistory, total, nil
	}

	return nil, 0, nil
}

func (db *borrowConnection) UpdateBorrowStatus(borrow model.Borrow) (*model.Borrow, error) {
	err := db.connection.Model(&borrow).Where("id = ?", borrow.ID).Updates(model.Borrow{
		Status: borrow.Status,
		Type:   borrow.Type,
	}).Error
	if err != nil {
		fmt.Println("----Here have error in update book repo -----")
		return nil, err

	}
	return &borrow, nil
}

func (db *borrowConnection) IsAlreadyBorrowThisBook(userUUID string, bookUUID string) (tx *gorm.DB) {
	var borrow model.Borrow
	return db.connection.Where("user_uuid = ? and book_uuid = ? and status = 1", userUUID, bookUUID).Take(&borrow)
}
