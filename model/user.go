package model

import "time"

type User struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name      string `gorm:"unique;not null" json:"name"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"type:varchar(255)" json:"password"`
	Token     string `gorm:"-" json:"token,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
