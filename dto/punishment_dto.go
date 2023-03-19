package dto

type CreatePunishmentDto struct {
	PackageName         string `json:"package_name" form:"package_name" binding:"required"`
	DurationStart       uint64 `json:"duration_start" form:"duration_start" binding:"required"`
	DurationEnd         uint64 `json:"duration_end" form:"duration_end" binding:"required"`
	TeacherPunishAmount uint64 `json:"teacher_punishment_amt" form:"teacher_punishment_amt" binding:"required"`
	StudentPunishAmount uint64 `json:"student_punishment_amt" form:"student_punishment_amt" binding:"required"`
	StaffPunishAmount   uint64 `json:"staff_punishment_amt" form:"staff_punishment_amt" binding:"required"`
}

type UpdatePunishmentDto struct {
	ID                  uint64 `json:"id" form:"id" binding:"required"`
	PackageName         string `json:"package_name" form:"package_name"`
	DurationStart       uint64 `json:"duration_start" form:"duration_start" binding:"required"`
	DurationEnd         uint64 `json:"duration_end" form:"duration_end" binding:"required"`
	TeacherPunishAmount uint64 `json:"teacher_punishment_amt" form:"teacher_punishment_amt"`
	StudentPunishAmount uint64 `json:"student_punishment_amt" form:"student_punishment_amt"`
	StaffPunishAmount   uint64 `json:"staff_punishment_amt" form:"staff_punishment_amt"`
}
