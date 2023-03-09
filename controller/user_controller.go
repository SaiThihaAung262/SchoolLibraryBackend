package controller

import (
	"fmt"
	"net/http"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/model"
	"MyGO.com/m/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUsers(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	
}

type userController struct {
	userService service.UserService
	jwtService  service.JwtService
}

func NewUserController(userService service.UserService, jwtService service.JwtService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

type ResponseUserListData struct {
	List  []model.User `json:"list"`
	Total int64        `json:"total"`
}

func (c *userController) GetAllUsers(ctx *gin.Context) {

	req := &dto.UserGetRequest{}

	if err := ctx.ShouldBind(&req); err != nil {
		response := helper.ResponseErrorData(500, "Internal server error !")
		ctx.JSON(http.StatusOK, response)
		return
	}

	result, count, err := c.userService.GetAllUsers(req)

	if count == 0 {
		response := helper.ResponseErrorData(512, "Record not found")
		ctx.JSON(http.StatusOK, response)
		return
	}

	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	var responseList ResponseUserListData

	responseList.List = result
	responseList.Total = count

	response := helper.ResponseData(0, "success", responseList)
	ctx.JSON(http.StatusOK, response)

}
func (c *userController) UpdateUser(ctx *gin.Context) {

	var updateUserDto dto.UpdateUserDto
	errDTO := ctx.ShouldBind(&updateUserDto)
	if errDTO != nil {
		fmt.Println("Chee pare ma bind twar bu")
		response := helper.ResponseErrorData(503, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	isExit := c.userService.IsUserExist(updateUserDto.ID)
	if !isExit {
		response := helper.ResponseErrorData(502, "Record not found !")
		ctx.JSON(http.StatusOK, response)
		return
	}

	isDuplicate := c.userService.IsDuplicateEmail(updateUserDto.Email)
	if isDuplicate {
		response := helper.ResponseErrorData(502, "Email Already Exit !")
		ctx.JSON(http.StatusOK, response)
		return
	}

	updateUser := c.userService.UpdateUser(updateUserDto)
	response := helper.ResponseData(0, "Success", updateUser)
	ctx.JSON(http.StatusOK, response)

}

func (c *userController) DeleteUser(ctx *gin.Context) {

	var deleteDTO dto.DeleteByIdDTO

	errDTO := ctx.ShouldBind(&deleteDTO)

	if errDTO != nil {
		fmt.Println("Chee pare ma bind twar bu")
		response := helper.ResponseErrorData(503, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	isExit := c.userService.IsUserExist(deleteDTO.ID)
	if !isExit {
		response := helper.ResponseErrorData(502, "Record not found !")
		ctx.JSON(http.StatusOK, response)
		return
	}

	err := c.userService.DeleteUser(deleteDTO.ID)
	if err != nil {
		response := helper.ResponseErrorData(503, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.ResponseData(0, "success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
