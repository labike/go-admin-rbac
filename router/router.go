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
	jwt := router.Group("/api", middleware.AuthMiddleware(), middleware.LogMiddleware())
	{
		jwt.POST("/post/create", controller.CreateSysPost)
		jwt.GET("/post/list", controller.GetSysPostList)
		jwt.GET("/post/info", controller.GetSysPostById)
		jwt.PUT("/post/update", controller.UpdateSysPost)
		jwt.DELETE("/post/delete", controller.DelSysPostById)
		jwt.DELETE("/post/batch/delete", controller.BatchDelSysPost)
		jwt.PUT("/post/updateStatus", controller.UpdateSysPostStatus)
		jwt.GET("/post/vo/list", controller.QuerySysPostList)
		jwt.GET("/dept/list", controller.GetSysDeptList)
		jwt.POST("/dept/create", controller.CreateSysDept)
		jwt.GET("/dept/info", controller.GetSysDeptById)
		jwt.PUT("/dept/update", controller.UpdateSysDept)
		jwt.DELETE("/dept/delete", controller.DeleteSysDeptById)
		jwt.GET("/dept/vo/list", controller.QuerySysDeptList)
		jwt.POST("/menu/create", controller.CreateSysMenu)
		jwt.GET("/menu/vo/list", controller.QuerySysMenuVoList)
		jwt.GET("/menu/info", controller.GetSysMenuById)
		jwt.PUT("/menu/update", controller.UpdateSysMenu)
		jwt.DELETE("/menu/delete", controller.DeleteSysMenu)
		jwt.GET("/menu/list", controller.GetSysMenuList)
		jwt.POST("/role/create", controller.CreateSysRole)
		jwt.GET("/role/info", controller.GetSysRoleById)
		jwt.PUT("/role/edit", controller.UpdateSysRole)
		jwt.DELETE("/role/delete", controller.DeleteSysRoleById)
		jwt.PUT("/role/update", controller.UpdateSysRoleStatus)
		jwt.GET("/role/list", controller.GetSysRoleList)
		jwt.GET("/role/vo/list", controller.GetSysRoleVoList)
		jwt.GET("/role/vo/idList", controller.QueryRoleMenuIdList)
		jwt.PUT("/role/assignPermissions", controller.AssignPermission)
		jwt.POST("/admin/create", controller.CreateSysAdmin)
		jwt.GET("/admin/info", controller.GetSysAdminInfo)
		jwt.PUT("/admin/edit", controller.UpdateSysAdmin)
		jwt.DELETE("/admin/delete", controller.DeleteSysAdmin)
		jwt.PUT("/admin/status", controller.UpdateSysAdminStatus)
		jwt.POST("/admin/reset/passwrod", controller.UpdateSysAdminPwd)
		jwt.GET("/admin/list", controller.GetSysAdminList)
		jwt.POST("/upload", controller.Upload)
		jwt.PUT("/user/update", controller.UpdateUser)
		jwt.PUT("/user/update/password", controller.UpdatePersonPwd)
		jwt.GET("/login/info", controller.GetLoginLog)
		jwt.DELETE("/login/delete", controller.DeleteLoginLog)
		jwt.DELETE("/login/batchDelete", controller.BatchDeleteLoginLog)
		jwt.DELETE("/login/clear", controller.ClearLoginLog)
		jwt.GET("/oper/list", controller.GetSysOperList)
		jwt.DELETE("/operr/delete", controller.DeleteSysOperLogById)
		jwt.DELETE("/oper/batchDelete", controller.BatchDeleteSysOperLog)
	}
}
