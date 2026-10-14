package controller

import (
	"go-admin/api/entity"
	"go-admin/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sysRole entity.SysRole

// @Summary 创建角色
// @Produce json
// @Description 创建角色
// @Param data body entity.SysRole true "data"
// @Success 200 {object} result.Result
// @router /api/role/create [post]
// @Security ApiKeyAuth
func CreateSysRole(c *gin.Context) {
	var dto entity.AddSysRole
	_ = c.BindJSON(&dto)
	service.SysRoleService().CreateSysRole(c, dto)
}

// @Summary 根据角色id查询角色信息
// @Produce json
// @Description 根据角色id查询角色信息
// @Param id query int true "id"
// @Success 200 {object} result.Result
// @router /api/role/info [get]
// @Security ApiKeyAuth
func GetSysRoleById(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysRoleService().GetSysRoleById(c, Id)
}

// @Summary 编辑角色
// @Produce json
// @Description 编辑角色
// @Param data body entity.UpdateSysRole true "data"
// @Success 200 {object} result.Result
// @router /api/role/edit [put]
// @Security ApiKeyAuth
func UpdateSysRole(c *gin.Context) {
	var dto entity.UpdateSysRole
	_ = c.BindJSON(&dto)
	service.SysRoleService().UpdateSysRole(c, dto)
}

// @Summary 删除角色
// @Produce json
// @Description 删除角色
// @Param data body entity.SysRoleDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/delete [delete]
// @Security ApiKeyAuth
func DeleteSysRoleById(c *gin.Context) {
	var dto entity.SysRoleDto
	_ = c.BindJSON(&dto)
	service.SysRoleService().DeleteSysRoleById(c, dto)
}

// @Summary 编辑角色状态
// @Produce json
// @Description 编辑角色状态
// @Param data body entity.UpdateSysRoleStatus true "data"
// @Success 200 {object} result.Result
// @router /api/role/update [put]
// @Security ApiKeyAuth
func UpdateSysRoleStatus(c *gin.Context) {
	var dto entity.UpdateSysRoleStatus
	_ = c.BindJSON(&dto)
	service.SysRoleService().UpdateSysRoleStatus(c, dto)
}

// @Summary 获取角色分页数据
// @Produce json
// @Description 获取角色分页数据
// @Param pageNum query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param roleName query string false "角色名称"
// @Param status query string false "角色状态"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/role/info [get]
// @Security ApiKeyAuth
func GetSysRoleList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	Status := c.Query("status")
	RoleName := c.Query("roleName")
	service.SysRoleService().GetSysRoleList(c, PageNum, PageSize, RoleName, Status, BeginTime, EndTime)
}

// @Summary 获取角色下拉列表数据
// @Produce json
// @Description 获取角色下拉列表数据
// @Success 200 {object} result.Result
// @router /api/role/vo/list [get]
// @Security ApiKeyAuth
func GetSysRoleVoList(c *gin.Context) {
	service.SysRoleService().GetSysRoleVoList(c)
}

// @Summary 根据角色id获取菜单权限数据
// @Produce json
// @Description 根据角色id获取菜单权限数据
// @Param id query int true "角色id"
// @Success 200 {object} result.Result
// @router /api/role/vo/idList [get]
// @Security ApiKeyAuth
func QueryRoleMenuIdList(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysRoleService().QueryRoleMenuIdList(c, Id)
}

// @Summary 分配权限
// @Produce json
// @Description 分配权限
// @Param data body entity.RoleMenu true "data"
// @Success 200 {object} result.Result
// @router /api/role/assignPermissions [put]
// @Security ApiKeyAuth
func AssignPermission(c *gin.Context) {
	var RoleMenu entity.RoleMenu
	_ = c.BindJSON(&RoleMenu)
	service.SysRoleService().AssignPermission(c, RoleMenu)
}
