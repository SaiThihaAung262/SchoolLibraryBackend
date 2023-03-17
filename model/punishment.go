package model

import (
	"time"

	"gorm.io/gorm"
)

type Punishment struct {
	ID                  uint64         `gorm:"column:id;primary_key:auto_increment" json:"id"`
	PackageName         string         `gorm:"column:package_name;type:varchar(125)" json:"package_name"`
	DurationStart       uint64         `gorm:"column:duration_start;" json:"duration_start"`
	DurationEnd         uint64         `gorm:"column:duration_end;" json:"duration_end"`
	TeacherPunishAmount uint64         `gorm:"column:teacher_punishment_amt;not null" json:"teacher_punishment_amt"`
	StudentPunishAmount uint64         `gorm:"column:student_punishment_amt;not null" json:"student_punishment_amt"`
	CreatedAt           time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt           gorm.DeletedAt `json:"-"`
}
