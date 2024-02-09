package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"bluebell/settings"

	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() *gin.Engine {
	if settings.Config.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Config.Version)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/userInfo", middlewares.JWTAuthMiddleware(), func(ctx *gin.Context) {
		userID, username, err := controller.GetCurrentUser(ctx)
		if err != nil {
			controller.ResponseError(ctx, controller.CodeNeedLogin)
			return
		}
		controller.ResponseSuccess(ctx, gin.H{
			"user_id":  userID,
			"username": username,
		})
	})

	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "404"})
	})
	return r
}
