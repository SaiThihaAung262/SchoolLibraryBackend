package repository

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book model.Book) (*model.Book, error)
	IsBookTitleDuplicate(title string) (tx *gorm.DB)
	GetAllBooks(req *dto.BookGetRequest) ([]model.Book, int64, error)
	UpdateBook(book model.Book) (*model.Book, error)
	DeleteBook(id uint64) error
}

type bookConnection struct {
	connection *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookConnection{
		connection: db,
	}
}

func (db *bookConnection) CreateBook(book model.Book) (*model.Book, error) {
	err := db.connection.Save(&book).Error
	if err != nil {
		fmt.Println("Here have error in create book repo")
		return nil, err
	}
	return &book, nil
}

func (db *bookConnection) IsBookTitleDuplicate(title string) (tx *gorm.DB) {
	var book model.Book
	return db.connection.Where("title = ?", title).Take(&book)
}

func (db *bookConnection) GetAllBooks(req *dto.BookGetRequest) ([]model.Book, int64, error) {

	var books []model.Book
	var offset uint64
	var pageSize uint64
	var total int64

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

	if req.ID != 0 {
		filter += fmt.Sprintf(" and id = %d", req.ID)

	}

	countQuery := fmt.Sprintf("select count(1) from books %s", filter)
	if err := db.connection.Raw(countQuery).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	sql := fmt.Sprintf("select * from books %s limit %v offset %v", filter, pageSize, offset)
	res := db.connection.Raw(sql).Scan(&books)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	return books, total, nil

}

func (db *bookConnection) UpdateBook(book model.Book) (*model.Book, error) {
	err := db.connection.Model(&book).Where("id = ?", book.ID).Updates(model.Book{
		Title:      book.Title,
		CategoryID: book.CategoryID,
		Author:     book.Author,
		Summary:    book.Summary,
		Status:     book.Status,
	}).Error
	if err != nil {
		fmt.Println("----Here have error in update book repo -----")
		return nil, err

	}
	return &book, nil
}

func (db *bookConnection) DeleteBook(id uint64) error {

	mydb := db.connection.Model(&model.Book{})
	mydb = mydb.Where(fmt.Sprintf("id  = %d", id))
	if err := mydb.Delete(&model.Book{}).Error; err != nil {
		return err
	}
	return nil
}
