package service

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"MyGO.com/m/repository"
	"github.com/mashingan/smapping"
)

type PunishmentService interface {
	InsertPunishment(punishment dto.CreatePunishmentDto) (*model.Punishment, error)
	GetPunishmentData() ([]model.Punishment, error)
	UpdatePunishment(punishment dto.UpdatePunishmentDto) (*model.Punishment, error)
	DeletePunishment(id uint64) error
}

type punishmentService struct {
	punishmentRepo repository.PunishmentRepository
}

func NewPunishmentService(punishmentRepo repository.PunishmentRepository) PunishmentService {
	return &punishmentService{
		punishmentRepo: punishmentRepo,
	}
}

func (service punishmentService) InsertPunishment(punishment dto.CreatePunishmentDto) (*model.Punishment, error) {
	var punishmentToCreate model.Punishment
	if err := smapping.FillStruct(&punishmentToCreate, smapping.MapFields(&punishment)); err != nil {
		fmt.Println("--------Here is error in repository ------", err)
	}
	res, err := service.punishmentRepo.InsertPunishment(punishmentToCreate)
	if err != nil {
		fmt.Println("----------Here is error in update service----------", err)
		return nil, err
	}
	return res, nil
}

func (service punishmentService) GetPunishmentData() ([]model.Punishment, error) {
	students, err := service.punishmentRepo.GetPunishmentData()
	if err != nil {
		return nil, err
	}
	return students, err
}

func (service punishmentService) UpdatePunishment(punishment dto.UpdatePunishmentDto) (*model.Punishment, error) {
	punishmentToUpdate := model.Punishment{}
	err := smapping.FillStruct(&punishmentToUpdate, smapping.MapFields(&punishment))
	if err != nil {
		fmt.Println("------Have error in update bookcategory servcie ------", err.Error())
	}

	res, errRepo := service.punishmentRepo.UpdatePunishment(punishmentToUpdate)
	if errRepo != nil {
		return nil, errRepo
	}

	return res, nil
}

func (service punishmentService) DeletePunishment(id uint64) error {
	return service.punishmentRepo.DeletePunishment(id)
}
