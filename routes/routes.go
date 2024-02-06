package routes

import (
	"bluebell/settings"
	"bluebell/logger"

	"net/http"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	// Set the mode to release mode
	gin.SetMode(settings.Config.Mode)
	r := gin.New()
	// Add a logger middleware
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// Register the routes
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Config.Version)
	})
	return r
}