package dto

type CreateBookDTO struct {
	Title      string `json:"title" form:"title" binding:"required"`
	CategoryID uint64 `json:"category_id" form:"category_id" binding:"required"`
	Status     uint64 `json:"status" form:"status" binding:"required"`
	Author     string `json:"author" form:"author" binding:"required"`
	Summary    string `json:"summary" from:"summary" binding:"required"`
}

type BookGetRequest struct {
	ID        uint64 `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	CategorID uint64 `json:"category_id" form:"category_id"`
	Status    uint64 `json:"status" form:"status"`
	Page      uint64 `json:"page" form:"page"`
	PageSize  uint64 `json:"page_size" form:"page_size"`
}

type UpdateBookDTO struct {
	ID         uint64 `json:"id" form:"id" binding:"required"`
	Title      string `json:"title" form:"title" binding:"required"`
	CategoryID uint64 `json:"category_id" form:"category_id" binding:"required"`
	Status     uint64 `json:"status" form:"status" binding:"required"`
	Author     string `json:"author" form:"author" binding:"required"`
	Summary    string `json:"summary" from:"summary" binding:"required"`
}
