package service

import (
	"go-admin/api/dao"
	"go-admin/api/entity"
	"go-admin/common/result"

	"github.com/gin-gonic/gin"
)

// 接口只声明方法签名，不关心具体实现
type ISysLoginService interface {
	GetLoginLog(c *gin.Context, Username, LoginStatus, BeginTime, EndTime string, PageSize, PageNum int)
	BatchDeleteLoginLog(c *gin.Context, dto entity.SysLoginDto)
	DeleteLoginLog(c *gin.Context, dto entity.SysLoginIdDto)
	ClearLoginLog(c *gin.Context)
}

// 怎么做
type SysLoginServiceImpl struct{}

// 登录日志
func (s SysLoginServiceImpl) GetLoginLog(c *gin.Context, Username, LoginStatus, BeginTime, EndTime string, PageSize, PageNum int) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysLoginInfo, count := dao.GetLoginLog(Username, LoginStatus, BeginTime, EndTime, PageSize, PageNum)
	result.Success(c, map[string]interface{}{
		"total":    count,
		"PageSize": PageSize,
		"PageNum":  PageNum,
		"list":     sysLoginInfo,
	})
}

func (s SysLoginServiceImpl) BatchDeleteLoginLog(c *gin.Context, dto entity.SysLoginDto) {
	dao.BatchDeleteLoginLog(dto)
	result.Success(c, true)
}

func (s SysLoginServiceImpl) DeleteLoginLog(c *gin.Context, dto entity.SysLoginIdDto) {
	dao.DeleteLoginById(dto)
	result.Success(c, true)
}

func (s SysLoginServiceImpl) ClearLoginLog(c *gin.Context) {
	dao.ClearLoginLog()
	result.Success(c, true)
}

// 单例对象，接口的实现实例
var sysLoginService = SysLoginServiceImpl{}

// 单例的获取入口
func SysLoginService() ISysLoginService {
	return &sysLoginService
}
