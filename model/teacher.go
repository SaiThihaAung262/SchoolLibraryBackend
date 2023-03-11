package model

import (
	"time"

	"gorm.io/gorm"
)

type Teacher struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UUID       string `gorm:"unique;not null" json:"uuid"`
	Name       string `gorm:"type:varchar(255)" json:"name"`
	Email      string `gorm:"unique;not null" json:"email"`
	Password   string `gorm:"type:varchar(255)" json:"password"`
	Department uint64 `gorm:"column:status;not null" json:"status"`
	Token      string `gorm:"-" json:"token,omitempty"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `json:"-"`
}
