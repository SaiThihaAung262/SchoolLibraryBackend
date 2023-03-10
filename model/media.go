package model

import (
	"time"

	"gorm.io/gorm"
)

type Media struct {
	Id        uint64 `gorm:"column:id;primaryKey" json:"id"`
	FileName  string `gorm:"column:filename;type:varchar(125)" json:"filename"`
	URL       string `gorm:"column:url;type:varchar(125)" json:"url"`
	Extension string `gorm:"column:extension;type:varchar(50)" json:"extension"`
	Type      string `gorm:"column:type;type:enum('image', 'video');default:image" json:"type"`
	// UserId    uint64         `gorm:"column:user_id;not null" json:"user_id"`
	// BookId    uint64         `gorm:"column:book_id;not null" json:"book_id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
