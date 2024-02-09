package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"errors"

	"go.uber.org/zap"
)

func SignUp(param *models.ParamSignUp) (err error) {
	var isExist bool
	isExist, err = mysql.IsUserExist(param.Username)
	if err != nil {
		zap.L().Error("CheckUserExist failed", zap.Error(err))
		return
	} else if isExist {
		zap.L().Warn("user already exists")
		err = errors.New("user already exists")
		return
	}
	newUserID := snowflake.GenID()
	newUser := &models.User{
		UserID:   newUserID,
		Username: param.Username,
		Password: param.Password,
	}
	mysql.InsertUser(newUser)
	return
}

func Login(param *models.ParamLogin) (access_token, refresh_token string, err error) {
	userToLogin := &models.User{
		Username: param.Username,
		Password: param.Password,
	}
	err = mysql.Login(userToLogin)
	if err != nil {
		return "", "", err
	} else {
		return jwt.GenFullToken(userToLogin.UserID, userToLogin.Username)
	}
}

func RefreshToken(param *models.ParamRefreshToken) (newAccessToken string, err error) {
	newAccessToken, err = jwt.RefreshToken(param.AcessToken, param.RefreshToken)
	return
}
