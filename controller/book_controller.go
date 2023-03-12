package controller

import (
	"net/http"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/model"
	"MyGO.com/m/repository/criteria"
	"MyGO.com/m/service"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	CreateBook(ctx *gin.Context)
	GetAllBooks(ctx *gin.Context)
	UpdateBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
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

	createdBook, err := c.bookService.CreateBook(bookToCreate)
	if err != nil {
		if criteria.IsErrNotFound(err) {
			response := helper.ResponseErrorData(500, "Cannot find")
			ctx.JSON(http.StatusOK, response)
			return
		}
		if criteria.IsDuplicate(err) {
			response := helper.ResponseErrorData(528, "This title already have!")
			ctx.JSON(http.StatusOK, response)
			return
		}
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}
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
		var responseData ResponseBookLists
		responseData.Books = result
		responseData.Total = count
		response := helper.ResponseData(0, "success", responseData)
		ctx.JSON(http.StatusOK, response)
		return
	}

	var responseData ResponseBookLists
	responseData.Books = result
	responseData.Total = count

	response := helper.ResponseData(0, "success", responseData)

	ctx.JSON(http.StatusOK, response)
}

func (c bookController) UpdateBook(ctx *gin.Context) {
	var bookToUpdate dto.UpdateBookDTO
	errDTO := ctx.ShouldBind(&bookToUpdate)
	if errDTO != nil {
		respons := helper.ResponseErrorData(500, errDTO.Error())
		ctx.JSON(http.StatusOK, respons)
		return
	}

	// isDuplicate := c.bookService.IsBookTitleDuplicate(bookToUpdate.Title)
	// if isDuplicate {
	// 	response := helper.ResponseErrorData(500, "The title with this book is already exit !")
	// 	ctx.JSON(http.StatusOK, response)
	// 	return
	// }
	updatedBook, err := c.bookService.UpdateBook(bookToUpdate)
	if err != nil {
		if criteria.IsErrNotFound(err) {
			response := helper.ResponseErrorData(500, "Cannot find")
			ctx.JSON(http.StatusOK, response)
			return
		}
		if criteria.IsDuplicate(err) {
			response := helper.ResponseErrorData(528, "This title already have!")
			ctx.JSON(http.StatusOK, response)
			return
		}
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.ResponseData(0, "success", updatedBook)
	ctx.JSON(http.StatusOK, response)
}

func (c bookController) DeleteBook(ctx *gin.Context) {
	var deleteId dto.DeleteByIdDTO
	errDto := ctx.ShouldBind(&deleteId)

	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	if err := c.bookService.DeleteBook(deleteId.ID); err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.ResponseData(0, "success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
