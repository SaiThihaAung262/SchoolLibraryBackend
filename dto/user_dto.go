package dto

type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserGetRequest struct {
	Page     uint64 `json:"page" form:"page"`
	PageSize uint64 `json:"page_size" form:"page_size"`
	ID       uint64 `json:"id" form:"id"`
}

type CreateUserDto struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Emial    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UpdateUserDto struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type DashboardResponse struct {
	TotalAdmin    int64 `json:"total_admin" form:"total_admin"`
	TotalTeacher  int64 `json:"total_teacher" form:"total_teacher"`
	TotalStudent  int64 `json:"total_student" form:"total_student"`
	TotalStaff    int64 `json:"total_staff" form:"total_staff"`
	TotalBook     int64 `json:"total_book" form:"total_book"`
	TotalCategory int64 `json:"total_category" form:"total_category"`
	TotalBorrow   int64 `json:"total_borrow" form:"total_borrow"`
	UnderBorrow   int64 `json:"under_borrow" form:"under_borrow"`
	HaveReturned  int64 `json:"have_returned" form:"have_returned"`
	ExpiredCount  int64 `json:"expired_count" form:"expired_count"`
}

type ReqMostBorrowData struct {
	Page     uint64 `json:"page" form:"page"`
	PageSize uint64 `json:"page_size" form:"page_size"`
}

type MostBorrowBookData struct {
	BookID      uint64 `json:"book_id" form:"book_id"`
	BookUUID    string `json:"book_uuid" form:"book_uuid"`
	BookTitle   string `json:"book_title" form:"book_title"`
	BorrowCount uint64 `json:"borrow_count" form:"borrow_count"`
}

type MostBorrowLogRespList struct {
	List  []MostBorrowBookData `json:"list" form:"list"`
	Total uint64               `json:"total" form:"total"`
}
