package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	StudentLoginType = 2
)

type Student struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UUID      string `gorm:"unique;not null" json:"uuid"`
	Name      string `gorm:"type:varchar(255)" json:"name"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"type:varchar(255)" json:"password"`
	RoleNo    string `gorm:"unique;not null" json:"role_no"`
	Year      uint64 `gorm:"column:year;not null" json:"year"`
	Token     string `gorm:"-" json:"token,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}
