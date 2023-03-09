package service

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"MyGO.com/m/repository"
	"github.com/mashingan/smapping"
)

type BookService interface {
	CreateBook(book dto.CreateBookDTO) model.Book
	IsBookTitleDuplicate(title string) bool
	GetAllBooks(req *dto.BookGetRequest) ([]model.Book, int64, error)
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepository,
	}

}

func (service bookService) CreateBook(book dto.CreateBookDTO) model.Book {
	bookToCreate := model.Book{}
	err := smapping.FillStruct(&bookToCreate, smapping.MapFields(book))
	if err != nil {
		fmt.Println("Here have error in Create book service")
	}
	res := service.bookRepository.CreateBook(bookToCreate)
	return res

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
