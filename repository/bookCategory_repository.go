package repository

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type BookCategoryRepository interface {
	CreateBookCategory(bookcategory model.BookCategory) (*model.BookCategory, error)
	IsDuplicateCategoryTitle(title string) (tx *gorm.DB)
	GetAllBookCategory(req *dto.BookCategoryGetRequest) ([]model.BookCategory, int64, error)
	UpdateBookCategory(category model.BookCategory) (*model.BookCategory, error)
	DeleteBookCategory(id uint64) error
}

type bookCategoryConnection struct {
	connection *gorm.DB
}

func NewBookCategoryRepository(db *gorm.DB) BookCategoryRepository {
	return &bookCategoryConnection{
		connection: db,
	}
}

func (db *bookCategoryConnection) CreateBookCategory(bookcategory model.BookCategory) (*model.BookCategory, error) {
	err := db.connection.Save(&bookcategory).Error
	if err != nil {
		fmt.Println("------------Have error in create book category ------------")
		return nil, err
	}

	return &bookcategory, nil
}

func (db *bookCategoryConnection) IsDuplicateCategoryTitle(title string) (tx *gorm.DB) {
	var category model.BookCategory
	return db.connection.Where("title = ?", title).Take(&category)
}

func (db *bookCategoryConnection) GetAllBookCategory(req *dto.BookCategoryGetRequest) ([]model.BookCategory, int64, error) {

	var categories []model.BookCategory
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

	// var filter string

	filter := " where deleted_at IS NULL"

	if req.ID != 0 {
		filter += fmt.Sprintf(" and id = %d", req.ID)
	}

	if req.CategoryName != "" {
		filter += fmt.Sprintf(" and category_name LIKE \"%s%s%s\"", "%", req.CategoryName, "%")
	}

	sql := fmt.Sprintf("select * from book_categories %s order by created_at desc limit %v offset %v", filter, pageSize, offset)
	res := db.connection.Raw(sql).Scan(&categories)

	countQuery := fmt.Sprintf("select count(1) from book_categories %s", filter)
	if err := db.connection.Raw(countQuery).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	if res.Error == nil {
		return categories, total, nil
	}

	return nil, 0, nil
}

func (db *bookCategoryConnection) UpdateBookCategory(category model.BookCategory) (*model.BookCategory, error) {
	err := db.connection.Model(&category).Where("id = ?", category.ID).Updates(model.BookCategory{
		CategoryName: category.CategoryName,
		// Description: category.Description,
	}).Error
	if err != nil {
		fmt.Println("Error at update book category repository----")
		return nil, err
	}
	return &category, nil
}

func (db *bookCategoryConnection) DeleteBookCategory(id uint64) error {

	mydb := db.connection.Model(&model.BookCategory{})
	mydb = mydb.Where(fmt.Sprintf("id  = %d", id))
	if err := mydb.Delete(&model.BookCategory{}).Error; err != nil {
		return err
	}
	return nil
}
