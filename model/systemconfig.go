package model

import (
	"time"

	"gorm.io/gorm"
)

type SystemConfig struct {
	Id                    uint   `gorm:"column:id;primary_key:auto_increment" json:"id"`
	TeacherCanBorrowCount uint64 `gorm:"column:teacher_can_borrow_count;" json:"teacher_can_borrow_count"`
	StudentCanBorrowCount uint64 `gorm:"column:student_can_borrow_count;" json:"student_can_borrow_count"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `json:"-"`
}
