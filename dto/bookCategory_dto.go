package dto

type CreateBookCategoryDTO struct {
	CategoryName string `json:"category_name" form:"category_name" binding:"required"`
}

type BookCategoryGetRequest struct {
	Page         uint64 `json:"page" form:"page"`
	PageSize     uint64 `json:"page_size" form:"page_size"`
	ID           uint64 `json:"id" form:"id" `
	CategoryName string `json:"category_name" form:"category_name"`
}

type UpdateBookCategoryDTO struct {
	ID           uint64 `json:"id" form:"id" binding:"required"`
	CategoryName string `json:"category_name" form:"category_name" binding:"required"`
}
