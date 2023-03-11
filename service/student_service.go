package service

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/model"
	"MyGO.com/m/repository"
	"github.com/mashingan/smapping"
)

type StudentService interface {
	InsertStudent(client dto.StudentRegisterDTO) (*model.Student, error)
	GetAllStudents(req *dto.StudentGetRequest) ([]model.Student, int64, error)
	UpdateStudent(client dto.UpdateStudentDTO) (*model.Student, error)
	DeleteStudent(id uint64) error
}
type studentService struct {
	studentRepo repository.SutudentRepository
}

func NewStudentService(studentRepo repository.SutudentRepository) StudentService {
	return &studentService{
		studentRepo: studentRepo,
	}
}

func (service studentService) InsertStudent(student dto.StudentRegisterDTO) (*model.Student, error) {
	var studentToCreate model.Student
	if err := smapping.FillStruct(&studentToCreate, smapping.MapFields(&student)); err != nil {
		fmt.Println("--------Here is error in repository ------", err)
	}
	studentToCreate.UUID = helper.GenerateUUID()
	res, err := service.studentRepo.InsertStudent(studentToCreate)
	if err != nil {
		fmt.Println("----------Here is error in update service----------", err)
		return nil, err
	}
	return res, nil
}

func (service studentService) GetAllStudents(req *dto.StudentGetRequest) ([]model.Student, int64, error) {

	students, count, err := service.studentRepo.GetAllStudents(req)
	if err != nil {
		return nil, 0, err
	}
	return students, count, err
}

func (service studentService) UpdateStudent(student dto.UpdateStudentDTO) (*model.Student, error) {
	studentToUpdate := model.Student{}
	err := smapping.FillStruct(&studentToUpdate, smapping.MapFields(student))
	if err != nil {
		fmt.Println("------Have error in update bookcategory servcie ------", err.Error())
	}

	res, errRepo := service.studentRepo.UpdateStudent(studentToUpdate)
	if errRepo != nil {
		return nil, errRepo
	}

	return res, nil

}

func (service studentService) DeleteStudent(id uint64) error {
	err := service.studentRepo.DeleteStudent(id)
	return err
}
