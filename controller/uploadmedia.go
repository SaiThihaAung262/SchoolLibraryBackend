package controller

import (
	"github.com/gin-gonic/gin"
)

func (ctr *mediaController) uploadMedia(c *gin.Context) {
	// admin := c.MustGet("admin").(*model.Admin)
	// lang := c.MustGet("acc_lang").(string)

	// res := dto.ResponseObject{}
	// dst := "media/"
	// file, err := c.FormFile("file")
	// if err != nil {
	// 	res.ErrCode = 400
	// 	res.ErrMsg = utils.ChangeLang(lang, "ENABLE_TO_GET_MEDIA_FILE")
	// 	c.JSON(200, res)
	// 	return
	// }

	// cType := file.Header["Content-Type"][0]
	// types := strings.Split(cType, "/")
	// if len(types) < 1 {
	// 	res.ErrCode = 400
	// 	res.ErrMsg = utils.ChangeLang(lang, "SERVER_CANNOT_PROCESS_REQUEST")
	// 	c.JSON(200, res)
	// 	return
	// }

	// mediaType := types[0]
	// extension := filepath.Ext(file.Filename)
	// if mediaType == "image" {
	// 	dst += "images/"
	// }
	// filename := utils.CreateFileName(extension)
	// fileLocation := fmt.Sprintf("%v/%v", dst, filename)
	// err = c.SaveUploadedFile(file, fileLocation)
	// if err != nil {
	// 	defer os.Remove(fileLocation)
	// 	logger.Sugar.Error("Error Saving file", err)
	// 	res.ErrCode = 500
	// 	res.ErrMsg = utils.ChangeLang(lang, "ERROR_SAVING_FILE")
	// 	c.JSON(500, res)
	// 	return
	// }
	// url := fmt.Sprintf("https://%v/api/media/%v", c.Request.Host, filename)

	// media := model.Media{
	// 	FileName:  filename,
	// 	URL:       url,
	// 	Extension: strings.Replace(extension, ".", "", 1),
	// 	Type:      types[0],
	// 	AdminId:   admin.Id,
	// }

	// createMedia, err := ctr.mediaRepo.Create(&media)
	// if err != nil {
	// 	defer os.Remove(fileLocation)
	// 	logger.Sugar.Error("Error Reading file", err)
	// 	res.ErrCode = 500
	// 	res.ErrMsg = utils.ChangeLang(lang, "ERROR_READING_FILE")
	// 	c.JSON(200, res)
	// 	return
	// }

	// res.ErrCode = 0
	// res.ErrMsg = utils.ChangeLang(lang, "SUCCESS")
	// res.Data = gin.H{
	// 	"media": createMedia,
	// }
	// c.JSON(http.StatusOK, res)
}
