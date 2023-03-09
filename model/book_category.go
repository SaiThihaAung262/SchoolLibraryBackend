package model

import (
	"time"

	"gorm.io/gorm"
)

type BookCategory struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `grom:"type:varchar(255)" json:"desc"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `json:"-"`
}
