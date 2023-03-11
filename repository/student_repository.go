package repository

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type SutudentRepository interface {
	InsertStudent(student model.Student) (*model.Student, error)
	GetAllStudents(req *dto.StudentGetRequest) ([]model.Student, int64, error)
	UpdateStudent(student model.Student) (*model.Student, error)
	DeleteStudent(id uint64) error
}

type studentConnection struct {
	connection *gorm.DB
}

func NewStudentRepository(db *gorm.DB) SutudentRepository {
	return &studentConnection{
		connection: db,
	}
}

func (db *studentConnection) InsertStudent(student model.Student) (*model.Student, error) {
	err := db.connection.Save(&student).Error
	if err != nil {
		fmt.Println("------------Here is error in student repository--------------", err)
		return nil, err
	}
	return &student, nil
}

func (db *studentConnection) GetAllStudents(req *dto.StudentGetRequest) ([]model.Student, int64, error) {
	var students []model.Student
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

	if req.RoleNo != "" {
		filter += fmt.Sprintf(" and role_no LIKE \"%s%s%s\"", "%", req.RoleNo, "%")
	}

	if req.Year != 0 {
		filter += fmt.Sprintf(" and year = %d", req.Year)
	}

	if req.Name != "" {
		filter += fmt.Sprintf(" and name LIKE \"%s%s%s\"", "%", req.Name, "%")
	}

	if req.Email != "" {
		filter += fmt.Sprintf(" and email LIKE \"%s%s%s\"", "%", req.Email, "%")

	}

	countQuery := fmt.Sprintf("select count(1) from students %s", filter)
	if err := db.connection.Raw(countQuery).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	sql := fmt.Sprintf("select * from students %s order by created_at desc limit %v offset %v", filter, pageSize, offset)
	res := db.connection.Raw(sql).Scan(&students)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	return students, total, nil
}

func (db *studentConnection) UpdateStudent(student model.Student) (*model.Student, error) {
	err := db.connection.Model(&student).Where("id = ?", student.ID).Updates(model.Student{
		UUID:     student.UUID,
		Name:     student.Name,
		Email:    student.Email,
		Password: student.Password,
		RoleNo:   student.RoleNo,
		Year:     student.Year,
	}).Error
	if err != nil {
		fmt.Println("Error at update student category repository----")
		return nil, err
	}
	return &student, nil
}

func (db *studentConnection) DeleteStudent(id uint64) error {

	mydb := db.connection.Model(&model.Student{})
	mydb = mydb.Where(fmt.Sprintf("id  = %d", id))
	if err := mydb.Delete(&model.Student{}).Error; err != nil {
		return err
	}
	return nil
}
