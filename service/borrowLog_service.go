package service

import (
	"MyGO.com/m/model"
	"MyGO.com/m/repository"
)

type BorrowLogService interface {
	CreateBorrowLog(createLog model.BorrowLog) error
}

type borrowLogService struct {
	borrowLogRepo repository.BorrowLogRepository
}

func NewBorrowLogService(borrowLogRepo repository.BorrowLogRepository) BorrowLogService {
	return &borrowLogService{
		borrowLogRepo: borrowLogRepo,
	}
}

func (service borrowLogService) CreateBorrowLog(createLog model.BorrowLog) error {
	return service.borrowLogRepo.CreateBorrowLog(createLog)
}
