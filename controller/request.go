package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

const CtxUserIDKey = "user_id"
const CtxUsernameKey = "username"

func GetCurrentUser(c *gin.Context) (user_id int64, username string, err error) {
	user_id_any_type, ok := c.Get(CtxUserIDKey)
	if !ok {
		return 0, "", ErrorUserNotLogin
	}
	user_id, ok = user_id_any_type.(int64)
	if !ok {
		return 0, "", ErrorUserNotLogin
	}
	username_any_type, ok := c.Get(CtxUsernameKey)
	if !ok {
		return 0, "", ErrorUserNotLogin
	}
	username, ok = username_any_type.(string)
	if !ok {
		return 0, "", ErrorUserNotLogin
	}
	return user_id, username, nil
}
