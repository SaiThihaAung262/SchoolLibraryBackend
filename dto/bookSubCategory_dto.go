package dto

type CreateBookSubCategoryDTO struct {
	CategoryID      uint64 `json:"category_id" form:"category_id" binding:"required"`
	SubCategoryName string `json:"sub_category_name" form:"sub_category_name" binding:"required"`
	Description     string `json:"desc" form:"desc" binding:"required"`
}

type BookSubCategoryGetRequest struct {
	Page            uint64 `json:"page" form:"page"`
	PageSize        uint64 `json:"page_size" form:"page_size"`
	ID              uint64 `json:"id" form:"id" `
	SubCategoryName string `json:"sub_category_name" form:"sub_category_name"`
}

type UpdateBookSubCategoryDTO struct {
	ID              uint64 `json:"id" form:"id" binding:"required"`
	CategoryID      uint64 `json:"category_id" form:"category_id" binding:"required"`
	SubCategoryName string `json:"sub_category_name" form:"sub_category_name" binding:"required"`
	Description     string `json:"desc" form:"desc" binding:"required"`
}
