package main

import (
	"context"
	"go-admin/common/config"
	_ "go-admin/docs"
	"go-admin/pkg/db"
	"go-admin/pkg/log"
	"go-admin/pkg/redis"
	"go-admin/router"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// @title 通用后台管理系统
// @version 1.0
// @description 后台管理系统api接口文档
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 加载日志
	log := log.Log()
	gin.SetMode(config.Config.Server.Mode)

	router := router.InitRouter()

	srv := &http.Server{
		Addr:    config.Config.Server.Address,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Info("listen: %s\n", err)
		}
		log.Info("listen: %s\n", config.Config.Server.Address)
	}()

	quit := make(chan os.Signal)
	<-quit
	log.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Info("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}

func init() {
	// mysql
	db.SetupDB()
	// redis
	redis.SetupRedis()

}
