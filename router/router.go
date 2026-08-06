package router

import (
	"go-admin/api/controller"
	"go-admin/common/config"
	"go-admin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())
	router.StaticFS(config.Config.ImgSetting.UploadDir, http.Dir(config.Config.ImgSetting.UploadDir))
	router.Use(middleware.Logger())
	register(router)
	return router
}

func register(router *gin.Engine) {
	router.GET("/api/captcha", controller.Captcha)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/api/login", controller.Login)
}
