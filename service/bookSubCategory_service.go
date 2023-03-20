package service

import (
	"fmt"

	"MyGO.com/m/dto"
	"MyGO.com/m/model"
	"MyGO.com/m/repository"
	"github.com/mashingan/smapping"
)

type BookSubCategoryService interface {
	CreateBookSubCategory(book_sub_category dto.CreateBookSubCategoryDTO) (*model.BookSubCategory, error)
	GetAllBookSubCategory(req *dto.BookSubCategoryGetRequest) ([]model.BookSubCategory, int64, error)
	UpdateBookSubCateogry(category dto.UpdateBookSubCategoryDTO) (*model.BookSubCategory, error)
	DeleteBookSubCategory(id uint64) error
}

type bookSubCategoryService struct {
	bookSubCategoryRepo repository.BookSubCategoryRepository
}

func NewBookSubCategoryService(bookSubCategoryRepo repository.BookSubCategoryRepository) BookSubCategoryService {
	return &bookSubCategoryService{
		bookSubCategoryRepo: bookSubCategoryRepo,
	}
}

func (service bookSubCategoryService) CreateBookSubCategory(book_sub_category dto.CreateBookSubCategoryDTO) (*model.BookSubCategory, error) {
	categoryToCreate := model.BookSubCategory{}
	err := smapping.FillStruct(&categoryToCreate, smapping.MapFields(&book_sub_category))
	if err != nil {
		fmt.Println("-----Here is error in category service -----", err)
	}
	res, errRepo := service.bookSubCategoryRepo.CreateBookSubCategory(categoryToCreate)
	if errRepo != nil {
		return nil, errRepo
	}
	return res, nil
}

func (service bookSubCategoryService) GetAllBookSubCategory(req *dto.BookSubCategoryGetRequest) ([]model.BookSubCategory, int64, error) {

	categories, count, err := service.bookSubCategoryRepo.GetAllBookSubCategory(req)
	if err != nil {
		return nil, 0, err
	}
	return categories, count, err
}

func (service bookSubCategoryService) UpdateBookSubCateogry(subCategory dto.UpdateBookSubCategoryDTO) (*model.BookSubCategory, error) {
	categoryToUpdate := model.BookSubCategory{}
	err := smapping.FillStruct(&categoryToUpdate, smapping.MapFields(&subCategory))
	if err != nil {
		fmt.Println("------Have error in update bookcategory servcie ------", err.Error())
	}

	res, errRepo := service.bookSubCategoryRepo.UpdateBookSubCategory(categoryToUpdate)
	if errRepo != nil {
		return nil, errRepo
	}
	return res, nil
}

func (service bookSubCategoryService) DeleteBookSubCategory(id uint64) error {
	err := service.bookSubCategoryRepo.DeleteBookSubCategory(id)
	return err
}
