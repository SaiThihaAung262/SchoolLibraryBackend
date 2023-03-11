package dto

type StudentRegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	RoleNo   string `json:"role_no" form:"role_no" binding:"required"`
	Year     uint64 `json:"year" form:"year" binding:"required"`
}

type UpdateStudentDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	UUID     string `json:"uuid" form:"uuid" binding:"required"`
	RoleNo   string `json:"role_no" form:"role_no" binding:"required"`
	Year     uint64 `json:"year" form:"year" binding:"required"`
}

type StudentGetRequest struct {
	Page     uint64 `json:"page" form:"page"`
	PageSize uint64 `json:"page_size" form:"page_size"`
	UUID     string `json:"uuid" form:"uuid"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	RoleNo   string `json:"role_no" form:"role_no" `
	Year     uint64 `json:"year" form:"year"`
}
