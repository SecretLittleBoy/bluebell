package logic

import (
	"crypto/md5"
	"encoding/hex"
	"mime/multipart"
	"path/filepath"

	"bluebell/dao/oss"

	"github.com/gin-gonic/gin"
)

const fileNameSecret = "bluebellFile"

func UploadImage(c *gin.Context, file *multipart.FileHeader) (string, error) {
	fileBaseName := filepath.Base(file.Filename)
	fileExt := filepath.Ext(file.Filename)

	h := md5.New()
	h.Write([]byte(fileNameSecret))
	saveFilename := hex.EncodeToString(h.Sum([]byte(fileBaseName))) + fileExt
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	return oss.UploadFileFromReader(saveFilename, src)
}

func GetImage(fileName string) ([]byte, error) {
	return oss.GetObject(fileName)
}
