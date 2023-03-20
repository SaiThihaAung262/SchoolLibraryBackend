package model

import (
	"time"

	"gorm.io/gorm"
)

type BookCategory struct {
	ID           uint64 `gorm:"primary_key:auto_increment" json:"id"`
	CategoryName string `gorm:"unique;not null" json:"category_name"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `json:"-"`
}
