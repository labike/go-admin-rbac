package controller

import (
	"go-admin/api/entity"
	"go-admin/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sysPost entity.SysPost

// @Summary 新增岗位接口
// @Produce json
// @Description 新增岗位接口
// @Param data body entity.SysPost true "data"
// @Success 200 {object} result.Result
// @router /api/post/create [post]
func CreateSysPost(c *gin.Context) {
	_ = c.BindJSON(&sysPost)
	service.SysPostService().CreateSysPost(c, sysPost)
}

// @Summary 分页获取岗位列表
// @Produce json
// @Description 分页获取岗位列表
// @Param pageNum query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param postName query string false "岗位名称"
// @Param postStatus query string false "岗位状态:1->正常,2->禁用"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/post/list [get]
func GetSysPostList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	PostName := c.Query("postName")
	PostStatus := c.Query("postStatus")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysPostService().GetSysPostList(c, PageNum, PageSize, PostName, PostStatus, BeginTime, EndTime)
}

// @Summary 根据id查询岗位
// @Produce json
// @Description 根据id查询岗位
// @Param id query int false "岗位id"
// @Success 200 {object} result.Result
// @router /api/post/info [get]
func GetSysPostById(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysPostService().GetSysPostById(c, Id)
}

// @Summary 编辑岗位接口
// @Produce json
// @Description 编辑岗位接口
// @Param data body entity.SysPost true "data"
// @Success 200 {object} result.Result
// @router /api/post/update [put]
func UpdateSysPost(c *gin.Context) {
	_ = c.BindJSON(&sysPost)
	service.SysPostService().UpdateSysPost(c, sysPost)
}

// @Summary 删除岗位接口
// @Produce json
// @Description 删除岗位接口
// @Param data body entity.SysPostIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/post/delete [delete]
func DelSysPostById(c *gin.Context) {
	dto := entity.SysPostIdDto{}
	_ = c.BindJSON(&dto)
	service.SysPostService().DelSysPostById(c, dto)
}

// @Summary 批量删除岗位接口
// @Produce json
// @Description 批量删除岗位接口
// @Param data body entity.DelSysPostDto true "data"
// @Success 200 {object} result.Result
// @router /api/post/batch/delete [delete]
func BatchDelSysPost(c *gin.Context) {
	dto := entity.DelSysPostDto{}
	_ = c.BindJSON(&dto)
	service.SysPostService().BatchDeleteSysPost(c, dto)
}

// @Summary 编辑岗位状态
// @Produce json
// @Description 编辑岗位状态
// @Param data body entity.UpdateSysPostStatusDto true "data"
// @Success 200 {object} result.Result
// @router /api/post/updateStatus [put]
func UpdateSysPostStatus(c *gin.Context) {
	var dto entity.UpdateSysPostStatusDto
	_ = c.BindJSON(&dto)
	service.SysPostService().UpdateSysPostStatus(c, dto)
}

// @Summary 岗位下拉列表
// @Produce json
// @Description 岗位下拉列表
// @Success 200 {object} result.Result
// @router /api/post/vo/list [get]
func QuerySysPostList(c *gin.Context) {
	service.SysPostService().QuerySysPostList(c)
}
