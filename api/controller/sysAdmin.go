package controller

import (
	"go-admin/api/entity"
	"go-admin/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 登录
// @Produce json
// @Description 登录
// @Param data body entity.LoginDto true "data"
// @Success 200 {object} result.Result
// @router /api/login [post]
// @Security ApiKeyAuth
func Login(c *gin.Context) {
	var dto entity.LoginDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().Login(c, dto)
}

// @Summary 新增用户
// @Produce json
// @Description 新增用户
// @Param data body entity.AddSysAdminDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/create [post]
// @Security ApiKeyAuth
func CreateSysAdmin(c *gin.Context) {
	var dto entity.AddSysAdminDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().CreateSysAdmin(c, dto)
}

// @Summary 根据id查询用户
// @Produce json
// @Description 根据id查询用户
// @Param id query int true "Id"
// @Success 200 {object} result.Result
// @router /api/admin/info [get]
// @Security ApiKeyAuth
func GetSysAdminInfo(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysAdminService().GetSysAdminInfo(c, Id)
}

// @Summary 编辑用户
// @Produce json
// @Description 编辑用户
// @Param data body entity.UpdateSysAdminDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/edit [put]
// @Security ApiKeyAuth
func UpdateSysAdmin(c *gin.Context) {
	var dto entity.UpdateSysAdminDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdateSysAdmin(c, dto)
}

// @Summary 根据用户id删除用户
// @Produce json
// @Description 根据用户id删除用户
// @Param data body entity.SysAdminIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/delete [delete]
// @Security ApiKeyAuth
func DeleteSysAdmin(c *gin.Context) {
	var dto entity.SysAdminIdDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().DeleteSysAdmin(c, dto)
}

// @Summary 编辑用户状态
// @Produce json
// @Description 编辑用户状态
// @Param data body entity.UpdateSysAdminStatusDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/status [put]
// @Security ApiKeyAuth
func UpdateSysAdminStatus(c *gin.Context) {
	var dto entity.UpdateSysAdminStatusDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdateSysAdminStatus(c, dto)
}

// @Summary 修改用户密码
// @Produce json
// @Description 修改用户密码
// @Param data body entity.ResetSysAdminPasswordDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/reset/password [post]
// @Security ApiKeyAuth
func UpdateSysAdminPwd(c *gin.Context) {
	var dto entity.ResetSysAdminPasswordDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdateSysAdminPwd(c, dto)
}

// @Summary 分页获取用户列表
// @Produce json
// @Description 分页获取用户列表
// @Param pageNum query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param userName query string false "用户名称"
// @Param status query string false "状态:1->正常,2->禁用"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/admin/list [get]
// @Security ApiKeyAuth
func GetSysAdminList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	Username := c.Query("username")
	Status := c.Query("status")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysAdminService().GetSysAdminList(c, PageSize, PageNum, Username, Status, BeginTime, EndTime)
}

// @Summary 编辑用户
// @Produce json
// @Description 编辑用户
// @Param data body entity.UpdateUserDto true "data"
// @Success 200 {object} result.Result
// @router /api/user/update [put]
// @Security ApiKeyAuth
func UpdateUser(c *gin.Context) {
	var dto entity.UpdateUserDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdateUser(c, dto)
}

// @Summary 更新当前用户密码
// @Produce json
// @Description 更新当前用户密码
// @Param data body entity.UpdatePersonPwdDto true "data"
// @Success 200 {object} result.Result
// @router /api/user/update/password [put]
// @Security ApiKeyAuth
func UpdatePersonPwd(c *gin.Context) {
	var dto entity.UpdatePersonPwdDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdatePersonPwd(c, dto)
}
