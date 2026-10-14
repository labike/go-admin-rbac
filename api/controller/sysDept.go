package controller

import (
	"go-admin/api/entity"
	"go-admin/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sysDept entity.SysDept

// @Summary 获取部门列表
// @Produce json
// @Description 获取部门列表
// @Param deptName query string false "部门名称"
// @Param deptStatus query string false "部门状态"
// @Success 200 {object} result.Result
// @router /api/dept/list [get]
// @Security ApiKeyAuth
func GetSysDeptList(c *gin.Context) {
	DeptName := c.Query("deptName")
	DeptStatus := c.Query("deptStatus")
	service.SysDeptService().GetSysDeptList(c, DeptName, DeptStatus)
}

// @Summary 创建部门
// @Produce json
// @Description 创建部门
// @Param data body entity.CreateSysDeptDto true "data"
// @Success 200 {object} result.Result
// @router /api/dept/create [post]
// @Security ApiKeyAuth
func CreateSysDept(c *gin.Context) {
	_ = c.BindJSON(&sysDept)
	service.SysDeptService().CreateSysDept(c, sysDept)
}

// @Summary 根据部门id查询
// @Produce json
// @Description 根据部门id查询
// @Param id query int false "部门id"
// @Success 200 {object} result.Result
// @router /api/dept/info [get]
// @Security ApiKeyAuth
func GetSysDeptById(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysDeptService().GetSysDeptById(c, Id)
}

// @Summary 编辑部门
// @Produce json
// @Description 编辑部门
// @Param data body entity.SysDept true "data"
// @Success 200 {object} result.Result
// @router /api/dept/update [PUT]
// @Security ApiKeyAuth
func UpdateSysDept(c *gin.Context) {
	_ = c.ShouldBindJSON(&sysDept)
	service.SysDeptService().UpdateSysDept(c, sysDept)
}

// @Summary 删除部门
// @Produce json
// @Description 删除部门
// @Param data body entity.SysDeptIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/dept/delete [delete]
// @Security ApiKeyAuth
func DeleteSysDeptById(c *gin.Context) {
	var dto entity.SysDeptIdDto
	_ = c.BindJSON(&dto)
	service.SysDeptService().DeleteSysDeptById(c, dto)
}

// @Summary 部门下拉列表
// @Produce json
// @Description 部门下拉列表
// @Success 200 {object} result.Result
// @router /api/dept/vo/list [get]
// @Security ApiKeyAuth
func QuerySysDeptList(c *gin.Context) {
	service.SysDeptService().QuerySysDeptList(c)
}
