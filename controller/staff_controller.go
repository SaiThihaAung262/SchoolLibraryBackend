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

type StaffController interface {
	CreateStaff(ctx *gin.Context)
	GetAllStaff(ctx *gin.Context)
	UpdateStaff(ctx *gin.Context)
	DeleteStaff(ctx *gin.Context)
}

type staffController struct {
	staffService service.StaffService
	jwtService   service.JwtService
}

func NewStaffController(staffService service.StaffService, jwtService service.JwtService) StaffController {
	return &staffController{
		staffService: staffService,
		jwtService:   jwtService,
	}
}

func (c *staffController) CreateStaff(ctx *gin.Context) {
	var registerDTO dto.StaffRegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.ResponseErrorData(501, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	createUser, err := c.staffService.InsertStaff(registerDTO)
	if err != nil {
		if criteria.IsErrNotFound(err) {
			response := helper.ResponseErrorData(500, "Cannot find")
			ctx.JSON(http.StatusOK, response)
			return
		}
		if criteria.IsDuplicate(err) {
			response := helper.ResponseErrorData(528, "Email is already exist !")
			ctx.JSON(http.StatusOK, response)
			return
		}
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.ResponseData(0, "Success", createUser)
	ctx.JSON(http.StatusOK, response)
}

type ResponseStaffData struct {
	List  []model.Staff `json:"list"`
	Total int64         `json:"total"`
}

func (c *staffController) GetAllStaff(ctx *gin.Context) {
	req := &dto.StaffGetRequest{}

	if err := ctx.ShouldBind(&req); err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	result, count, err := c.staffService.GetAllStaff(req)

	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	if count == 0 {
		var staffs []model.Staff
		var responseList ResponseStaffData
		responseList.List = staffs
		responseList.Total = count

		response := helper.ResponseData(0, "success", responseList)
		ctx.JSON(http.StatusOK, response)
		return
	}

	var responseList ResponseStaffData
	responseList.List = result
	responseList.Total = count

	response := helper.ResponseData(0, "success", responseList)
	ctx.JSON(http.StatusOK, response)

}

func (c *staffController) UpdateStaff(ctx *gin.Context) {
	var updateDto dto.UpdateStaffDTO
	errDTO := ctx.ShouldBind(&updateDto)
	if errDTO != nil {
		response := helper.ResponseErrorData(500, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	res, err := c.staffService.UpdateStaff(updateDto)
	if err != nil {
		if criteria.IsErrNotFound(err) {
			response := helper.ResponseErrorData(500, "Cannot find")
			ctx.JSON(http.StatusOK, response)
			return
		}
		if criteria.IsDuplicate(err) {
			response := helper.ResponseErrorData(528, "Email already exist!")
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

func (c *staffController) DeleteStaff(ctx *gin.Context) {
	var deleteDto dto.DeleteByIdDTO
	errDto := ctx.ShouldBind(&deleteDto)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	err := c.staffService.DeleteStaff(deleteDto.ID)
	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.ResponseData(0, "success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
