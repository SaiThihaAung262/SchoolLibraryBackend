package service

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"MyGO.com/m/repository"
	"github.com/mashingan/smapping"
)

type Borrowservice interface {
	CreateBorrow(createDto dto.CreateBorrowDTO) error
}

type borrowService struct {
	borrowRepo repository.BorrowRepository
}

func NewBorrowService(borrowRepo repository.BorrowRepository) Borrowservice {
	return &borrowService{
		borrowRepo: borrowRepo,
	}
}

func (service borrowService) CreateBorrow(createDto dto.CreateBorrowDTO) error {
	var borrow model.Borrow

	err := smapping.FillStruct(&borrow, smapping.MapFields(&createDto))
	if err != nil {
		fmt.Println("---Error in fill struct service ------", err.Error())
	}
	return service.borrowRepo.CreateBorrow(borrow)
}
