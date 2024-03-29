package dto

type StaffRegisterDTO struct {
	Name       string `json:"name" form:"name" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	Department uint64 `json:"department" form:"department"`
}

type UpdateStaffDTO struct {
	ID         uint64 `json:"id" form:"id" binding:"required"`
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	UUID       string `json:"uuid" form:"uuid"`
	Department uint64 `json:"department" form:"department"`
}

type StaffGetRequest struct {
	Page       uint64 `json:"page" form:"page"`
	PageSize   uint64 `json:"page_size" form:"page_size"`
	UUID       string `json:"uuid" form:"uuid"`
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Department uint64 `json:"department" form:"department"`
}
