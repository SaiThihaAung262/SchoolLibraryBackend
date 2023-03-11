package controller

import (
	"net/http"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/service"
	"github.com/gin-gonic/gin"
)

type BorrowController interface {
	CreateBorrow(ctx *gin.Context)
}

type borrowController struct {
	borrowService service.Borrowservice
}

func NewBorrowController(borrowService service.Borrowservice) BorrowController {
	return &borrowController{
		borrowService: borrowService,
	}
}

func (c borrowController) CreateBorrow(ctx *gin.Context) {
	var createDto dto.CreateBorrowDTO
	errDto := ctx.ShouldBind(&createDto)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	err := c.borrowService.CreateBorrow(createDto)
	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.ResponseData(0, "success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)

}
