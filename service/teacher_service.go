package service

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/model"
	"MyGO.com/m/repository"
	"github.com/mashingan/smapping"
)

type TeacherService interface {
	InsertTeacher(teacher dto.TeacherRegisterDTO) (*model.Teacher, error)
	GetAllTeachers(req *dto.TeacherGetRequest) ([]model.Teacher, int64, error)
	UpdateTeacher(teacher dto.UpdateTeacherDTO) (*model.Teacher, error)
	DeleteTeacher(id uint64) error
	GetTeacherByUUID(uuid string) (*model.Teacher, error)
	VerifyLogin(name string, password string) interface{}
	ChangePassword(id uint64, password string) error
}
type teacherService struct {
	teacherRepo repository.TeacherRepository
}

func NewTeacherService(teacherRepo repository.TeacherRepository) TeacherService {
	return &teacherService{
		teacherRepo: teacherRepo,
	}
}

func (service teacherService) InsertTeacher(teacher dto.TeacherRegisterDTO) (*model.Teacher, error) {
	var teacherToCreate model.Teacher
	if err := smapping.FillStruct(&teacherToCreate, smapping.MapFields(&teacher)); err != nil {
		fmt.Println("--------Here is error in repository ------", err)
	}
	teacherToCreate.UUID = helper.GenerateUUID()
	res, err := service.teacherRepo.InsertTeacher(teacherToCreate)
	if err != nil {
		fmt.Println("----------Here is error in update service----------", err)
		return nil, err
	}
	return res, nil
}

func (service teacherService) GetAllTeachers(req *dto.TeacherGetRequest) ([]model.Teacher, int64, error) {

	teachers, count, err := service.teacherRepo.GetAllTeachers(req)
	if err != nil {
		return nil, 0, err
	}
	return teachers, count, err
}

func (service teacherService) UpdateTeacher(teacher dto.UpdateTeacherDTO) (*model.Teacher, error) {
	teacherToUpdate := model.Teacher{}
	err := smapping.FillStruct(&teacherToUpdate, smapping.MapFields(&teacher))
	if err != nil {
		fmt.Println("------Have error in update bookcategory servcie ------", err.Error())
	}

	res, errRepo := service.teacherRepo.UpdateTeacher(teacherToUpdate)
	if errRepo != nil {
		return nil, errRepo
	}

	return res, nil

}

func (service teacherService) DeleteTeacher(id uint64) error {
	err := service.teacherRepo.DeleteTeacher(id)
	return err
}

func (service teacherService) GetTeacherByUUID(uuid string) (*model.Teacher, error) {
	teacher, err := service.teacherRepo.GetTeacherByUUID(uuid)
	return teacher, err
}

func (service teacherService) VerifyLogin(email string, password string) interface{} {
	res := service.teacherRepo.VerifyLogin(email)
	if v, ok := res.(model.Teacher); ok {
		isPassword := password == v.Password
		if v.Email == email && isPassword {
			return res
		}
		return false
	}
	return false
}

func (service teacherService) ChangePassword(id uint64, password string) error {
	return service.teacherRepo.ChangePassword(id, password)
}
