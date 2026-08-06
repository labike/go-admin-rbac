package controller

import (
	"go-admin/api/entity"
	"go-admin/api/service"

	"github.com/gin-gonic/gin"
)

// @Summary 登录
// @Produce json
// @Description 登录
// @Param data body entity.LoginDto true "data"
// @Success 200 {object} result.Result
// @router /api/login [post]
func Login(c *gin.Context) {
	var dto entity.LoginDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().Login(c, dto)
}
