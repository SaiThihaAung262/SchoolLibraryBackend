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

type BorrowController interface {
	CreateBorrow(ctx *gin.Context)
}

type borrowController struct {
	borrowService  service.Borrowservice
	bookService    service.BookService
	teacherService service.TeacherService
	studentService service.StudentService
}

func NewBorrowController(borrowService service.Borrowservice,
	bookService service.BookService,
	teacherService service.TeacherService,
	studentService service.StudentService,
) BorrowController {
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

	fmt.Println("HEre is book uuid >>>>>>>>", createDto.BookUUID)

	_, errGetBook := c.bookService.GetBookByUUID(createDto.BookUUID)
	if errGetBook != nil {
		if criteria.IsErrNotFound(errGetBook) {
			response := helper.ResponseErrorData(500, "Cannot find book")
			ctx.JSON(http.StatusOK, response)
			return
		}
		response := helper.ResponseErrorData(500, errGetBook.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	// if createDto.Type == 1 {
	// 	_, errGetTeacher := c.teacherService.GetTeacherByUUID(createDto.UserUUID)
	// 	if errGetTeacher != nil {
	// 		if criteria.IsErrNotFound(errGetBook) {
	// 			response := helper.ResponseErrorData(500, "Cannot find teacher")
	// 			ctx.JSON(http.StatusOK, response)
	// 			return
	// 		}
	// 		response := helper.ResponseErrorData(500, errGetBook.Error())
	// 		ctx.JSON(http.StatusOK, response)
	// 		return
	// 	}
	// } else {
	// 	_, errGetStudent := c.studentService.GetStudentByUUID(createDto.UserUUID)
	// 	if errGetStudent != nil {
	// 		if criteria.IsErrNotFound(errGetBook) {
	// 			response := helper.ResponseErrorData(500, "Cannot find student")
	// 			ctx.JSON(http.StatusOK, response)
	// 			return
	// 		}
	// 		response := helper.ResponseErrorData(500, errGetBook.Error())
	// 		ctx.JSON(http.StatusOK, response)
	// 		return
	// 	}
	// }

	err := c.borrowService.CreateBorrow(createDto)
	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.ResponseData(0, "success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)

}
