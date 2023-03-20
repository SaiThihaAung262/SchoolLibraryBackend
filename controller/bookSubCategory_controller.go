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

type BookSubCategoryController interface {
	CreateBookSubCategory(ctx *gin.Context)
	GetAllBookSubCategory(ctx *gin.Context)
	UpdateBookSubCategory(ctx *gin.Context)
	DeleteBookSubCategory(ctx *gin.Context)
}

type bookSubCategoryController struct {
	bookSubCategoryService service.BookSubCategoryService
}

func NewBookSubCategoryControlle(bookSubCategoryService service.BookSubCategoryService) BookSubCategoryController {
	return &bookSubCategoryController{
		bookSubCategoryService: bookSubCategoryService,
	}
}

type ResponseSubCategoryListData struct {
	List  []model.BookSubCategory `json:"list"`
	Total int64                   `json:"total"`
}

func (c *bookSubCategoryController) CreateBookSubCategory(ctx *gin.Context) {
	var createDto dto.CreateBookSubCategoryDTO
	errDto := ctx.ShouldBind(&createDto)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	_, err := c.bookSubCategoryService.CreateBookSubCategory(createDto)
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

func (c *bookSubCategoryController) GetAllBookSubCategory(ctx *gin.Context) {
	req := &dto.BookSubCategoryGetRequest{}

	if err := ctx.ShouldBind(&req); err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	result, count, err := c.bookSubCategoryService.GetAllBookSubCategory(req)

	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	var responseList ResponseSubCategoryListData
	responseList.List = result
	responseList.Total = count

	response := helper.ResponseData(0, "success", responseList)
	ctx.JSON(http.StatusOK, response)

}

func (c *bookSubCategoryController) UpdateBookSubCategory(ctx *gin.Context) {
	var updateCategoryDto dto.UpdateBookSubCategoryDTO
	errDTO := ctx.ShouldBind(&updateCategoryDto)
	if errDTO != nil {
		response := helper.ResponseErrorData(500, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	res, err := c.bookSubCategoryService.UpdateBookSubCateogry(updateCategoryDto)
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
	response := helper.ResponseData(0, "success", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *bookSubCategoryController) DeleteBookSubCategory(ctx *gin.Context) {
	var deleteDto dto.DeleteByIdDTO
	errDto := ctx.ShouldBind(&deleteDto)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	err := c.bookSubCategoryService.DeleteBookSubCategory(deleteDto.ID)
	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.ResponseData(0, "success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
