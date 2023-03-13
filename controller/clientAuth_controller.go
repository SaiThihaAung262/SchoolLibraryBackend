package controller

import (
	"net/http"
	"strconv"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/model"
	"MyGO.com/m/service"
	"github.com/gin-gonic/gin"
)

type ClientAuthController interface {
	ClientLogin(ctx *gin.Context)
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
	Name  string `json:"name"`
	Token string `json:"token"`
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

	} else {
		loginResult := c.studentService.VerifyLogin(loginDTO.Name, loginDTO.Password)

		if v, ok := loginResult.(model.User); ok {
			generateToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
			// v.Token =
			responseData := LoginResponse{
				Name:  v.Name,
				Token: generateToken,
			}

			response := helper.ResponseData(0, "Login successfull", responseData)
			ctx.JSON(http.StatusOK, response)
			return
		}
	}

	response := helper.ResponseErrorData(504, "Invalid uesr name or password")
	ctx.JSON(http.StatusOK, response)
}
