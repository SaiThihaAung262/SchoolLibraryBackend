package model

import (
	"time"

	"gorm.io/gorm"
)

type BookSubCategory struct {
	ID              uint64 `gorm:"primary_key:auto_increment" json:"id"`
	CategoryID      uint64 `gorm:"column:category_id;not null" json:"category_id"`
	SubCategoryName string `gorm:"unique;not null" json:"sub_category_name"`
	// Description     string `grom:"type:varchar(255)" json:"desc"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}
