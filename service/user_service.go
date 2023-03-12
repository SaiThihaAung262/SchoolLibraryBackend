package service

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"MyGO.com/m/repository"
	"github.com/mashingan/smapping"
)

type UserService interface {
	CreateUser(user dto.RegisterDTO) (*model.User, error)
	IsDuplicateEmail(email string) bool
	IsDuplicateName(name string) bool
	VerifyLogin(name string, password string) interface{}
	GetAllUsers(req *dto.UserGetRequest) ([]model.User, int64, error)
	UpdateUser(user dto.UpdateUserDto) (*model.User, error)
	IsUserExist(id uint64) bool
	DeleteUser(id uint64) error
	GetUserDashBoard() (*dto.DashboardResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (service userService) CreateUser(user dto.RegisterDTO) (*model.User, error) {
	userToCreate := model.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		fmt.Println("--------Here is error in repository ------", err)
	}
	res, err := service.userRepo.InsertUser(userToCreate)
	if err != nil {
		fmt.Println("----------Here is error in update service----------", err)
		return nil, err
	}
	return res, nil
}

func (service userService) IsDuplicateEmail(email string) bool {
	res := service.userRepo.IsDuplicateEmail(email)
	fmt.Println("____________res____________", res.Error)

	return (res.Error == nil)
}

func (service userService) IsDuplicateName(name string) bool {
	res := service.userRepo.IsDuplicateName(name)
	fmt.Println("____________res____________", res.Error)
	return (res.Error == nil)
}

func (service userService) VerifyLogin(name string, password string) interface{} {
	res := service.userRepo.VerifyLogin(name)
	if v, ok := res.(model.User); ok {
		// isPassword := comparePassword(password, v.Password)
		isPassword := password == v.Password
		if v.Name == name && isPassword {
			return res
		}
		return false
	}
	return false
}

// func comparePassword(enterPass string, resPassword string) bool {
// 	return enterPass == resPassword
// }

func (service userService) GetAllUsers(req *dto.UserGetRequest) ([]model.User, int64, error) {
	users, count, errr := service.userRepo.GetAllUser(req)
	if errr != nil {
		return nil, 0, errr
	}
	return users, count, nil
}

func (service userService) UpdateUser(user dto.UpdateUserDto) (*model.User, error) {
	userToUpdate := model.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		fmt.Println("----------Here is error in update service----------", err.Error())
	}
	res, err := service.userRepo.UpdateUser(userToUpdate)
	if err != nil {
		fmt.Println("----------Here is error in update service----------", err)
		return nil, err
	}
	return res, nil
}

func (service userService) IsUserExist(id uint64) bool {
	res := service.userRepo.IsUserExist(id)
	return (res.Error == nil)
}

func (service userService) DeleteUser(id uint64) error {
	res := service.userRepo.DeleteUser(id)
	return res
}

func (service userService) GetUserDashBoard() (*dto.DashboardResponse, error) {
	return service.userRepo.GetUserDashBoard()
}
