package controller

import (
	"net/http"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/model"
	"MyGO.com/m/service"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	CreateBook(ctx *gin.Context)
	GetAllBooks(ctx *gin.Context)
}

type bookController struct {
	bookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &bookController{
		bookService: bookService,
	}
}

func (c bookController) CreateBook(ctx *gin.Context) {
	var bookToCreate dto.CreateBookDTO

	errDto := ctx.ShouldBind(&bookToCreate)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	isDuplicateTitle := c.bookService.IsBookTitleDuplicate(bookToCreate.Title)
	if isDuplicateTitle {
		response := helper.ResponseErrorData(500, "The title with this book is already exit !")
		ctx.JSON(http.StatusOK, response)
		return
	}

	createdBook := c.bookService.CreateBook(bookToCreate)
	response := helper.ResponseData(0, "success", createdBook)
	ctx.JSON(http.StatusOK, response)

}

type ResponseBookLists struct {
	Books []model.Book `json:"list"`
	Total int64        `json:"total"`
}

func (c bookController) GetAllBooks(ctx *gin.Context) {
	var req dto.BookGetRequest
	if errReqDto := ctx.ShouldBind(&req); errReqDto != nil {
		response := helper.ResponseErrorData(500, errReqDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	result, count, err := c.bookService.GetAllBooks(&req)

	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	if count == 0 {
		response := helper.ResponseErrorData(500, "No data")
		ctx.JSON(http.StatusOK, response)
		return
	}

	var responseData ResponseBookLists
	responseData.Books = result
	responseData.Total = count

	response := helper.ResponseData(0, "success", responseData)

	ctx.JSON(http.StatusOK, response)
}
