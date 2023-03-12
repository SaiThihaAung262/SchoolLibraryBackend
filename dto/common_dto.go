package dto

type DeleteByIdDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
}

type BorrowHistoryRequest struct {
	Page     uint64 `json:"page" form:"page"`
	PageSize uint64 `json:"page_size" form:"page_size"`
	ID       uint64 `json:"id" form:"id"`
}
