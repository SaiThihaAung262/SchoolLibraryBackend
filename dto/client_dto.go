package dto

type ClientRegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Type     uint64 `json:"type" form:"type" binding:"required"`
}

type UpdateClientDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Type     uint64 `json:"type" form:"type" binding:"required"`
	UUID     string `json:"uuid" form:"uuid" binding:"required"`
}

type ClientGetRequest struct {
	Page     uint64 `json:"page" form:"page"`
	PageSize uint64 `json:"page_size" form:"page_size"`
	UUID     string `json:"uuid" form:"uuid"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Type     uint64 `json:"Type" form:"Type"`
}
