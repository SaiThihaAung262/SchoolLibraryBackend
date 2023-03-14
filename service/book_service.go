package service

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/model"
	"MyGO.com/m/repository"
	"github.com/mashingan/smapping"
)

type BookService interface {
	CreateBook(book dto.CreateBookDTO) (*model.Book, error)
	IsBookTitleDuplicate(title string) bool
	GetAllBooks(req *dto.BookGetRequest) ([]model.Book, int64, error)
	UpdateBook(book dto.UpdateBookDTO) (*model.Book, error)
	DeleteBook(id uint64) error
	GetBookByUUID(uuid string) (*model.Book, error)
	GetBookByUUIDAndDate(req *dto.ReqBorrowCountByBookUUIDAndDateDto) (*model.Book, error)
	UpdateBookBorrowQTY(id uint64, borrowQty uint64) error
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepository,
	}

}

func (service bookService) CreateBook(book dto.CreateBookDTO) (*model.Book, error) {
	bookToCreate := model.Book{}

	err := smapping.FillStruct(&bookToCreate, smapping.MapFields(&book))
	if err != nil {
		fmt.Println("Here have error in Create book service")
	}
	bookToCreate.UUID = helper.GenerateUUID()
	res, errRepo := service.bookRepository.CreateBook(bookToCreate)
	if errRepo != nil {
		return nil, errRepo
	}
	return res, nil

}

func (service bookService) IsBookTitleDuplicate(title string) bool {
	res := service.bookRepository.IsBookTitleDuplicate(title)
	return (res.Error == nil)
}

func (service bookService) GetAllBooks(req *dto.BookGetRequest) ([]model.Book, int64, error) {
	books, total, err := service.bookRepository.GetAllBooks(req)
	if err != nil {
		return nil, 0, err
	}
	return books, total, err
}

func (service bookService) UpdateBook(book dto.UpdateBookDTO) (*model.Book, error) {
	bookToUpdate := model.Book{}

	err := smapping.FillStruct(&bookToUpdate, smapping.MapFields(&book))
	if err != nil {
		fmt.Println("-------- Here have error in book service update -------")
	}

	res, errRepo := service.bookRepository.UpdateBook(bookToUpdate)
	if errRepo != nil {
		return nil, errRepo
	}
	return res, nil
}

func (service bookService) DeleteBook(id uint64) error {
	err := service.bookRepository.DeleteBook(id)
	return err
}

func (service bookService) GetBookByUUID(uuid string) (*model.Book, error) {
	book, err := service.bookRepository.GetBookByUUID(uuid)
	return book, err
}

func (service bookService) GetBookByUUIDAndDate(req *dto.ReqBorrowCountByBookUUIDAndDateDto) (*model.Book, error) {
	return service.bookRepository.GetBookByUUIDAndDate(req)
}

func (service bookService) UpdateBookBorrowQTY(id uint64, borrowQty uint64) error {
	return service.bookRepository.UpdateBookBorrowQTY(id, borrowQty)
}
