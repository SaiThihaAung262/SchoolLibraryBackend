package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/model"
	"MyGO.com/m/repository/criteria"
	"MyGO.com/m/service"
	"github.com/gin-gonic/gin"
)

type ClientAuthController interface {
	ClientLogin(ctx *gin.Context)
	GetClientByUUID(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
}

type clientAuthController struct {
	studentService service.StudentService
	teacherService service.TeacherService
	staffService   service.StaffService
	jwtService     service.JwtService
}

func NewClientAuthController(
	studentService service.StudentService,
	teacherService service.TeacherService,
	staffService service.StaffService,
	jwtService service.JwtService) ClientAuthController {
	return &clientAuthController{
		studentService: studentService,
		teacherService: teacherService,
		staffService:   staffService,
		jwtService:     jwtService,
	}

}

type ClientLoginResponse struct {
	Name     string `json:"name"`
	UserType uint64 `json:"user_type"`
	Token    string `json:"token"`
	UUID     string `json:"uuid"`
}

func (c *clientAuthController) ClientLogin(ctx *gin.Context) {
	var loginDTO dto.ClientLoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.ResponseErrorData(503, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	if loginDTO.Type == model.TeacherLoginType {
		loginResult := c.teacherService.VerifyLogin(loginDTO.Email, loginDTO.Password)

		if v, ok := loginResult.(model.Teacher); ok {
			generateToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
			// v.Token =
			responseData := ClientLoginResponse{
				Name:     v.Name,
				UserType: loginDTO.Type,
				Token:    generateToken,
				UUID:     v.UUID,
			}

			response := helper.ResponseData(0, "Login successfull", responseData)
			ctx.JSON(http.StatusOK, response)
			return
		}

	} else if loginDTO.Type == model.StudentLoginType {

		loginResult := c.studentService.VerifyLogin(loginDTO.Email, loginDTO.Password)

		if v, ok := loginResult.(model.Student); ok {
			generateToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
			// v.Token =
			responseData := ClientLoginResponse{
				Name:     v.Name,
				UserType: loginDTO.Type,
				Token:    generateToken,
				UUID:     v.UUID,
			}

			response := helper.ResponseData(0, "Login successfull", responseData)
			ctx.JSON(http.StatusOK, response)
			return
		}
	} else if loginDTO.Type == model.StaffLoginType {
		loginResult := c.staffService.VerifyLogin(loginDTO.Email, loginDTO.Password)

		if v, ok := loginResult.(model.Staff); ok {
			generateToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
			// v.Token =
			responseData := ClientLoginResponse{
				Name:     v.Name,
				UserType: loginDTO.Type,
				Token:    generateToken,
				UUID:     v.UUID,
			}

			response := helper.ResponseData(0, "Login successfull", responseData)
			ctx.JSON(http.StatusOK, response)
			return
		}
	}

	response := helper.ResponseErrorData(504, "Invalid email or password")
	ctx.JSON(http.StatusOK, response)
}

func (c clientAuthController) GetClientByUUID(ctx *gin.Context) {
	var getDto dto.GetUserByUUIDDto
	errDto := ctx.ShouldBind(&getDto)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	var responseData helper.ResponseUserData
	responseData.Type = getDto.Type

	if getDto.Type == model.TeacherLoginType {
		teacher, errGetTeacher := c.teacherService.GetTeacherByUUID(getDto.UUID)
		if errGetTeacher != nil {
			if criteria.IsErrNotFound(errGetTeacher) {
				response := helper.ResponseErrorData(500, "Cannot find teacher")
				ctx.JSON(http.StatusOK, response)
				return
			}
			response := helper.ResponseErrorData(500, errGetTeacher.Error())
			ctx.JSON(http.StatusOK, response)
			return
		}
		responseData.UserData = teacher

	} else if getDto.Type == model.StudentLoginType {
		student, errGetStudent := c.studentService.GetStudentByUUID(getDto.UUID)
		if errGetStudent != nil {
			if criteria.IsErrNotFound(errGetStudent) {
				response := helper.ResponseErrorData(500, "Cannot find student")
				ctx.JSON(http.StatusOK, response)
				return
			}
			response := helper.ResponseErrorData(500, errGetStudent.Error())
			ctx.JSON(http.StatusOK, response)
			return
		}
		responseData.UserData = student

	} else if getDto.Type == model.StaffLoginType {
		staff, errGetStudent := c.staffService.GetStaffByUUID(getDto.UUID)
		if errGetStudent != nil {
			if criteria.IsErrNotFound(errGetStudent) {
				response := helper.ResponseErrorData(500, "Cannot find student")
				ctx.JSON(http.StatusOK, response)
				return
			}
			response := helper.ResponseErrorData(500, errGetStudent.Error())
			ctx.JSON(http.StatusOK, response)
			return
		}
		responseData.UserData = staff
	}

	respnse := helper.ResponseData(0, "success", responseData)
	ctx.JSON(http.StatusOK, respnse)

}

func (c clientAuthController) ChangePassword(ctx *gin.Context) {
	var updatePassDto dto.ChangePasswordDto

	errDTO := ctx.ShouldBind(&updatePassDto)
	if errDTO != nil {
		response := helper.ResponseErrorData(503, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	fmt.Println("Here is the username and password", updatePassDto.ID, updatePassDto.Password, updatePassDto.Emial)
	if updatePassDto.Type == model.TeacherLoginType {
		loginResult := c.teacherService.VerifyLogin(updatePassDto.Emial, updatePassDto.Password)
		if _, ok := loginResult.(model.Teacher); ok {

			err := c.teacherService.ChangePassword(updatePassDto.ID, updatePassDto.NewPassword)
			if err != nil {
				response := helper.ResponseErrorData(504, err.Error())
				ctx.JSON(http.StatusOK, response)
				return
			}

			response := helper.ResponseData(0, "Change Password successful", helper.EmptyObj{})
			ctx.JSON(http.StatusOK, response)
			return
		}

	} else if updatePassDto.Type == model.StudentLoginType {
		loginResult := c.studentService.VerifyLogin(updatePassDto.Emial, updatePassDto.Password)
		if _, ok := loginResult.(model.Student); ok {

			err := c.studentService.ChangePassword(updatePassDto.ID, updatePassDto.NewPassword)
			if err != nil {
				response := helper.ResponseErrorData(504, err.Error())
				ctx.JSON(http.StatusOK, response)
				return
			}

			response := helper.ResponseData(0, "Change Password successful", helper.EmptyObj{})
			ctx.JSON(http.StatusOK, response)
			return
		}
	} else if updatePassDto.Type == model.StaffLoginType {
		loginResult := c.staffService.VerifyLogin(updatePassDto.Emial, updatePassDto.Password)
		if _, ok := loginResult.(model.Staff); ok {

			err := c.staffService.ChangePassword(updatePassDto.ID, updatePassDto.NewPassword)
			if err != nil {
				response := helper.ResponseErrorData(504, err.Error())
				ctx.JSON(http.StatusOK, response)
				return
			}

			response := helper.ResponseData(0, "Change Password successful", helper.EmptyObj{})
			ctx.JSON(http.StatusOK, response)
			return
		}
	}

	response := helper.ResponseErrorData(504, "Old password is wrong!")
	ctx.JSON(http.StatusOK, response)

}
