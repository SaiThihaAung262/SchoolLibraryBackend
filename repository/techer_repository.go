package repository

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type TeacherRepository interface {
	InsertTeacher(teacher model.Teacher) (*model.Teacher, error)
	GetAllTeachers(req *dto.TeacherGetRequest) ([]model.Teacher, int64, error)
	UpdateTeacher(teacher model.Teacher) (*model.Teacher, error)
	DeleteTeacher(id uint64) error
	GetTeacherByUUID(uuid string) (*model.Teacher, error)
	VerifyLogin(name string) interface{}
}

type teacherConnection struct {
	connection *gorm.DB
}

func NeweTeacherRepository(db *gorm.DB) TeacherRepository {
	return &teacherConnection{
		connection: db,
	}
}

func (db *teacherConnection) InsertTeacher(teacher model.Teacher) (*model.Teacher, error) {
	err := db.connection.Save(&teacher).Error
	if err != nil {
		fmt.Println("------------Here is error in teacher repository--------------", err)
		return nil, err
	}
	return &teacher, nil
}

func (db *teacherConnection) GetAllTeachers(req *dto.TeacherGetRequest) ([]model.Teacher, int64, error) {
	var teachers []model.Teacher
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

	countQuery := fmt.Sprintf("select count(1) from teachers %s", filter)
	if err := db.connection.Raw(countQuery).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	sql := fmt.Sprintf("select * from teachers %s order by created_at desc limit %v offset %v", filter, pageSize, offset)
	res := db.connection.Raw(sql).Scan(&teachers)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	return teachers, total, nil
}

func (db *teacherConnection) UpdateTeacher(teacher model.Teacher) (*model.Teacher, error) {
	err := db.connection.Model(&teacher).Where("id = ?", teacher.ID).Updates(model.Teacher{
		UUID:       teacher.UUID,
		Name:       teacher.Name,
		Email:      teacher.Email,
		Password:   teacher.Password,
		Department: teacher.Department,
	}).Error
	if err != nil {
		fmt.Println("Error at update teacher category repository----")
		return nil, err
	}
	return &teacher, nil
}

func (db *teacherConnection) DeleteTeacher(id uint64) error {

	mydb := db.connection.Model(&model.Teacher{})
	mydb = mydb.Where(fmt.Sprintf("id  = %d", id))
	if err := mydb.Delete(&model.Teacher{}).Error; err != nil {
		return err
	}
	return nil
}

func (db *teacherConnection) GetTeacherByUUID(uuid string) (*model.Teacher, error) {
	teacher := &model.Teacher{}
	myDb := db.connection.Model(&model.Teacher{})
	myDb = myDb.Where("uuid = ?", uuid)
	if err := myDb.First(&teacher).Error; err != nil {
		return nil, err
	}
	return teacher, nil
}

func (db *teacherConnection) VerifyLogin(name string) interface{} {
	var teacher model.Teacher

	res := db.connection.Where("email = ?", name).Take(&teacher)

	if res.Error == nil {
		return teacher
	}
	return nil
}
