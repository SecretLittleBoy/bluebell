package controller

import (
	"bluebell/logic"
	"bluebell/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		err_in_ValidationErrors_type, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(err_in_ValidationErrors_type.Translate(trans)))
		}
		return
	}

	if err := logic.SignUp(&p); err != nil {
		zap.L().Error("SignUp failed", zap.Error(err))
		ResponseError(c, CodeUserExist)
		return
	}

	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	var p models.ParamLogin
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		err_in_ValidationErrors_type, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(err_in_ValidationErrors_type.Translate(trans)))
		}
		return
	}
	token, err := logic.Login(&p)
	if err != nil {
		zap.L().Error("Login failed", zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	}
	ResponseSuccess(c, token)
}
