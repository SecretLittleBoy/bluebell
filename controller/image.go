package controller

import (
	"bluebell/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImageHandler(c *gin.Context) {
	file, err := c.FormFile("image") // 从请求中解析文件
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	if savedFileName, err := logic.UploadImage(c, file); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	} else {
		ResponseSuccess(c, gin.H{
			"file_name": savedFileName,
		})
	}
}

func GetImageHandler(c *gin.Context) {
	fileName := c.Param("filename")
	if fileName == "" {
		ResponseError(c, CodeInvalidParam)
		return
	}
	file, err := logic.GetImage(fileName)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	c.Data(http.StatusOK, "image/jpg", file)
}