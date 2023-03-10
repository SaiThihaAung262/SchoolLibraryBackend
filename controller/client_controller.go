package controller

import (
	"net/http"

	"MyGO.com/m/dto"
	"MyGO.com/m/helper"
	"MyGO.com/m/repository/criteria"
	"MyGO.com/m/service"
	"github.com/gin-gonic/gin"
)

type ClientController interface {
	CreateClient(ctx *gin.Context)
	// GetAllUsers(ctx *gin.Context)
	// UpdateUser(ctx *gin.Context)
	// DeleteUser(ctx *gin.Context)
}

type clientController struct {
	clientService service.ClientService
	jwtService    service.JwtService
}

func NewClientController(clientService service.ClientService, jwtService service.JwtService) ClientController {
	return &clientController{
		clientService: clientService,
		jwtService:    jwtService,
	}
}

func (c *clientController) CreateClient(ctx *gin.Context) {
	var registerDTO dto.ClientRegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.ResponseErrorData(501, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	createUser, err := c.clientService.InsertClient(registerDTO)
	if err != nil {
		if criteria.IsErrNotFound(err) {
			response := helper.ResponseErrorData(500, "Cannot find")
			ctx.JSON(http.StatusOK, response)
			return
		}
		if criteria.IsDuplicate(err) {
			response := helper.ResponseErrorData(528, "Email is Duplicate !")
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
