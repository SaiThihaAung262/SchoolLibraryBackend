package repository

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user model.User) (*model.User, error)
	IsDuplicateEmail(email string) (tx *gorm.DB)
	IsDuplicateName(name string) (tx *gorm.DB)
	VerifyLogin(name string) interface{}
	GetAllUser(req *dto.UserGetRequest) ([]model.User, int64, error)
	UpdateUser(user model.User) (*model.User, error)
	IsUserExist(id uint64) (tx *gorm.DB)
	DeleteUser(id uint64) error
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user model.User) (*model.User, error) {
	err := db.connection.Save(&user).Error
	if err != nil {
		fmt.Println("------------Here is error in user repository--------------", err)
		return nil, err
	}
	return &user, nil
}

func (db *userConnection) VerifyLogin(name string) interface{} {
	var user model.User

	res := db.connection.Where("name = ?", name).Take(&user)

	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) GetAllUser(req *dto.UserGetRequest) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	var offset uint64
	var pageSize uint64
	if req.Page != 0 {
		offset = (req.Page - 1) * req.PageSize
	} else {
		offset = 0
	}

	if req.PageSize != 0 {
		pageSize = req.PageSize
	} else {
		pageSize = 10

	}

	var filter string

	if req.ID != 0 {
		filter = fmt.Sprintf("where id = %v", req.ID)
	}

	sql := fmt.Sprintf("select * from users %s limit %v offset %v", filter, pageSize, offset)
	res := db.connection.Raw(sql).Scan(&users)

	countQuery := fmt.Sprintf("select count(1) from users %s", filter)
	if err := db.connection.Raw(countQuery).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	if res.Error == nil {
		return users, total, nil
	}

	return nil, 0, nil
}

func (db *userConnection) IsUserExist(id uint64) (tx *gorm.DB) {
	var user model.User
	return db.connection.Where("id = ?", id).Take(&user)
}

func (db *userConnection) UpdateUser(user model.User) (*model.User, error) {

	fmt.Println("-----------------Here is error in update user id ------------------", user.ID)

	err := db.connection.Model(&user).Where("id = ?", user.ID).Updates(model.User{Name: user.Name, Email: user.Email, Password: user.Password}).Error
	if err != nil {
		fmt.Println("-----------------Here is error in update user repository ------------------", err)
		return nil, err
	}
	return &user, nil
}

func (db *userConnection) DeleteUser(id uint64) error {

	sql := fmt.Sprintf("delete from users where id = %d", id)
	if err := db.connection.Exec(sql); err != nil {
		return err.Error
	}

	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user model.User
	return db.connection.Where("email = ?", email).Take(&user)

}

func (db *userConnection) IsDuplicateName(name string) (tx *gorm.DB) {
	var user model.User
	return db.connection.Where("name = ?", name).Take(&user)
}
