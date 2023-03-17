package model

import (
	"time"
)

type BorrowLog struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	BorrowID   uint64 `gorm:"column:borrow_id;not null" json:"borrow_id"`
	Type       uint64 `gorm:"column:type;not null" json:"type"`
	UserID     uint64 `gorm:"column:user_id;not null" json:"user_id"`
	UserUUID   string `gorm:"column:user_uuid;not null" json:"user_uuid"`
	UserName   string `gorm:"column:username;not null" json:"username"`
	BookID     uint64 `gorm:"column:book_id;not null" json:"book_id"`
	BookUUID   string `gorm:"column:book_uuid;not null" json:"book_uuid"`
	BookTitle  string `gorm:"column:book_title;not null" json:"book_title"`
	Department uint64 `gorm:"column:department;" json:"department"`
	RoleNo     string `gorm:"column:role_no;" json:"role_no"`
	Year       uint64 `gorm:"column:year;" json:"year"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
