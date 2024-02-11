package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"strconv"

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
	access_token, refresh_token, err := logic.Login(&p)
	if err != nil {
		zap.L().Error("Login failed", zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	}
	ResponseSuccess(c, gin.H{
		"access_token":  access_token,
		"refresh_token": refresh_token})
}

func UserInfoHander(ctx *gin.Context) {
	userID, username, err := GetCurrentUser(ctx)
	if err != nil {
		ResponseError(ctx, CodeNeedLogin)
		return
	}
	ResponseSuccess(ctx, gin.H{
		"user_id":  strconv.FormatInt(userID, 10),
		"username": username,
	})
}

func RefreshTokenHandler(ctx *gin.Context) {
	var p models.ParamRefreshToken
	if err := ctx.ShouldBindJSON(&p); err != nil {
		zap.L().Error("RefreshToken with invalid param", zap.Error(err))
		err_in_ValidationErrors_type, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(err_in_ValidationErrors_type.Translate(trans)))
		}
		return
	}
	newAccessToken, err := logic.RefreshToken(&p)
	if err != nil {
		zap.L().Error("RefreshToken failed", zap.Error(err))
		ResponseError(ctx, CodeInvalidToken)
		return
	}
	ResponseSuccess(ctx, gin.H{
		"access_token": newAccessToken,
	})
}
