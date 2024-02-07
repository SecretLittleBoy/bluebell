package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"net/http"

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
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(err_in_ValidationErrors_type.Translate(trans)),
			})
		}
		return
	}

	if err := logic.SignUp(&p); err != nil {
		zap.L().Error("SignUp failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "signup failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "signup success",
	})
}

func LoginHandler(c *gin.Context) {
	var p models.ParamLogin
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		err_in_ValidationErrors_type, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(err_in_ValidationErrors_type.Translate(trans)),
			})
		}
		return
	}

	if err := logic.Login(&p); err != nil {
		zap.L().Error("Login failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "login failed:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "login success"})
}
