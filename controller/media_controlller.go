package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"MyGO.com/m/helper"
	"MyGO.com/m/model"
	"MyGO.com/m/service"
	"github.com/gin-gonic/gin"
)

type MediaController interface {
	CreateMedia(ctx *gin.Context)
}

type mediaController struct {
	mediaService service.MedeiaService
}

func NewMediaController(mediaService service.MedeiaService) MediaController {
	return &mediaController{
		mediaService: mediaService,
	}
}

func CreateFileName(ext string) string {
	return time.Now().Format("20060102150405") + ext
}

func (c mediaController) CreateMedia(ctx *gin.Context) {
	// pwd, _ := os.Getwd()
	// fmt.Println("here is my file path", pwd)
	// dst := pwd

	dst := "/Users/thihaaung/Documents/SchoolLibraryProject"
	dstClient := "/Users/thihaaung/Documents/SchoolLibraryProject"

	file, err := ctx.FormFile("file")
	if err != nil {
		response := helper.ResponseErrorData(500, "ENABLE_TO_GET_MEDIA_FILE")
		ctx.JSON(200, response)
		return
	}

	cType := file.Header["Content-Type"][0]
	types := strings.Split(cType, "/")
	if len(types) < 1 {
		response := helper.ResponseErrorData(500, "SERVER_CANNOT_PROCESS_REQUEST")
		ctx.JSON(200, response)
		return
	}

	mediaType := types[0]
	extension := filepath.Ext(file.Filename)
	if mediaType == "image" {
		dst += "/AdminSide/public/images"
		dstClient += "/ClinetSide/public/images"
	}
	filename := CreateFileName(extension)
	fileLocation := fmt.Sprintf("%v/%v", dst, filename)
	fileLocationClient := fmt.Sprintf("%v/%v", dstClient, filename)

	err = ctx.SaveUploadedFile(file, fileLocation)
	if err != nil {
		defer os.Remove(fileLocation)

		fmt.Println("Error Saving file>>>>>>>>>>", err)
		respone := helper.ResponseErrorData(500, "ERROR_SAVING_FILE")
		ctx.JSON(500, respone)
		return
	}
	errClientFile := ctx.SaveUploadedFile(file, fileLocationClient)
	if errClientFile != nil {
		defer os.Remove(fileLocationClient)

		fmt.Println("Error Saving file>>>>>>>>>>", errClientFile)
		respone := helper.ResponseErrorData(500, "ERROR_SAVING_FILE")
		ctx.JSON(500, respone)
		return
	}

	url := fmt.Sprintf("/images/%v", filename)

	media := model.Media{
		FileName:  filename,
		URL:       url,
		Extension: strings.Replace(extension, ".", "", 1),
		Type:      types[0],
	}

	createMedia, err := c.mediaService.CreateMedia(&media)
	if err != nil {
		defer os.Remove(fileLocation)
		response := helper.ResponseErrorData(500, "ERROR_READING_FILE")
		ctx.JSON(200, response)
		return
	}

	response := helper.ResponseData(0, "success", createMedia)

	ctx.JSON(http.StatusOK, response)
}
