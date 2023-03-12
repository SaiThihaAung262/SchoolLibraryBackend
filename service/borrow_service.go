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
	GetBorrowHistory(req *dto.BorrowHistoryRequest) ([]model.Borrow, int64, error)
	UpdateBorrowStatus(borrow dto.UpdateBorrowStatusDTO) (*model.Borrow, error)
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
	borrow.Status = 1
	return service.borrowRepo.CreateBorrow(borrow)
}

func (service borrowService) GetBorrowHistory(req *dto.BorrowHistoryRequest) ([]model.Borrow, int64, error) {
	res, total, err := service.borrowRepo.GetBorrowHistory(req)
	return res, total, err
}

func (service borrowService) UpdateBorrowStatus(borrow dto.UpdateBorrowStatusDTO) (*model.Borrow, error) {
	toUpdateBorrow := model.Borrow{}
	err := smapping.FillStruct(&toUpdateBorrow, smapping.MapFields(&borrow))
	if err != nil {
		fmt.Println("------Have error in update bookcategory servcie ------", err.Error())
	}

	res, errRepo := service.borrowRepo.UpdateBorrowStatus(toUpdateBorrow)
	if errRepo != nil {
		return nil, errRepo
	}

	return res, nil

}
