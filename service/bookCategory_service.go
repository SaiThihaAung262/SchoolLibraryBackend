package service

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"MyGO.com/m/repository"
	"github.com/mashingan/smapping"
)

type BookCategoryService interface {
	CreateBookCategory(book_category dto.CreateBookCategoryDTO) (*model.BookCategory, error)
	IsDuplicateCategoryTitle(title string) bool
	GetAllBookCategory(req *dto.BookCategoryGetRequest) ([]model.BookCategory, int64, error)
	UpdateBookCateogry(category dto.UpdateBookCategoryDTO) (*model.BookCategory, error)
	DeleteBookCategory(id uint64) error
}

type bookCategoryService struct {
	bookCategoryRepo repository.BookCategoryRepository
}

func NewBookCategoryService(bookCategoryRepo repository.BookCategoryRepository) BookCategoryService {
	return &bookCategoryService{
		bookCategoryRepo: bookCategoryRepo,
	}
}

func (service bookCategoryService) CreateBookCategory(book_category dto.CreateBookCategoryDTO) (*model.BookCategory, error) {
	categoryToCreate := model.BookCategory{}
	err := smapping.FillStruct(&categoryToCreate, smapping.MapFields(&book_category))
	if err != nil {
		fmt.Println("-----Here is error in category service -----", err)
	}
	res, errRepo := service.bookCategoryRepo.CreateBookCategory(categoryToCreate)
	if errRepo != nil {
		return nil, errRepo
	}
	return res, nil
}

func (service bookCategoryService) IsDuplicateCategoryTitle(title string) bool {
	res := service.bookCategoryRepo.IsDuplicateCategoryTitle(title)
	fmt.Println("____________res____________", res.Error)

	return (res.Error == nil)
}

func (service bookCategoryService) GetAllBookCategory(req *dto.BookCategoryGetRequest) ([]model.BookCategory, int64, error) {

	categories, count, err := service.bookCategoryRepo.GetAllBookCategory(req)
	if err != nil {
		return nil, 0, err
	}
	return categories, count, err
}

func (service bookCategoryService) UpdateBookCateogry(category dto.UpdateBookCategoryDTO) (*model.BookCategory, error) {
	categoryToUpdate := model.BookCategory{}
	err := smapping.FillStruct(&categoryToUpdate, smapping.MapFields(&category))
	if err != nil {
		fmt.Println("------Have error in update bookcategory servcie ------", err.Error())
	}

	res, errRepo := service.bookCategoryRepo.UpdateBookCategory(categoryToUpdate)
	if errRepo != nil {
		return nil, errRepo
	}

	return res, nil

}

func (service bookCategoryService) DeleteBookCategory(id uint64) error {
	err := service.bookCategoryRepo.DeleteBookCategory(id)
	return err
}
