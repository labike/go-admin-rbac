package controller

import (
	"go-admin/api/entity"
	"go-admin/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sysMenu entity.SysMenu

// @Summary 创建菜单
// @Produce json
// @Description 创建菜单
// @Param data body entity.SysMenu true "data"
// @Success 200 {object} result.Result
// @router /api/menu/create [post]
// @Security ApiKeyAuth
func CreateSysMenu(c *gin.Context) {
	_ = c.BindJSON(&sysMenu)
	service.SysMenuService().CreateSysMenu(c, sysMenu)
}

// @Summary 获取菜单下拉列表
// @Produce json
// @Description 获取菜单下拉列表
// @Success 200 {object} result.Result
// @router /api/menu/vo/list [get]
// @Security ApiKeyAuth
func QuerySysMenuVoList(c *gin.Context) {
	service.SysMenuService().QuerySysMenuVoList(c)
}

// @Summary 根据菜单id查询菜单信息
// @Produce json
// @Description 根据菜单id查询菜单信息
// @Param id query int true "id"
// @Success 200 {object} result.Result
// @router /api/menu/info [get]
// @Security ApiKeyAuth
func GetSysMenuById(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysMenuService().GetSysMenuById(c, Id)
}

// @Summary 编辑菜单
// @Produce json
// @Description 编辑菜单
// @Param data body entity.SysMenu true "data"
// @Success 200 {object} result.Result
// @router /api/menu/update [put]
// @Security ApiKeyAuth
func UpdateSysMenu(c *gin.Context) {
	_ = c.BindJSON(&sysMenu)
	service.SysMenuService().UpdateSysMenu(c, sysMenu)
}

// @Summary 删除菜单
// @Produce json
// @Description 删除菜单
// @Param data body entity.SysMenuIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/menu/delete [delete]
// @Security ApiKeyAuth
func DeleteSysMenu(c *gin.Context) {
	var dto entity.SysMenuIdDto
	_ = c.BindJSON(&dto)
	service.SysMenuService().DeleteSysMenu(c, dto)
}

// @Summary 查询菜单
// @Produce json
// @Description 查询菜单
// @Param menuName query string false "菜单名称"
// @Param menuStatus query string false "菜单状态"
// @Success 200 {object} result.Result
// @router /api/menu/list [get]
// @Security ApiKeyAuth
func GetSysMenuList(c *gin.Context) {
	MenuName := c.Query("menuName")
	MenuStatus := c.Query("menuStatus")
	service.SysMenuService().GetSysMenuList(c, MenuName, MenuStatus)
}
