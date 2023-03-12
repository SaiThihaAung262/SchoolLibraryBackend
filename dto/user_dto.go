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
	TotalBook     int64 `json:"total_book" form:"total_book"`
	TotalCategory int64 `json:"total_category" form:"total_category"`
	TotalBorrow   int64 `json:"total_borrow" form:"total_borrow"`
	UnderBorrow   int64 `json:"under_borrow" form:"under_borrow"`
	HaveReturned  int64 `json:"have_returned" form:"have_returned"`
}
