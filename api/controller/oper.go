package controller

import (
	"go-admin/api/entity"
	"go-admin/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 分页获取操作日志列表
// @Produce json
// @Description 分页获取操作日志列表
// @Param pageNum query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param userName query string false "用户名称"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/oper/list [get]
// @Security ApiKeyAuth
func GetSysOperList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	Username := c.Query("username")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysOperService().GetSysOperList(c, Username, BeginTime, EndTime, PageSize, PageNum)
}

// @Summary 根据id删除操作日志
// @Produce json
// @Description 根据id删除操作日志
// @Param data body entity.SysOperateLogDto true "data"
// @Success 200 {object} result.Result
// @router /api/oper/delete [delete]
// @Security ApiKeyAuth
func DeleteSysOperLogById(c *gin.Context) {
	var dto entity.SysOperateLogDto
	_ = c.BindJSON(&dto)
	service.SysOperService().DeleteSysOperLogById(c, dto)
}

// @Summary 根据id批量删除操作日志
// @Produce json
// @Description 根据id批量删除操作日志
// @Param data body entity.SysOperateIdsDto true "data"
// @Success 200 {object} result.Result
// @router /api/oper/batchDelete [delete]
// @Security ApiKeyAuth
func BatchDeleteSysOperLog(c *gin.Context) {
	var dto entity.SysOperateIdsDto
	_ = c.BindJSON(&dto)
	service.SysOperService().BatchDeleteSysOperLog(c, dto)
}
