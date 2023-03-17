package dto

type CreatePunishmentDto struct {
	PackageName         string `json:"package_name" form:"package_name" binding:"required"`
	Duration            uint64 `json:"duration" form:"duration" binding:"required"`
	TeacherPunishAmount uint64 `json:"teacher_punishment_amt" form:"teacher_punishment_amt" binding:"required"`
	StudentPunishAmount uint64 `json:"student_punishment_amt" form:"student_punishment_amt" binding:"required"`
}

type UpdatePunishmentDto struct {
	ID                  uint64 `json:"id" form:"id" binding:"required"`
	PackageName         string `json:"package_name" form:"package_name"`
	Duration            uint64 `json:"duration" form:"duration"`
	TeacherPunishAmount uint64 `json:"teacher_punishment_amt" form:"teacher_punishment_amt"`
	StudentPunishAmount uint64 `json:"student_punishment_amt" form:"student_punishment_amt"`
}
