package repository

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type BookSubCategoryRepository interface {
	CreateBookSubCategory(bookSubCategory model.BookSubCategory) (*model.BookSubCategory, error)
	IsDuplicateSubCategoryTitle(name string) (tx *gorm.DB)
	GetAllBookSubCategory(req *dto.BookSubCategoryGetRequest) ([]model.BookSubCategory, int64, error)
	UpdateBookSubCategory(subCategory model.BookSubCategory) (*model.BookSubCategory, error)
	DeleteBookSubCategory(id uint64) error
}

type bookSubCategoryConnection struct {
	connection *gorm.DB
}

func NewBookSubCategoryRepository(db *gorm.DB) BookSubCategoryRepository {
	return &bookSubCategoryConnection{
		connection: db,
	}
}

func (db *bookSubCategoryConnection) CreateBookSubCategory(bookSubCategory model.BookSubCategory) (*model.BookSubCategory, error) {
	err := db.connection.Save(&bookSubCategory).Error
	if err != nil {
		fmt.Println("------------Have error in create book sub category ------------")
		return nil, err
	}

	return &bookSubCategory, nil
}

func (db *bookSubCategoryConnection) IsDuplicateSubCategoryTitle(name string) (tx *gorm.DB) {
	var subCategory model.BookSubCategory
	return db.connection.Where("title = ?", name).Take(&subCategory)
}

func (db *bookSubCategoryConnection) GetAllBookSubCategory(req *dto.BookSubCategoryGetRequest) ([]model.BookSubCategory, int64, error) {

	var subCategories []model.BookSubCategory
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

	if req.SubCategoryName != "" {
		filter += fmt.Sprintf(" and title LIKE \"%s%s%s\"", "%", req.SubCategoryName, "%")
	}

	sql := fmt.Sprintf("select * from book_sub_categories %s order by created_at desc limit %v offset %v", filter, pageSize, offset)
	res := db.connection.Raw(sql).Scan(&subCategories)

	countQuery := fmt.Sprintf("select count(1) from book_sub_categories %s", filter)
	if err := db.connection.Raw(countQuery).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	if res.Error == nil {
		return subCategories, total, nil
	}

	return nil, 0, nil
}

func (db *bookSubCategoryConnection) UpdateBookSubCategory(subCategory model.BookSubCategory) (*model.BookSubCategory, error) {
	err := db.connection.Model(&subCategory).Where("id = ?", subCategory.ID).Updates(model.BookSubCategory{
		CategoryID:      subCategory.CategoryID,
		SubCategoryName: subCategory.SubCategoryName,
		Description:     subCategory.Description,
	}).Error
	if err != nil {
		fmt.Println("Error at update book category repository----")
		return nil, err
	}
	return &subCategory, nil
}

func (db *bookSubCategoryConnection) DeleteBookSubCategory(id uint64) error {

	mydb := db.connection.Model(&model.BookSubCategory{})
	mydb = mydb.Where(fmt.Sprintf("id  = %d", id))
	if err := mydb.Delete(&model.BookSubCategory{}).Error; err != nil {
		return err
	}
	return nil
}
