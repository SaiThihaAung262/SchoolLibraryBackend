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
	GetBorrowHistory(ctx *gin.Context)
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
		borrowService:  borrowService,
		bookService:    bookService,
		teacherService: teacherService,
		studentService: studentService,
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

	if createDto.Type == 1 {
		_, errGetTeacher := c.teacherService.GetTeacherByUUID(createDto.UserUUID)
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
	} else {
		_, errGetStudent := c.studentService.GetStudentByUUID(createDto.UserUUID)
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

func (c borrowController) GetBorrowHistory(ctx *gin.Context) {
	reqDto := &dto.BorrowHistoryRequest{}
	errDto := ctx.ShouldBind(&reqDto)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	resultData, total, err := c.borrowService.GetBorrowHistory(reqDto)
	if err != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	fmt.Println("----- here is total count ----", total)

	var responseList dto.BorrowHistoryList
	responseList.Total = total

	for _, item := range resultData {
		responseData := dto.BorrowHistoryResponse{}

		book, errGetBook := c.bookService.GetBookByUUID(item.BookUUID)
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

		var borrowUser dto.BorrowUser
		if item.Type == 1 {
			teacher, errGetTeacher := c.teacherService.GetTeacherByUUID(item.UserUUID)
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
			borrowUser.ID = teacher.ID
			borrowUser.UUID = teacher.UUID
			borrowUser.Name = teacher.Name
			borrowUser.Email = teacher.Email
			borrowUser.Department = teacher.Department
			borrowUser.Year = 0
			borrowUser.RoleNo = ""
		} else {
			student, errGetStudent := c.studentService.GetStudentByUUID(item.UserUUID)
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
			borrowUser.ID = student.ID
			borrowUser.UUID = student.UUID
			borrowUser.Name = student.Name
			borrowUser.Email = student.Email
			borrowUser.Year = student.Year
			borrowUser.RoleNo = student.RoleNo
			borrowUser.Department = 0
		}

		responseData.ID = item.ID
		responseData.Type = item.Type
		responseData.User = borrowUser
		responseData.Book = book
		responseData.CreatedAt = item.CreatedAt
		responseData.UpdatedAt = item.UpdatedAt
		responseData.ExpiredAt = helper.AddSevenDay(item.CreatedAt)

		responseList.List = append(responseList.List, responseData)

	}

	response := helper.ResponseData(0, "success", responseList)
	ctx.JSON(http.StatusOK, response)

}
