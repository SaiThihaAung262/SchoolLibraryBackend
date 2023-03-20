package repository

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type StaffRepository interface {
	InsertStaff(staff model.Staff) (*model.Staff, error)
	GetAllStaff(req *dto.StaffGetRequest) ([]model.Staff, int64, error)
	UpdateStaff(staff model.Staff) (*model.Staff, error)
	DeleteStaff(id uint64) error
	GetStaffByUUID(uuid string) (*model.Staff, error)
	VerifyLogin(name string) interface{}
	ChangePassword(id uint64, password string) error
}

type staffConnection struct {
	connection *gorm.DB
}

func NeweStaffRepository(db *gorm.DB) StaffRepository {
	return &staffConnection{
		connection: db,
	}
}

func (db *staffConnection) InsertStaff(staff model.Staff) (*model.Staff, error) {
	err := db.connection.Save(&staff).Error
	if err != nil {
		fmt.Println("------------Here is error in staff repository--------------", err)
		return nil, err
	}
	return &staff, nil
}

func (db *staffConnection) GetAllStaff(req *dto.StaffGetRequest) ([]model.Staff, int64, error) {
	var staffs []model.Staff
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

	filter := " where deleted_at IS NULL"

	if req.UUID != "" {
		filter += fmt.Sprintf(" and uuid = %s", req.UUID)
	}

	if req.Department != 0 {
		filter += fmt.Sprintf(" and department = %d", req.Department)
	}

	if req.Name != "" {
		filter += fmt.Sprintf(" and name LIKE \"%s%s%s\"", "%", req.Name, "%")
	}

	if req.Email != "" {
		filter += fmt.Sprintf(" and email LIKE \"%s%s%s\"", "%", req.Email, "%")

	}

	countQuery := fmt.Sprintf("select count(1) from staffs %s", filter)
	if err := db.connection.Raw(countQuery).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	sql := fmt.Sprintf("select * from staffs %s order by created_at desc limit %v offset %v", filter, pageSize, offset)
	res := db.connection.Raw(sql).Scan(&staffs)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	return staffs, total, nil
}

func (db *staffConnection) UpdateStaff(staff model.Staff) (*model.Staff, error) {
	err := db.connection.Model(&staff).Where("id = ?", staff.ID).Updates(model.Staff{
		UUID:       staff.UUID,
		Name:       staff.Name,
		Email:      staff.Email,
		Password:   staff.Password,
		Department: staff.Department,
	}).Error
	if err != nil {
		fmt.Println("Error at update staff repository----")
		return nil, err
	}
	return &staff, nil
}

func (db *staffConnection) DeleteStaff(id uint64) error {

	mydb := db.connection.Model(&model.Staff{})
	mydb = mydb.Where(fmt.Sprintf("id  = %d", id))
	if err := mydb.Delete(&model.Staff{}).Error; err != nil {
		return err
	}
	return nil
}

func (db *staffConnection) GetStaffByUUID(uuid string) (*model.Staff, error) {
	staff := &model.Staff{}
	myDb := db.connection.Model(&model.Staff{})
	myDb = myDb.Where("uuid = ?", uuid)
	if err := myDb.First(&staff).Error; err != nil {
		return nil, err
	}
	return staff, nil
}

func (db *staffConnection) VerifyLogin(name string) interface{} {
	var staff model.Staff

	res := db.connection.Where("email = ?", name).Take(&staff)

	if res.Error == nil {
		return staff
	}
	return nil
}

func (db *staffConnection) ChangePassword(id uint64, password string) error {
	updateStaff := model.Staff{}

	fmt.Println("----------Hre is punishent ID--------", id)

	err := db.connection.Model(&updateStaff).Where("id = ?", id).Select("password").Updates(model.Staff{
		Password: password,
	}).Error

	if err != nil {
		fmt.Println("----Here have error in update borrow qty book repo -----")
		return err
	}
	return nil
}
