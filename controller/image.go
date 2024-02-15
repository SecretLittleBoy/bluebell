package controller

import (
	"bluebell/logic"
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
