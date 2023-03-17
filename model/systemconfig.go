package model

import (
	"time"

	"gorm.io/gorm"
)

type SystemConfig struct {
	ID                    uint64 `gorm:"column:id;primary_key:auto_increment" json:"id"`
	TeacherCanBorrowCount uint64 `gorm:"column:teacher_can_borrow_count;not null" json:"teacher_can_borrow_count"`
	StudentCanBorrowCount uint64 `gorm:"column:student_can_borrow_count;not null" json:"student_can_borrow_count"`
	TeacherPunishAmt      uint64 `gorm:"column:teacher_punishment_amt;not null" json:"teacher_punishment_amt"`
	StudentPunishAmt      uint64 `gorm:"column:student_punishment_amt;not null" json:"student_punishment_amt"`
	TeacherCanBorrowDay   uint64 `gorm:"column:teacher_can_borrow_day;not null" json:"teacher_can_borrow_day"`
	StudentCanBorrowDay   uint64 `gorm:"column:student_can_borrow_day;not null" json:"student_can_borrow_day"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `json:"-"`
}
