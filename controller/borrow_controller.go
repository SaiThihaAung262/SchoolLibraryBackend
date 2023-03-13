package controller

import (
	"fmt"
	"net/http"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/model"
	"MyGO.com/m/repository/criteria"
	"MyGO.com/m/service"
	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type BorrowController interface {
	CreateBorrow(ctx *gin.Context)
	GetBorrowHistory(ctx *gin.Context)
	UpdateBorrowStatus(ctx *gin.Context)
}

type borrowController struct {
	borrowService    service.Borrowservice
	bookService      service.BookService
	teacherService   service.TeacherService
	studentService   service.StudentService
	borrowLogService service.BorrowLogService
}

func NewBorrowController(borrowService service.Borrowservice,
	bookService service.BookService,
	teacherService service.TeacherService,
	studentService service.StudentService,
	borrowLogService service.BorrowLogService,
) BorrowController {
	return &borrowController{
		borrowService:    borrowService,
		bookService:      bookService,
		teacherService:   teacherService,
		studentService:   studentService,
		borrowLogService: borrowLogService,
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

	//---------------Check is book exist---------------
	book, errGetBook := c.bookService.GetBookByUUID(createDto.BookUUID)
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

	if book.BorrowQty >= book.AvailableQty {
		response := helper.ResponseErrorData(500, "These books are under borrow")
		ctx.JSON(http.StatusOK, response)
		return
	}

	var userTeacher model.Teacher
	var userStudent model.Student
	//---------------Check type for Teacher or student---------------
	if createDto.Type == 1 {
		teacher, errGetTeacher := c.teacherService.GetTeacherByUUID(createDto.UserUUID)
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
		userTeacher = *teacher
	} else {
		student, errGetStudent := c.studentService.GetStudentByUUID(createDto.UserUUID)
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
		userStudent = *student
	}

	//---------------Create borrow---------------
	err := c.borrowService.CreateBorrow(createDto)
	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	//---------------Create borrow logs---------------
	var borrowLog model.BorrowLog
	borrowLog.Type = createDto.Type
	borrowLog.BookID = book.ID
	borrowLog.BookUUID = book.UUID
	borrowLog.BookTitle = book.Title

	if createDto.Type == 1 {
		borrowLog.UserID = userTeacher.ID
		borrowLog.UserUUID = userTeacher.UUID
		borrowLog.UserName = userTeacher.Name
		borrowLog.Department = userTeacher.Department

	} else {
		borrowLog.UserID = userStudent.ID
		borrowLog.UserUUID = userStudent.UUID
		borrowLog.UserName = userStudent.Name
		borrowLog.RoleNo = userStudent.RoleNo
		borrowLog.Year = userStudent.Year
	}
	errCreateLog := c.borrowLogService.CreateBorrowLog(borrowLog)
	if errCreateLog != nil {
		response := helper.ResponseErrorData(500, errCreateLog.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	//---------------Update qty of borrow---------------
	var bookToUpdate dto.UpdateBookDTO
	bookToUpdate.ID = book.ID
	bookToUpdate.AvailableQty = book.AvailableQty
	bookToUpdate.BorrowQty = book.BorrowQty + 1
	_, errUpdateBook := c.bookService.UpdateBook(bookToUpdate)
	if err != nil {
		response := helper.ResponseErrorData(500, errUpdateBook.Error())
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
		fmt.Println("here have error <>>>>>>>>>>>>>>>")
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
		//---------------Get Book Data to response---------------
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

		//---------------Get User Data to response---------------
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

			err := smapping.FillStruct(&borrowUser, smapping.MapFields(&teacher))
			if err != nil {
				response := helper.ResponseErrorData(500, err.Error())
				ctx.JSON(http.StatusOK, response)
				return
			}
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

			err := smapping.FillStruct(&borrowUser, smapping.MapFields(&student))
			if err != nil {
				response := helper.ResponseErrorData(500, err.Error())
				ctx.JSON(http.StatusOK, response)
				return
			}
		}

		responseData.ID = item.ID
		responseData.Type = item.Type
		responseData.Status = item.Status
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

func (c borrowController) UpdateBorrowStatus(ctx *gin.Context) {
	var updateDto dto.UpdateBorrowStatusDTO
	errDTO := ctx.ShouldBind(&updateDto)
	if errDTO != nil {
		response := helper.ResponseErrorData(500, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	res, err := c.borrowService.UpdateBorrowStatus(updateDto)
	if err != nil {
		if criteria.IsErrNotFound(err) {
			response := helper.ResponseErrorData(500, "Cannot find")
			ctx.JSON(http.StatusOK, response)
			return
		}
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	fmt.Println("Here is res . book >>>>>>>>>>>>>>>>>>>>>>>", res.BookUUID)

	// Check is book exist
	book, errGetBook := c.bookService.GetBookByUUID(res.BookUUID)
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

	//Update qty of borrow
	var bookToUpdate dto.UpdateBookDTO
	bookToUpdate.ID = book.ID
	bookToUpdate.AvailableQty = book.AvailableQty
	bookToUpdate.BorrowQty = book.BorrowQty - 1
	_, errUpdateBook := c.bookService.UpdateBook(bookToUpdate)
	if err != nil {
		response := helper.ResponseErrorData(500, errUpdateBook.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.ResponseData(0, "success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)

}
