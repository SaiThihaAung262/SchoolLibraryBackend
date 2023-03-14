package controller

import (
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
}

type clientAuthController struct {
	studentService service.StudentService
	teacherService service.TeacherService
	jwtService     service.JwtService
}

func NewClientAuthController(studentService service.StudentService,
	teacherService service.TeacherService,
	jwtService service.JwtService) ClientAuthController {
	return &clientAuthController{
		studentService: studentService,
		teacherService: teacherService,
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

	if loginDTO.Type == 1 {
		loginResult := c.teacherService.VerifyLogin(loginDTO.Name, loginDTO.Password)

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

	} else {

		loginResult := c.studentService.VerifyLogin(loginDTO.Name, loginDTO.Password)

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
	}

	response := helper.ResponseErrorData(504, "Invalid username or password")
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

	if getDto.Type == 1 {
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

	} else {
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

	}

	respnse := helper.ResponseData(0, "success", responseData)
	ctx.JSON(http.StatusOK, respnse)

}
