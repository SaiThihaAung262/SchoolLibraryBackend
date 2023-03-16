package model

import (
	"time"

	"gorm.io/gorm"
)

type SystemConfig struct {
	ID                    uint   `gorm:"column:id;primary_key:auto_increment" json:"id"`
	TeacherCanBorrowCount uint64 `gorm:"column:teacher_can_borrow_count;not null" json:"teacher_can_borrow_count"`
	StudentCanBorrowCount uint64 `gorm:"column:student_can_borrow_count;not null" json:"student_can_borrow_count"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `json:"-"`
}
