package oss

import (
	"bluebell/settings"
	"errors"
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var ossClient *oss.Client
var bucket *oss.Bucket

func Init() (err error) {
	ossClient, err = oss.New(settings.Config.OssConfig.Endpoint, settings.Config.OssConfig.AccessKeyId, settings.Config.OssConfig.AccessKeySecret)
	if err != nil {
		return
	}
	isBucketExist, err := ossClient.IsBucketExist(settings.Config.OssConfig.Bucket)
	if err != nil {
		return
	}
	if !isBucketExist {
		err = errors.New("bucket does not exist")
		return
	}
	bucket, err = ossClient.Bucket(settings.Config.OssConfig.Bucket)
	return
}

func UploadFileFromLocal(newName string, localPath string) (string, error) {
	err := bucket.PutObjectFromFile(newName, localPath)
	if err != nil {
		return "", err
	}
	return settings.Config.OssConfig.Bucket + "." + settings.Config.OssConfig.Endpoint + "/" + newName, nil
}

func UploadFileFromReader(newName string, reader io.Reader) (string, error) {
	err := bucket.PutObject(newName, reader)
	if err != nil {
		return "", err
	}
	return settings.Config.OssConfig.Bucket + "." + settings.Config.OssConfig.Endpoint + "/" + newName, nil
}
