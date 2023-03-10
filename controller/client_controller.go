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

type ClientController interface {
	CreateClient(ctx *gin.Context)
	GetAllClients(ctx *gin.Context)
	UpdateClient(ctx *gin.Context)
	DeleteClient(ctx *gin.Context)
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

type ResponseClientData struct {
	List  []model.Client `json:"list"`
	Total int64          `json:"total"`
}

func (c *clientController) GetAllClients(ctx *gin.Context) {
	req := &dto.ClientGetRequest{}

	if err := ctx.ShouldBind(&req); err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	result, count, err := c.clientService.GetAllClients(req)

	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	if count == 0 {
		response := helper.ResponseErrorData(500, "Request not found")
		ctx.JSON(http.StatusOK, response)
		return
	}

	var responseList ResponseClientData
	responseList.List = result
	responseList.Total = count

	response := helper.ResponseData(0, "success", responseList)
	ctx.JSON(http.StatusOK, response)

}

func (c *clientController) UpdateClient(ctx *gin.Context) {
	var updateDto dto.UpdateClientDTO
	errDTO := ctx.ShouldBind(&updateDto)
	if errDTO != nil {
		response := helper.ResponseErrorData(500, errDTO.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	res, err := c.clientService.UpdateClient(updateDto)
	if err != nil {
		if criteria.IsErrNotFound(err) {
			response := helper.ResponseErrorData(500, "Cannot find")
			ctx.JSON(http.StatusOK, response)
			return
		}
		if criteria.IsDuplicate(err) {
			response := helper.ResponseErrorData(528, "This email already exist!")
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

func (c *clientController) DeleteClient(ctx *gin.Context) {
	var deleteDto dto.DeleteByIdDTO
	errDto := ctx.ShouldBind(&deleteDto)
	if errDto != nil {
		response := helper.ResponseErrorData(500, errDto.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	err := c.clientService.DeleteClient(deleteDto.ID)
	if err != nil {
		response := helper.ResponseErrorData(500, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.ResponseData(0, "success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
