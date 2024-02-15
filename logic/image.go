package logic

import (
	"crypto/md5"
	"encoding/hex"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const fileNameSecret = "bluebellFile"

func UploadImage(c *gin.Context, file *multipart.FileHeader) (string, error) {
	fileBaseName := filepath.Base(file.Filename)
	fileExt := filepath.Ext(file.Filename)

	h := md5.New()
	h.Write([]byte(fileNameSecret))
	saveFilename := hex.EncodeToString(h.Sum([]byte(fileBaseName))) + fileExt
	// 保存文件到本地
	if err := c.SaveUploadedFile(file, "./file/"+saveFilename); err != nil {
		return "", err
	}
	return saveFilename, nil
}
