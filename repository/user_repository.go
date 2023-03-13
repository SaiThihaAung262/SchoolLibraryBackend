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
	GetUserDashBoard() (*dto.DashboardResponse, error)
	GetMostBorrowLog(req *dto.ReqMostBorrowData) ([]dto.MostBorrowBookData, uint64, error)
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

	sql := fmt.Sprintf("select * from users %s order by created_at desc limit %v offset %v", filter, pageSize, offset)
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

func (db *userConnection) GetUserDashBoard() (*dto.DashboardResponse, error) {
	dashboardData := &dto.DashboardResponse{}

	var adminCount int64
	adminCountSQL := "select count(1) from users"
	if err := db.connection.Raw(adminCountSQL).Scan(&adminCount).Error; err != nil {
		return nil, err
	}

	var teacherCount int64
	teacherCountSQL := "select count(1) from teachers where deleted_at IS NULL"
	if err := db.connection.Raw(teacherCountSQL).Scan(&teacherCount).Error; err != nil {
		return nil, err
	}

	var studentCount int64
	studentCountSQL := "select count(1) from students where deleted_at IS NULL"
	if err := db.connection.Raw(studentCountSQL).Scan(&studentCount).Error; err != nil {
		return nil, err
	}

	var bookCount int64
	bookCountSQL := "select count(1) from books where deleted_at IS NULL"
	if err := db.connection.Raw(bookCountSQL).Scan(&bookCount).Error; err != nil {
		return nil, err
	}

	var categoryCount int64
	categoryCountSQL := "select count(1) from book_categories where deleted_at IS NULL"
	if err := db.connection.Raw(categoryCountSQL).Scan(&categoryCount).Error; err != nil {
		return nil, err
	}

	var totalBorrowCount int64
	totalBorrowCountSQL := "select count(1) from borrows where deleted_at IS NULL"
	if err := db.connection.Raw(totalBorrowCountSQL).Scan(&totalBorrowCount).Error; err != nil {
		return nil, err
	}

	var underBorrowing int64
	underBorrowingCountSQL := "select count(1) from borrows where status = 1 and deleted_at IS NULL"
	if err := db.connection.Raw(underBorrowingCountSQL).Scan(&underBorrowing).Error; err != nil {
		return nil, err
	}

	var haveReturned int64
	haveReturnedCountSQL := "select count(1) from borrows where status = 2 and deleted_at IS NULL"
	if err := db.connection.Raw(haveReturnedCountSQL).Scan(&haveReturned).Error; err != nil {
		return nil, err
	}

	dashboardData.TotalAdmin = adminCount
	dashboardData.TotalTeacher = teacherCount
	dashboardData.TotalStudent = studentCount
	dashboardData.TotalBook = bookCount
	dashboardData.TotalCategory = categoryCount
	dashboardData.TotalBorrow = totalBorrowCount
	dashboardData.UnderBorrow = underBorrowing
	dashboardData.HaveReturned = haveReturned
	return dashboardData, nil
}

func (db *userConnection) GetMostBorrowLog(req *dto.ReqMostBorrowData) ([]dto.MostBorrowBookData, uint64, error) {
	var mostBorrowBooks []dto.MostBorrowBookData
	var total uint64

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

	// filter := " where type = 2"
	sql := fmt.Sprintf("SELECT book_id, book_uuid, book_title, COUNT(1) AS 'borrow_count' FROM borrow_logs WHERE type = 2 GROUP BY book_uuid ORDER BY borrow_count DESC limit %v offset %v", pageSize, offset)
	res := db.connection.Raw(sql).Scan(&mostBorrowBooks)
	if res.Error != nil {
		return nil, 0, res.Error
	}

	countQuery := "SELECT COUNT(*) as 'total_count' FROM (SELECT COUNT(1) AS 'borrow_count' FROM borrow_logs WHERE type = 2 GROUP BY book_uuid) as t1"
	if err := db.connection.Raw(countQuery).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	return mostBorrowBooks, total, nil
}
