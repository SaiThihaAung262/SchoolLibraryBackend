package model

import (
	"time"

	"gorm.io/gorm"
)

type Borrow struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Type      uint64 `gorm:"column:type;not null" json:"type"`
	UserUUID  string `gorm:"column:user_uuid;not null" json:"user_uuid"`
	BookUUID  string `gorm:"column:book_uuid;not null" json:"book_uuid"`
	Status    uint64 `gorm:"column:status;not null" json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}
