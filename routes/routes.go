package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"bluebell/settings"

	"net/http"

	"github.com/gin-gonic/gin"
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

	v1 := r.Group("/api/v1")
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware())

	v1.GET("/community", controller.CommunityHandler)
	v1.GET("/community/:id", controller.CommunityDetailHandler)
	v1.GET("/userInfo", controller.UserInfoHander)
	v1.POST("/refreshToken", controller.RefreshTokenHandler)
	v1.POST("/post", controller.CreatePostHandler)
	v1.GET("/post/:id", controller.GetPostDetailHandler)
	v1.GET("/post/", controller.GetPostListHandler)
	v1.POST("/vote", controller.PostVoteController)

	v2 := r.Group("/api/v2")
	v2.Use(middlewares.JWTAuthMiddleware())
	v2.GET("/post", controller.GetPostListHandler2)//可以按时间或者分数排序
	v2.GET("/community/:id/post", controller.GetCommunityPostListHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "404"})
	})
	return r
}
