package controller

import (
	"fmt"
	"net/http"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/repository/criteria"
	"MyGO.com/m/service"
	"github.com/gin-gonic/gin"
)

type PunishmentController interface {
	CreatePunishment(ctx *gin.Context)
	GetPunishmentData(ctx *gin.Context)
	UpdatePunishment(ctx *gin.Context)
	DeletePunishment(ctx *gin.Context)
}

type punsihmentController struct {
	service service.PunishmentService
}

func NewPunishmentController(service service.PunishmentService) PunishmentController {
	return &punsihmentController{
		service: service,
	}
}

func (c punsihmentController) CreatePunishment(ctx *gin.Context) {
	var createDto dto.CreatePunishmentDto
	errDto := ctx.ShouldBind(&createDto)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	_, err := c.service.InsertPunishment(createDto)
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

func (c punsihmentController) GetPunishmentData(ctx *gin.Context) {

	result, err := c.service.GetPunishmentData()

	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.ResponseData(0, "success", result)
	ctx.JSON(http.StatusOK, response)
}

func (c punsihmentController) UpdatePunishment(ctx *gin.Context) {
	var updateDto dto.UpdatePunishmentDto
	errDTO := ctx.ShouldBind(&updateDto)
	if errDTO != nil {
		response := helper.ResponseErrorData(500, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	fmt.Println("-------------herei s update id in controller----------", updateDto.ID)

	res, err := c.service.UpdatePunishment(updateDto)
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

func (c punsihmentController) DeletePunishment(ctx *gin.Context) {
	var deleteDto dto.DeleteByIdDTO
	errDto := ctx.ShouldBind(&deleteDto)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	err := c.service.DeletePunishment(deleteDto.ID)
	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.ResponseData(0, "success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
