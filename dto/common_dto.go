package dto

type DeleteByIdDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
}

type BorrowHistoryRequest struct {
	Page     uint64 `json:"page" form:"page"`
	PageSize uint64 `json:"page_size" form:"page_size"`
	ID       uint64 `json:"id" form:"id"`
	UserUUID string `json:"user_uuid" form:"user_uuid"`
	BookUUID string `json:"book_uuid" form:"book_uuid"`
	Status   uint64 `json:"status" form:"status"`
}

type ClientLoginDTO struct {
	Type     uint64 `json:"type" form:"type" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type GetUserByUUIDDto struct {
	Type uint64 `json:"type" form:"type" binding:"required"`
	UUID string `json:"uuid" form:"uuid" binding:"required"`
}
