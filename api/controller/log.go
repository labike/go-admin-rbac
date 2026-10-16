package controller

import (
	"go-admin/api/entity"
	"go-admin/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 分页获取登录日志列表
// @Produce json
// @Description 分页获取登录日志列表
// @Param pageNum query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param userName query string false "用户名称"
// @Param loginStatus query string false "状态:1->正常,2->禁用"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/login/info [get]
// @Security ApiKeyAuth
func GetLoginLog(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	Username := c.Query("username")
	LoginStatus := c.Query("loginStatus")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysLoginService().GetLoginLog(c, Username, LoginStatus, BeginTime, EndTime, PageSize, PageNum)
}

// @Summary 根据id删除登录日志
// @Produce json
// @Description 根据id删除登录日志
// @Param data body entity.SysLoginDto true "data"
// @Success 200 {object} result.Result
// @router /api/login/delete [delete]
// @Security ApiKeyAuth
func DeleteLoginLog(c *gin.Context) {
	var dto entity.SysLoginIdDto
	_ = c.BindJSON(&dto)
	service.SysLoginService().DeleteLoginLog(c, dto)
}

// @Summary 根据id批量删除登录日志
// @Produce json
// @Description 根据id批量删除登录日志
// @Param data body entity.SysLoginIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/login/batchDelete [delete]
// @Security ApiKeyAuth
func BatchDeleteLoginLog(c *gin.Context) {
	var dto entity.SysLoginDto
	_ = c.BindJSON(&dto)
	service.SysLoginService().BatchDeleteLoginLog(c, dto)
}

// @Summary 清空登录日志
// @Produce json
// @Description 清空登录日志
// @Success 200 {object} result.Result
// @router /api/login/clear [delete]
// @Security ApiKeyAuth
func ClearLoginLog(c *gin.Context) {
	service.SysLoginService().ClearLoginLog(c)
}
