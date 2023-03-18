package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	BookAvailiableStatus = 1
	BookDamageLostStatus = 2
	BookAllGoneStatus    = 3
)

type Book struct {
	ID           uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UUID         string `gorm:"unique;not null" json:"uuid"`
	CategoryID   uint64 `gorm:"column:category_id;not null" json:"category_id"`
	Title        string `gorm:"unique;not null" json:"title"`
	Author       string `gorm:"type:varchar(250)" json:"author"`
	Summary      string `gorm:"type:varchar(250)" json:"summary"`
	Status       uint64 `gorm:"column:status;not null" json:"status"`
	BookImage    string `gorm:"type:varchar(250)" json:"book_image"`
	PublishDate  string `gorm:"column:publish_date;not null" json:"publish_date"`
	DownloadLink string `gorm:"column:download_link;not null" json:"download_link"`
	AvailableQty uint64 `gorm:"column:available_qty;not null" json:"available_qty"`
	BorrowQty    uint64 `gorm:"column:borrow_qty;not null" json:"borrow_qty"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `json:"-"`
}
