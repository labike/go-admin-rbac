package service

import (
	"go-admin/api/dao"
	"go-admin/api/entity"
	"go-admin/common/result"

	"github.com/gin-gonic/gin"
)

type ISysMenuService interface {
	GetSysMenuByName(c *gin.Context, menuName string)
	CreateSysMenu(c *gin.Context, addMenu entity.SysMenu)
	QuerySysMenuVoList(c *gin.Context)
	GetSysMenuById(c *gin.Context, Id int)
	UpdateSysMenu(c *gin.Context, menu entity.SysMenu)
	DeleteSysMenu(c *gin.Context, dto entity.SysMenuIdDto)
	GetSysMenuList(c *gin.Context, MenuName string, MenuStatus string)
}

type SysMenuServiceImpl struct{}

var sysMenuService = SysMenuServiceImpl{}

func (s *SysMenuServiceImpl) GetSysMenuByName(c *gin.Context, menuName string) {
	res := dao.GetSysMenuByName(menuName)
	result.Success(c, res)
}

func (s *SysMenuServiceImpl) CreateSysMenu(c *gin.Context, addMenu entity.SysMenu) {
	res := dao.CreateSysMenu(addMenu)
	if !res {
		result.Failed(c, int(result.ApiCode.MENUEXISTS), result.ApiCode.GetMessage(result.ApiCode.MENUEXISTS))
		return
	}
	result.Success(c, true)
}

func (s *SysMenuServiceImpl) QuerySysMenuVoList(c *gin.Context) {
	res := dao.QueryDeptList()
	result.Success(c, res)
}

func (s *SysMenuServiceImpl) GetSysMenuById(c *gin.Context, Id int) {
	res := dao.GetSysMenuById(Id)
	result.Success(c, res)
}

func (s *SysMenuServiceImpl) UpdateSysMenu(c *gin.Context, menu entity.SysMenu) {
	res := dao.UpdateSysMenu(menu)
	result.Success(c, res)
}

func (s *SysMenuServiceImpl) DeleteSysMenu(c *gin.Context, dto entity.SysMenuIdDto) {
	res := dao.DeleteSysMenu(dto)
	if !res {
		result.Failed(c, int(result.ApiCode.DELSYSMENUFAILED), result.ApiCode.GetMessage(result.ApiCode.DELSYSMENUFAILED))
		return
	}
	result.Success(c, true)
}

func (s *SysMenuServiceImpl) GetSysMenuList(c *gin.Context, MenuName string, MenuStatus string) {
	res := dao.GetSysMenuList(MenuName, MenuStatus)
	result.Success(c, res)
}

func SysMenuService() ISysMenuService {
	return &sysMenuService
}
