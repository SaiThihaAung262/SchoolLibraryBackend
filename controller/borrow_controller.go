package controller

import (
	"fmt"
	"net/http"
	"time"

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
	GetBookSummaryData(ctx *gin.Context)
	GetBookByUUID(ctx *gin.Context)
	ReBorrow(ctx *gin.Context)
}

type borrowController struct {
	borrowService       service.Borrowservice
	bookService         service.BookService
	teacherService      service.TeacherService
	studentService      service.StudentService
	staffService        service.StaffService
	borrowLogService    service.BorrowLogService
	systemConfigService service.SystemConfigService
	punishmentService   service.PunishmentService
}

func NewBorrowController(borrowService service.Borrowservice,
	bookService service.BookService,
	teacherService service.TeacherService,
	studentService service.StudentService,
	staffService service.StaffService,
	borrowLogService service.BorrowLogService,
	systemConfigService service.SystemConfigService,
	punishmentService service.PunishmentService,

) BorrowController {
	return &borrowController{
		borrowService:       borrowService,
		bookService:         bookService,
		teacherService:      teacherService,
		studentService:      studentService,
		staffService:        staffService,
		borrowLogService:    borrowLogService,
		systemConfigService: systemConfigService,
		punishmentService:   punishmentService,
	}
}

// --------------Create Borrow-------------
func (c borrowController) CreateBorrow(ctx *gin.Context) {
	var createDto dto.CreateBorrowDTO
	errDto := ctx.ShouldBind(&createDto)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	//*Check user alreay borrow this book or not
	// isAlreadyBoorowThisBook := c.borrowService.IsAlreadyBorrowThisBook(createDto.UserUUID, createDto.BookUUID)
	// if isAlreadyBoorowThisBook {
	// 	response := helper.ResponseErrorData(888, "Already borrow this book!")
	// 	ctx.JSON(http.StatusOK, response)
	// 	return
	// }
	//*Get expired count with uuid

	reqExpiredCountDto := &dto.BorrowHistoryRequest{
		UserUUID: createDto.UserUUID,
		Status:   model.BookBorrowExpireStatus,
	}

	_, expiredCount, errGetExpiredCount := c.borrowService.GetBorrowHistory(reqExpiredCountDto)
	if errGetExpiredCount != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	fmt.Println("here is expired Count!", expiredCount)

	//*---------------Check is book exist---------------
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

	if book.Status == model.BookDamageLostStatus {
		response := helper.ResponseErrorData(500, "These books have been Damage or Lost !")
		ctx.JSON(http.StatusOK, response)
		return
	}

	//*Get User Borrowing book count
	reqBorrowCountDto := &dto.BorrowHistoryRequest{
		UserUUID: createDto.UserUUID,
	}

	_, borrowingCount, errGetBorrowingCount := c.borrowService.GetBorrowingAndExpireData(reqBorrowCountDto)
	if errGetBorrowingCount != nil {
		response := helper.ResponseErrorData(500, errGetBorrowingCount.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	//*Get Config data
	configData, errGetConfig := c.systemConfigService.GetSystemConfig()

	if errGetConfig != nil {
		response := helper.ResponseErrorData(500, errGetConfig.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	//*---------------Check type for Teacher or student---------------
	var userTeacher model.Teacher
	var userStudent model.Student
	var userStaff model.Staff
	if createDto.Type == model.TeacherBorrow {

		//*Check can borrowcount is greather than or equal borrowing count
		if uint64(borrowingCount) >= configData.TeacherCanBorrowCount && expiredCount > 0 {
			respMsg := fmt.Sprintf("Borrowing Limit is full and there has %d expired borrow book", expiredCount)
			response := helper.ResponseErrorData(500, respMsg)
			ctx.JSON(http.StatusOK, response)
			return
		}
		if uint64(borrowingCount) >= configData.TeacherCanBorrowCount {
			response := helper.ResponseErrorData(500, "Borrowing Limit is Full!")
			ctx.JSON(http.StatusOK, response)
			return
		}

		//*Get teacher by UUID
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
	} else if createDto.Type == model.StudentBorrow {

		//*Check can borrowcount is greather than or equal borrowing count
		if uint64(borrowingCount) >= configData.StudentCanBorrowCount && expiredCount > 0 {
			respMsg := fmt.Sprintf("Borrowing Limit is full and there has %d expired borrow book", expiredCount)
			response := helper.ResponseErrorData(500, respMsg)
			ctx.JSON(http.StatusOK, response)
			return
		}
		if uint64(borrowingCount) >= configData.StudentCanBorrowCount {
			response := helper.ResponseErrorData(500, "Borrowing limit is Full!")
			ctx.JSON(http.StatusOK, response)
			return
		}

		//*Get student by UUID
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
	} else if createDto.Type == model.StaffBorrow {

		//*Check can borrowcount is greather than or equal borrowing count
		if uint64(borrowingCount) >= configData.StaffCanBorrowCount && expiredCount > 0 {
			respMsg := fmt.Sprintf("Borrowing Limit is full and there has %d expired borrow book", expiredCount)
			response := helper.ResponseErrorData(500, respMsg)
			ctx.JSON(http.StatusOK, response)
			return
		}
		if uint64(borrowingCount) >= configData.StaffCanBorrowCount {
			response := helper.ResponseErrorData(500, "Borrowing limit is Full!")
			ctx.JSON(http.StatusOK, response)
			return
		}

		//*Get staff by UUID
		staff, errGetStaff := c.staffService.GetStaffByUUID(createDto.UserUUID)
		if errGetStaff != nil {
			if criteria.IsErrNotFound(errGetStaff) {
				response := helper.ResponseErrorData(500, "Cannot find staff")
				ctx.JSON(http.StatusOK, response)
				return
			}
			response := helper.ResponseErrorData(500, errGetStaff.Error())
			ctx.JSON(http.StatusOK, response)
			return
		}
		userStaff = *staff
	}

	//*---------------Create borrow---------------
	borrowRes, err := c.borrowService.CreateBorrow(createDto)
	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	//*---------------Create borrow logs---------------
	var borrowLog model.BorrowLog
	borrowLog.BorrowID = borrowRes.ID
	borrowLog.Type = createDto.Type
	borrowLog.BookID = book.ID
	borrowLog.BookUUID = book.UUID
	borrowLog.BookTitle = book.Title

	if createDto.Type == model.TeacherBorrow {
		borrowLog.UserID = userTeacher.ID
		borrowLog.UserUUID = userTeacher.UUID
		borrowLog.UserName = userTeacher.Name
		borrowLog.Department = userTeacher.Department

	} else if createDto.Type == model.StudentBorrow {
		borrowLog.UserID = userStudent.ID
		borrowLog.UserUUID = userStudent.UUID
		borrowLog.UserName = userStudent.Name
		borrowLog.RoleNo = userStudent.RoleNo
		borrowLog.Year = userStudent.Year
	} else if createDto.Type == model.StaffBorrow {
		borrowLog.UserID = userStaff.ID
		borrowLog.UserUUID = userStaff.UUID
		borrowLog.UserName = userStaff.Name
		borrowLog.Department = userStaff.Department
	}

	errCreateLog := c.borrowLogService.CreateBorrowLog(borrowLog)
	if errCreateLog != nil {
		response := helper.ResponseErrorData(500, errCreateLog.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	//*---------------Update qty of borrow---------------
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

// *--------------Get Borrow History-------------
func (c borrowController) GetBorrowHistory(ctx *gin.Context) {
	reqDto := &dto.BorrowHistoryRequest{}
	errDto := ctx.ShouldBind(&reqDto)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	//*Get Config data
	configData, errGetConfig := c.systemConfigService.GetSystemConfig()

	if errGetConfig != nil {
		response := helper.ResponseErrorData(500, errGetConfig.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	resultData, total, err := c.borrowService.GetBorrowHistory(reqDto)
	if err != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	var responseList dto.BorrowHistoryList
	responseList.Total = total

	for _, item := range resultData {
		responseData := dto.BorrowHistoryResponse{}

		//*---------------Get Book Data to response---------------
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

		//*---------------Get User Data to response---------------
		var borrowUser dto.BorrowUser
		if item.Type == model.TeacherBorrow {
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

			responseData.ExpiredAt = helper.CalculatExpireDate(item.CreatedAt, int(configData.TeacherCanBorrowDay))

		} else if item.Type == model.StudentBorrow {
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

			responseData.ExpiredAt = helper.CalculatExpireDate(item.CreatedAt, int(configData.StudentCanBorrowDay))
		} else if item.Type == model.StaffBorrow {
			staff, errGetStaff := c.staffService.GetStaffByUUID(item.UserUUID)
			if errGetStaff != nil {
				if criteria.IsErrNotFound(errGetStaff) {
					response := helper.ResponseErrorData(500, "Cannot find teacher")
					ctx.JSON(http.StatusOK, response)
					return
				}
				response := helper.ResponseErrorData(500, errGetStaff.Error())
				ctx.JSON(http.StatusOK, response)
				return
			}

			err := smapping.FillStruct(&borrowUser, smapping.MapFields(&staff))
			if err != nil {
				response := helper.ResponseErrorData(500, err.Error())
				ctx.JSON(http.StatusOK, response)
				return
			}

			responseData.ExpiredAt = helper.CalculatExpireDate(item.CreatedAt, int(configData.StaffCanBorrowDay))

		}

		responseData.ID = item.ID
		responseData.Type = item.Type
		responseData.Status = item.Status
		responseData.User = borrowUser
		responseData.Book = book
		responseData.CreatedAt = item.CreatedAt
		responseData.UpdatedAt = item.UpdatedAt
		expiredDay, punishmentAmt := CalcExpireDayAndPunishAmt(c, ctx, responseData.ExpiredAt, responseData)
		responseData.ExpiredDay = expiredDay
		responseData.PunishAmount = punishmentAmt
		responseList.List = append(responseList.List, responseData)

	}

	response := helper.ResponseData(0, "success", responseList)
	ctx.JSON(http.StatusOK, response)
}

// *Caluclate Expire day and Punishment amount
func CalcExpireDayAndPunishAmt(c borrowController, ctx *gin.Context, expireTime time.Time, data dto.BorrowHistoryResponse) (uint64, uint64) {
	var expiredDay uint64
	var punishmentAmt uint64

	//*---------Caluclate expire time from now---------
	calcExpireTime := time.Since(expireTime)

	expiredDay = uint64(calcExpireTime.Hours() / 24) // asssing value to Expired Day

	//*---------Get Punishment Data and calculate punishment amount---------
	punishmentLists, err := c.punishmentService.GetPunishmentData()

	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		ctx.Abort()
	}

	for _, item := range punishmentLists {
		// if expiredDay <= item.DurationEnd {
		// 	if data.Type == model.TeacherBorrow {
		// 		punishmentAmt = item.TeacherPunishAmount
		// 	} else if data.Type == model.StudentBorrow {
		// 		punishmentAmt = item.StudentPunishAmount
		// 	} else if data.Type == model.StaffBorrow {
		// 		punishmentAmt = item.StaffPunishAmount
		// 	}
		// }

		if expiredDay >= item.DurationStart && expiredDay < item.DurationEnd && expiredDay < punishmentLists[0].DurationEnd {
			if data.Type == model.TeacherBorrow {
				punishmentAmt = item.TeacherPunishAmount
			} else if data.Type == model.StudentBorrow {
				punishmentAmt = item.StudentPunishAmount
			} else if data.Type == model.StaffBorrow {
				punishmentAmt = item.StaffPunishAmount
			}
		}

		if expiredDay >= item.DurationEnd {
			var expiredDayOrWeekOryear uint64
			myRemainDay := expiredDay % item.DurationEnd
			if myRemainDay > 0 {
				expiredDayOrWeekOryear = (expiredDay / item.DurationEnd) + 1
			} else {
				expiredDayOrWeekOryear = expiredDay / item.DurationEnd
			}

			if data.Type == model.TeacherBorrow {
				punishmentAmt = expiredDayOrWeekOryear * item.TeacherPunishAmount
			} else if data.Type == model.StudentBorrow {
				punishmentAmt = expiredDayOrWeekOryear * item.StudentPunishAmount
			} else if data.Type == model.StaffBorrow {
				punishmentAmt = expiredDayOrWeekOryear * item.StaffPunishAmount
			}

		}

	}

	//*---------check borrow status and expired days to Update borrow status to expired---------
	if data.Status == 1 && expiredDay > 0 {
		updatStatusDto := dto.UpdateBorrowStatusDTO{
			ID:       data.ID,
			UserUUID: data.User.UUID,
			BookUUID: data.Book.UUID,
			Type:     data.Type,
			Status:   model.BookBorrowExpireStatus,
		}
		_, updateErr := c.borrowService.UpdateBorrowStatus(updatStatusDto)
		if err != nil {
			if criteria.IsErrNotFound(err) {
				response := helper.ResponseErrorData(500, "Cannot find")
				ctx.JSON(http.StatusOK, response)
				ctx.Abort()
			}
			response := helper.ResponseErrorData(500, updateErr.Error())
			ctx.JSON(http.StatusOK, response)
			ctx.Abort()
		}
	}

	return expiredDay, punishmentAmt
}

// * --------------Update Borrow Status-------------
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

	// *Check is book exist
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

	//*Update qty of borrow
	var bookToUpdate dto.UpdateBookDTO
	bookToUpdate.ID = book.ID
	bookToUpdate.AvailableQty = book.AvailableQty
	bookToUpdate.BorrowQty = book.BorrowQty - 1
	// _, errUpdateBook := c.bookService.UpdateBook(bookToUpdate)
	errUpdateBook := c.bookService.UpdateBookBorrowQTY(bookToUpdate.ID, bookToUpdate.AvailableQty, bookToUpdate.BorrowQty)

	if errUpdateBook != nil {
		response := helper.ResponseErrorData(500, errUpdateBook.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.ResponseData(0, "success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)

}

// *--------------Get Book Summary-------------
func (c borrowController) GetBookSummaryData(ctx *gin.Context) {
	reqSummaryDto := &dto.ReqBookSummary{}
	errReqSummaryDto := ctx.ShouldBind(&reqSummaryDto)
	if errReqSummaryDto != nil {
		response := helper.ResponseErrorData(500, errReqSummaryDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	reqBookDto := &dto.BookGetRequest{}

	//*----------Bind Book req dto with req Summary dto----------
	err := smapping.FillStruct(&reqBookDto, smapping.MapFields(&reqSummaryDto))
	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	//*----------Request Summary Data dto----------
	reqSummaryDataDto := &dto.ReqBorrowCountByBookUUIDAndDateDto{}
	errReqSummaryDataDto := smapping.FillStruct(&reqSummaryDataDto, smapping.MapFields(&reqSummaryDto))
	if errReqSummaryDataDto != nil {
		response := helper.ResponseErrorData(500, errReqSummaryDataDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	//*Get all books
	resultData, count, err := c.bookService.GetAllBooks(reqBookDto)

	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	var bookSummaryData helper.ResponseSummaryData

	var bookSummaryDataList helper.ResponseSummaryDataList
	bookSummaryDataList.Total = uint64(count)

	for _, item := range resultData {
		reqSummaryDataDto.BookUUID = item.UUID
		bookSummaryData.BookDetail = item

		total, errReq := c.borrowLogService.GetBorrowCountByBookUUIDAndDate(reqSummaryDataDto)
		if errReq != nil {
			response := helper.ResponseErrorData(500, errReq.Error())
			ctx.JSON(http.StatusOK, response)
			return
		}
		bookSummaryData.BorrowCount = total

		book, errReq := c.bookService.GetBookByUUIDAndDate(reqSummaryDataDto)
		if errReq != nil {
			bookSummaryData.TotalBook = 0
		} else {
			bookSummaryData.TotalBook = book.AvailableQty
		}
		bookSummaryDataList.List = append(bookSummaryDataList.List, bookSummaryData)
	}
	response := helper.ResponseData(0, "success", bookSummaryDataList)
	ctx.JSON(http.StatusOK, response)
}

// *--------------Get Book by UUID-------------
func (c borrowController) GetBookByUUID(ctx *gin.Context) {
	var dtoBook dto.GetBookByUUIDDto
	errDto := ctx.ShouldBind(&dtoBook)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	reqSummaryDataDto := dto.ReqBorrowCountByBookUUIDAndDateDto{}
	reqSummaryDataDto.BookUUID = dtoBook.UUID

	borrowCount, errReq := c.borrowLogService.GetBorrowCountByBookUUIDAndDate(&reqSummaryDataDto)
	if errReq != nil {
		response := helper.ResponseErrorData(500, errReq.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	book, errGetBook := c.bookService.GetBookByUUID(dtoBook.UUID)
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

	responseData := helper.ResponBookDetailByUUID{
		BookDetail:  *book,
		BorrowCount: borrowCount,
	}

	response := helper.ResponseData(0, "success", responseData)
	ctx.JSON(http.StatusOK, response)
}

func (c borrowController) ReBorrow(ctx *gin.Context) {
	c.CreateBorrow(ctx)
}
