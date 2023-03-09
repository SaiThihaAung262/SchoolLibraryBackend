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
