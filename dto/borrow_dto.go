package dto

import (
	"time"

	"MyGO.com/m/model"
)

type CreateBorrowDTO struct {
	UserUUID string `json:"user_uuid" form:"user_uuid" binding:"required"`
	BookUUID string `json:"book_uuid" form:"book_uuid" binding:"required"`
	Type     uint64 `json:"type" form:"type" binding:"required"`
}

type UpdateBorrowStatusDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	UserUUID string `json:"user_uuid" form:"user_uuid" binding:"required"`
	BookUUID string `json:"book_uuid" form:"book_uuid" binding:"required"`
	Type     uint64 `json:"type" form:"type" binding:"required"`
	Status   uint64 `json:"status" form:"status" binding:"required"`
}

type BorrowUser struct {
	ID         uint64 `json:"id"`
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department uint64 `json:"department"`
	Year       uint64 `json:"year"`
	RoleNo     string `json:"role_no"`
}

type BorrowHistoryResponse struct {
	ID        uint64      `json:"id"`
	Type      uint64      `json:"type"`
	Status    uint64      `json:"status"`
	User      BorrowUser  `json:"user_data"`
	Book      *model.Book `json:"book_data"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	ExpiredAt time.Time   `json:"expired_at"`
}

type BorrowHistoryList struct {
	List  []BorrowHistoryResponse `json:"list"`
	Total int64                   `json:"total"`
}
