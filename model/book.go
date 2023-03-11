package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UUID       string `gorm:"unique;not null" json:"uuid"`
	CategoryID uint64 `gorm:"column:category_id;not null" json:"category_id"`
	Title      string `gorm:"unique;not null" json:"title"`
	Author     string `gorm:"type:varchar(250)" json:"author"`
	Summary    string `gorm:"type:varchar(250)" json:"summary"`
	Status     uint64 `gorm:"column:status;not null" json:"status"`
	BookImage  string `gorm:"type:varchar(250)" json:"book_image"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `json:"-"`
}
