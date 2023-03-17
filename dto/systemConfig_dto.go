package dto

type CreateSystemConfigDto struct {
	TeacherCanBorrowCount uint64 `json:"teacher_can_borrow_count" form:"teacher_can_borrow_count" binding:"required"`
	StudentCanBorrowCount uint64 `json:"student_can_borrow_count" form:"student_can_borrow_count" binding:"required"`
	TeacherPunishAmt      uint64 `json:"teacher_punishment_amt" form:"teacher_punishment_amt"`
	StudentPunishAmt      uint64 `json:"student_punishment_amt" form:"student_punishment_amt"`
	TeacherCanBorrowDay   uint64 `json:"teacher_can_borrow_day" form:"teacher_can_borrow_day"`
	StudentCanBorrowDay   uint64 `json:"student_can_borrow_day" form:"student_can_borrow_day"`
}

type UpdateSystemConfigDto struct {
	ID                    uint64 `json:"id" form:"id" binding:"required"`
	TeacherCanBorrowCount uint64 `json:"teacher_can_borrow_count" form:"teacher_can_borrow_count" binding:"required"`
	StudentCanBorrowCount uint64 `json:"student_can_borrow_count" form:"student_can_borrow_count" binding:"required"`
	TeacherPunishAmt      uint64 `json:"teacher_punishment_amt" form:"teacher_punishment_amt"`
	StudentPunishAmt      uint64 `json:"student_punishment_amt" form:"student_punishment_amt"`
	TeacherCanBorrowDay   uint64 `json:"teacher_can_borrow_day" form:"teacher_can_borrow_day"`
	StudentCanBorrowDay   uint64 `json:"student_can_borrow_day" form:"student_can_borrow_day"`
}
