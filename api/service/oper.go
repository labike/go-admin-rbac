package service

import (
	"go-admin/api/dao"
	"go-admin/api/entity"
	"go-admin/common/result"

	"github.com/gin-gonic/gin"
)

// 接口只声明方法签名，不关心具体实现
type ISysOperService interface {
	GetSysOperList(c *gin.Context, Username string, BeginTime string, EndTime string, PageSize int, PageNum int)
	BatchDeleteSysOperLog(c *gin.Context, dto entity.SysOperateIdsDto)
	DeleteSysOperLogById(c *gin.Context, dto entity.SysOperateLogDto)
}

// 怎么做
type SysOperServiceImpl struct{}

// 登录日志
func (s SysOperServiceImpl) GetSysOperList(c *gin.Context, Username string, BeginTime string, EndTime string, PageSize int, PageNum int) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysLoginInfo, count := dao.GetSysOperList(Username, BeginTime, EndTime, PageSize, PageNum)
	result.Success(c, map[string]interface{}{
		"total":    count,
		"PageSize": PageSize,
		"PageNum":  PageNum,
		"list":     sysLoginInfo,
	})
}

func (s SysOperServiceImpl) BatchDeleteSysOperLog(c *gin.Context, dto entity.SysOperateIdsDto) {
	dao.BatchDeleteSysOperLog(dto)
	result.Success(c, true)
}

func (s SysOperServiceImpl) DeleteSysOperLogById(c *gin.Context, dto entity.SysOperateLogDto) {
	dao.DeleteSysOperLogById(dto)
	result.Success(c, true)
}

// 单例对象，接口的实现实例
var sysOperService = SysOperServiceImpl{}

// 单例的获取入口
func SysOperService() ISysOperService {
	return &sysOperService
}
