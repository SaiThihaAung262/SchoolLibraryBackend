package dto

type CreateBorrowDTO struct {
	UserUUID string `json:"user_uuid" form:"user_uuid" binding:"required"`
	BookUUID string `json:"book_uuid" form:"book_uuid" binding:"required"`
	Type     uint64 `json:"type" form:"type" binding:"required"`
}
