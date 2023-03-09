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

type BookCategoryController interface {
	CreateBookCategory(ctx *gin.Context)
	GetAllBookCategory(ctx *gin.Context)
	UpdateBookCategory(ctx *gin.Context)
	DeleteBookCategory(ctx *gin.Context)
}

type bookCategoryController struct {
	bookCategoryService service.BookCategoryService
}

func NewBookCategoryControlle(bookCategoryService service.BookCategoryService) BookCategoryController {
	return &bookCategoryController{
		bookCategoryService: bookCategoryService,
	}
}

type ResponseCategoryListData struct {
	List  []model.BookCategory `json:"list"`
	Total int64                `json:"total"`
}

func (c *bookCategoryController) CreateBookCategory(ctx *gin.Context) {
	var createDto dto.CreateBookCategoryDTO
	errDto := ctx.ShouldBind(&createDto)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	_, err := c.bookCategoryService.CreateBookCategory(createDto)
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

	response := helper.ResponseData(0, "success", nil)
	ctx.JSON(http.StatusOK, response)
}

func (c *bookCategoryController) GetAllBookCategory(ctx *gin.Context) {
	req := &dto.BookCategoryGetRequest{}

	if err := ctx.ShouldBind(&req); err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	result, count, err := c.bookCategoryService.GetAllBookCategory(req)

	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	if count == 0 {
		response := helper.ResponseErrorData(500, "Request not found")
		ctx.JSON(http.StatusOK, response)
		return
	}

	var responseList ResponseCategoryListData
	responseList.List = result
	responseList.Total = count

	response := helper.ResponseData(0, "success", responseList)
	ctx.JSON(http.StatusOK, response)

}

func (c *bookCategoryController) UpdateBookCategory(ctx *gin.Context) {
	var updateCategoryDto dto.UpdateBookCategoryDTO
	errDTO := ctx.ShouldBind(&updateCategoryDto)
	if errDTO != nil {
		response := helper.ResponseErrorData(500, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	_, err := c.bookCategoryService.UpdateBookCateogry(updateCategoryDto)
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
	response := helper.ResponseData(0, "success", nil)
	ctx.JSON(http.StatusOK, response)
}

func (c *bookCategoryController) DeleteBookCategory(ctx *gin.Context) {
	var deleteDto dto.DeleteByIdDTO
	errDto := ctx.ShouldBind(&deleteDto)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	err := c.bookCategoryService.DeleteBookCategory(deleteDto.ID)
	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.ResponseData(0, "success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
