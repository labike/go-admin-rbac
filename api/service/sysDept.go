package service

import (
	"go-admin/api/dao"
	"go-admin/api/entity"
	"go-admin/common/result"

	"github.com/gin-gonic/gin"
)

type ISysDeptService interface {
	GetSysDeptList(c *gin.Context, DeptName, DeptStatus string)
	CreateSysDept(c *gin.Context, sysDept entity.SysDept)
	GetSysDeptById(c *gin.Context, Id int)
	UpdateSysDept(c *gin.Context, sysDept entity.SysDept)
	DeleteSysDeptById(c *gin.Context, dto entity.SysDeptIdDto)
	QuerySysDeptList(c *gin.Context)
}

type ISysDeptServiceImpl struct{}

var sysDeptService = ISysDeptServiceImpl{}

func (d *ISysDeptServiceImpl) GetSysDeptList(c *gin.Context, DeptName, DeptStatus string) {
	res := dao.GetSysDeptList(DeptName, DeptStatus)
	result.Success(c, res)
}

func (d *ISysDeptServiceImpl) CreateSysDept(c *gin.Context, sysDept entity.SysDept) {
	bool := dao.CreateSysDept(sysDept)
	if !bool {
		result.Failed(c, int(result.ApiCode.DEPTEXISTS), result.ApiCode.GetMessage(result.ApiCode.DEPTEXISTS))
		return
	}
	result.Success(c, true)
}

func (d *ISysDeptServiceImpl) GetSysDeptById(c *gin.Context, Id int) {
	res := dao.GetSysDeptById(Id)
	result.Success(c, res)
}

func (d *ISysDeptServiceImpl) UpdateSysDept(c *gin.Context, sysDept entity.SysDept) {
	res := dao.UpdateSysDept(sysDept)
	result.Success(c, res)
}

func (d *ISysDeptServiceImpl) DeleteSysDeptById(c *gin.Context, dto entity.SysDeptIdDto) {
	bool := dao.DeleteSysDeptById(dto)
	if !bool {
		result.Failed(c, int(result.ApiCode.DEPTDISABLED), result.ApiCode.GetMessage(result.ApiCode.DEPTDISABLED))
		return
	}
	result.Success(c, true)
}

func (d *ISysDeptServiceImpl) QuerySysDeptList(c *gin.Context) {
	res := dao.QueryDeptList()
	result.Success(c, res)
}

func SysDeptService() ISysDeptService {
	return &sysDeptService
}
