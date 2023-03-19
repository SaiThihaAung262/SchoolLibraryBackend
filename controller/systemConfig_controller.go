package controller

import (
	"net/http"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/repository/criteria"
	"MyGO.com/m/service"
	"github.com/gin-gonic/gin"
)

type SystemConfigController interface {
	CreateSystemConfg(ctx *gin.Context)
	GetSystemConfig(ctx *gin.Context)
	UpdateSystemConfig(ctx *gin.Context)
	DeleteSystemConfig(ctx *gin.Context)
}

type systemConfigController struct {
	service service.SystemConfigService
}

func NewSystemConfigController(service service.SystemConfigService) SystemConfigController {
	return &systemConfigController{
		service: service,
	}
}

type SystemConfigResp struct {
	ID                    uint64 `json:"id"`
	TeacherCanBorrowCount uint64 `json:"teacher_can_borrow_count"`
	StudentCanBorrowCount uint64 `json:"student_can_borrow_count"`
	StaffCanBorrowCount   uint64 `json:"staff_can_borrow_count"`
	TeacherPunishAmt      uint64 `json:"teacher_punishment_amt"`
	StudentPunishAmt      uint64 `json:"student_punishment_amt"`
	TeacherCanBorrowDay   uint64 `json:"teacher_can_borrow_day"`
	StudentCanBorrowDay   uint64 `json:"student_can_borrow_day"`
	StaffCanBorrowDay     uint64 `json:"staff_can_borrow_day"`
}

func (c systemConfigController) CreateSystemConfg(ctx *gin.Context) {
	var createDto dto.CreateSystemConfigDto
	errDto := ctx.ShouldBind(&createDto)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	_, err := c.service.InsertSystemConfig(createDto)
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

func (c systemConfigController) GetSystemConfig(ctx *gin.Context) {
	result, err := c.service.GetSystemConfig()

	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	responseData := SystemConfigResp{
		ID:                    uint64(result.ID),
		TeacherCanBorrowCount: result.TeacherCanBorrowCount,
		StudentCanBorrowCount: result.StudentCanBorrowCount,
		StaffCanBorrowCount:   result.StaffCanBorrowCount,
		TeacherPunishAmt:      result.TeacherPunishAmt,
		StudentPunishAmt:      result.StudentPunishAmt,
		TeacherCanBorrowDay:   result.TeacherCanBorrowDay,
		StudentCanBorrowDay:   result.StudentCanBorrowDay,
		StaffCanBorrowDay:     result.StaffCanBorrowDay,
	}

	response := helper.ResponseData(0, "success", responseData)
	ctx.JSON(http.StatusOK, response)
}

func (c systemConfigController) UpdateSystemConfig(ctx *gin.Context) {
	var updateDto dto.UpdateSystemConfigDto
	errDTO := ctx.ShouldBind(&updateDto)
	if errDTO != nil {
		response := helper.ResponseErrorData(500, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	res, err := c.service.UpdateSystemConfig(updateDto)
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

func (c systemConfigController) DeleteSystemConfig(ctx *gin.Context) {
	var deleteDto dto.DeleteByIdDTO
	errDto := ctx.ShouldBind(&deleteDto)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	err := c.service.DeleteSystemConfig(deleteDto.ID)
	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.ResponseData(0, "success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
