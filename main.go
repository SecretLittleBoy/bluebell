package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/routes"
	"bluebell/settings"
	"bluebell/pkg/snowflake"
	"bluebell/controller"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"context"

	"go.uber.org/zap"
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Printf("Fatal error settings.Init: %s \n", err)
		return
	}

	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	if err := snowflake.Init(); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	if err := redis.Init(); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()	

	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init gin translator failed, err:%v\n", err)
	}

	r := routes.Init()
	fmt.Println(settings.Config.AppConfig.Port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Config.AppConfig.Port),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen: ", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown:", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
