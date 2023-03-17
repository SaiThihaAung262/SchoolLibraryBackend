package model

import (
	"time"

	"gorm.io/gorm"
)

type Punishment struct {
	ID           uint64         `gorm:"column:id;primary_key:auto_increment" json:"id"`
	PackageName  string         `gorm:"column:package_name;type:varchar(125)" json:"package_name"`
	Duration     uint64         `gorm:"column:duration;" json:"duration"`
	PunishAmount uint64         `gorm:"column:punishment_amt;not null" json:"punishment_amt"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}
