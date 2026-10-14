package service

import (
	"go-admin/api/dao"
	"go-admin/api/entity"
	"go-admin/common/result"

	"github.com/gin-gonic/gin"
)

type ISysRoleService interface {
	CreateSysRole(c *gin.Context, dto entity.AddSysRole)
	GetSysRoleById(c *gin.Context, Id int)
	UpdateSysRole(c *gin.Context, Dto entity.UpdateSysRole)
	UpdateSysRoleStatus(c *gin.Context, dto entity.UpdateSysRoleStatus)
	DeleteSysRoleById(c *gin.Context, dto entity.SysRoleDto)
	GetSysRoleList(c *gin.Context, PageNum, PageSize int, RoleName, Status, BeginTime, EndTime string)
	GetSysRoleVoList(c *gin.Context)
	QueryRoleMenuIdList(c *gin.Context, Id int)
	AssignPermission(c *gin.Context, menu entity.RoleMenu)
}

type SysRoleServiceImpl struct{}

var sysRoleService = SysRoleServiceImpl{}

func (s *SysRoleServiceImpl) CreateSysRole(c *gin.Context, dto entity.AddSysRole) {
	res := dao.CreateSysRole(dto)
	if !res {
		result.Failed(c, int(result.ApiCode.ROLEEXISTS), result.ApiCode.GetMessage(result.ApiCode.ROLEEXISTS))
		return
	}
	result.Success(c, true)
}

func (s *SysRoleServiceImpl) GetSysRoleById(c *gin.Context, Id int) {
	res := dao.GetSysRoleById(Id)
	result.Success(c, res)
}

func (s *SysRoleServiceImpl) DeleteSysRoleById(c *gin.Context, dto entity.SysRoleDto) {
	dao.DeleteSysRoleById(dto)
	result.Success(c, true)
}

func (s *SysRoleServiceImpl) UpdateSysRole(c *gin.Context, dto entity.UpdateSysRole) {
	res := dao.UpdateSysRole(dto)
	result.Success(c, res)
}

func (s *SysRoleServiceImpl) UpdateSysRoleStatus(c *gin.Context, dto entity.UpdateSysRoleStatus) {
	res := dao.UpdateSysRoleStatus(dto)
	result.Success(c, res)
}

func (s *SysRoleServiceImpl) GetSysRoleList(c *gin.Context, PageNum, PageSize int, RoleName, Status, BeginTime, EndTime string) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	res, count := dao.GetSysRoleList(PageNum, PageSize, RoleName, Status, BeginTime, EndTime)
	result.Success(c, map[string]interface{}{
		"list":     res,
		"total":    count,
		"pageSize": PageSize,
		"pageNum":  PageNum,
	})
	return
}

func (s *SysRoleServiceImpl) GetSysRoleVoList(c *gin.Context) {
	res := dao.GetSysRoleVoList()
	result.Success(c, res)
}

func (s *SysRoleServiceImpl) QueryRoleMenuIdList(c *gin.Context, Id int) {
	res := dao.QueryRoleMenuIdList(Id)
	var idList = make([]int, 0)
	for _, id := range res {
		idList = append(idList, id.Id)
	}
	result.Success(c, idList)
}

func (s *SysRoleServiceImpl) AssignPermission(c *gin.Context, menu entity.RoleMenu) {
	res := dao.AssignPermission(menu)
	result.Success(c, res)
}

func SysRoleService() ISysRoleService {
	return &sysRoleService
}
