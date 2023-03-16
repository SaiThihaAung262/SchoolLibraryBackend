package dto

type CreateSystemConfigDto struct {
	TeacherCanBorrowCount uint64 `json:"teacher_can_borrow_count" form:"teacher_can_borrow_count" binding:"required"`
	StudentCanBorrowCount uint64 `json:"student_can_borrow_count" form:"student_can_borrow_count" binding:"required"`
}

type UpdateSystemConfigDto struct {
	ID                    uint64 `json:"id" form:"id" binding:"required"`
	TeacherCanBorrowCount uint64 `json:"teacher_can_borrow_count" form:"teacher_can_borrow_count" binding:"required"`
	StudentCanBorrowCount uint64 `json:"student_can_borrow_count" form:"student_can_borrow_count" binding:"required"`
}
