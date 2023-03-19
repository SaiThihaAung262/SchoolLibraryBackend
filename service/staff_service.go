package service

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/model"
	"MyGO.com/m/repository"
	"github.com/mashingan/smapping"
)

type StaffService interface {
	InsertStaff(staff dto.StaffRegisterDTO) (*model.Staff, error)
	GetAllStaff(req *dto.StaffGetRequest) ([]model.Staff, int64, error)
	UpdateStaff(staff dto.UpdateStaffDTO) (*model.Staff, error)
	DeleteStaff(id uint64) error
	GetStaffByUUID(uuid string) (*model.Staff, error)
	VerifyLogin(name string, password string) interface{}
}
type staffService struct {
	staffRepo repository.StaffRepository
}

func NewStaffService(staffRepo repository.StaffRepository) StaffService {
	return &staffService{
		staffRepo: staffRepo,
	}
}

func (service staffService) InsertStaff(staff dto.StaffRegisterDTO) (*model.Staff, error) {
	var staffToCreate model.Staff
	if err := smapping.FillStruct(&staffToCreate, smapping.MapFields(&staff)); err != nil {
		fmt.Println("--------Here is error in repository ------", err)
	}
	staffToCreate.UUID = helper.GenerateUUID()
	res, err := service.staffRepo.InsertStaff(staffToCreate)
	if err != nil {
		fmt.Println("----------Here is error in update service----------", err)
		return nil, err
	}
	return res, nil
}

func (service staffService) GetAllStaff(req *dto.StaffGetRequest) ([]model.Staff, int64, error) {

	staffs, count, err := service.staffRepo.GetAllStaff(req)
	if err != nil {
		return nil, 0, err
	}
	return staffs, count, err
}

func (service staffService) UpdateStaff(staff dto.UpdateStaffDTO) (*model.Staff, error) {
	staffToUpdate := model.Staff{}
	err := smapping.FillStruct(&staffToUpdate, smapping.MapFields(&staff))
	if err != nil {
		fmt.Println("------Have error in update bookcategory servcie ------", err.Error())
	}

	res, errRepo := service.staffRepo.UpdateStaff(staffToUpdate)
	if errRepo != nil {
		return nil, errRepo
	}

	return res, nil

}

func (service staffService) DeleteStaff(id uint64) error {
	err := service.staffRepo.DeleteStaff(id)
	return err
}

func (service staffService) GetStaffByUUID(uuid string) (*model.Staff, error) {
	staff, err := service.staffRepo.GetStaffByUUID(uuid)
	return staff, err
}

func (service staffService) VerifyLogin(email string, password string) interface{} {
	res := service.staffRepo.VerifyLogin(email)
	if v, ok := res.(model.Staff); ok {
		isPassword := password == v.Password
		if v.Email == email && isPassword {
			return res
		}
		return false
	}
	return false
}
