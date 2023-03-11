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

type StudentController interface {
	CreateStudent(ctx *gin.Context)
	GetAllStudents(ctx *gin.Context)
	UpdateStudent(ctx *gin.Context)
	DeleteStudent(ctx *gin.Context)
}

type studentController struct {
	studentService service.StudentService
	jwtService     service.JwtService
}

func NewStudentController(studentService service.StudentService, jwtService service.JwtService) StudentController {
	return &studentController{
		studentService: studentService,
		jwtService:     jwtService,
	}
}

func (c *studentController) CreateStudent(ctx *gin.Context) {
	var registerDTO dto.StudentRegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.ResponseErrorData(501, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	createUser, err := c.studentService.InsertStudent(registerDTO)
	if err != nil {
		if criteria.IsErrNotFound(err) {
			response := helper.ResponseErrorData(500, "Cannot find")
			ctx.JSON(http.StatusOK, response)
			return
		}
		if criteria.IsDuplicate(err) {
			response := helper.ResponseErrorData(528, "Email or Role number is already exist !")
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

type ResponseStudentData struct {
	List  []model.Student `json:"list"`
	Total int64           `json:"total"`
}

func (c *studentController) GetAllStudents(ctx *gin.Context) {
	req := &dto.StudentGetRequest{}

	if err := ctx.ShouldBind(&req); err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	result, count, err := c.studentService.GetAllStudents(req)

	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	if count == 0 {
		var students []model.Student
		var responseList ResponseStudentData
		responseList.List = students
		responseList.Total = count

		response := helper.ResponseData(0, "success", responseList)
		ctx.JSON(http.StatusOK, response)
		return
	}

	var responseList ResponseStudentData
	responseList.List = result
	responseList.Total = count

	response := helper.ResponseData(0, "success", responseList)
	ctx.JSON(http.StatusOK, response)

}

func (c *studentController) UpdateStudent(ctx *gin.Context) {
	var updateDto dto.UpdateStudentDTO
	errDTO := ctx.ShouldBind(&updateDto)
	if errDTO != nil {
		response := helper.ResponseErrorData(500, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	res, err := c.studentService.UpdateStudent(updateDto)
	if err != nil {
		if criteria.IsErrNotFound(err) {
			response := helper.ResponseErrorData(500, "Cannot find")
			ctx.JSON(http.StatusOK, response)
			return
		}
		if criteria.IsDuplicate(err) {
			response := helper.ResponseErrorData(528, "Email or role number already exist!")
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

func (c *studentController) DeleteStudent(ctx *gin.Context) {
	var deleteDto dto.DeleteByIdDTO
	errDto := ctx.ShouldBind(&deleteDto)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	err := c.studentService.DeleteStudent(deleteDto.ID)
	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.ResponseData(0, "success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
