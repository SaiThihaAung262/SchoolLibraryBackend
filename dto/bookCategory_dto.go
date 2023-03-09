package dto

type CreateBookCategoryDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"desc" form:"desc" binding:"required"`
}

type BookCategoryGetRequest struct {
	Page     uint64 `json:"page" form:"page"`
	PageSize uint64 `json:"page_size" form:"page_size"`
	ID       uint64 `json:"id" form:"id" `
	Title    string `json:"title" form:"title"`
}

type UpdateBookCategoryDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"desc" form:"desc" binding:"required"`
}
