package service

import (
	"go-admin/api/dao"
	"go-admin/api/entity"
	"go-admin/common/result"

	"github.com/gin-gonic/gin"
)

type ISysPostService interface {
	CreateSysPost(c *gin.Context, sysPost entity.SysPost)
	GetSysPostList(c *gin.Context, PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string)
	GetSysPostById(c *gin.Context, Id int)
	UpdateSysPost(c *gin.Context, sysPost entity.SysPost)
	DelSysPostById(c *gin.Context, dto entity.SysPostIdDto)
	BatchDeleteSysPost(c *gin.Context, dto entity.DelSysPostDto)
	UpdateSysPostStatus(c *gin.Context, dto entity.UpdateSysPostStatusDto)
	QuerySysPostList(c *gin.Context)
}

type SysPostServiceImpl struct{}

func (s *SysPostServiceImpl) CreateSysPost(c *gin.Context, sysPost entity.SysPost) {
	res := dao.CreateSysPost(sysPost)
	if !res {
		result.Failed(c, int(result.ApiCode.POSTEXISTS), result.ApiCode.GetMessage(result.ApiCode.POSTEXISTS))
		return
	}
	result.Success(c, true)
}

func (s *SysPostServiceImpl) GetSysPostList(c *gin.Context, PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string) {
	if PageSize < 1 {
		PageSize = 10
	}

	if PageNum < 1 {
		PageNum = 1
	}

	sysPost, count := dao.GetSysPostList(PageNum, PageSize, PostName, PostStatus, BeginTime, EndTime)
	result.Success(c, map[string]interface{}{
		"total":    count,
		"pageSize": PageSize,
		"pageNum":  PageNum,
		"list":     sysPost,
	})
}

func (s *SysPostServiceImpl) GetSysPostById(c *gin.Context, Id int) {
	res := dao.GetSysPostListById(Id)
	result.Success(c, res)
}

func (s *SysPostServiceImpl) UpdateSysPost(c *gin.Context, sysPost entity.SysPost) {
	res := dao.UpdateSysPost(sysPost)
	result.Success(c, res)
}

func (s *SysPostServiceImpl) DelSysPostById(c *gin.Context, dto entity.SysPostIdDto) {
	dao.DeleteSysPost(dto)
	result.Success(c, true)
}

func (s *SysPostServiceImpl) BatchDeleteSysPost(c *gin.Context, dto entity.DelSysPostDto) {
	dao.BatchDeleteSysPost(dto)
	result.Success(c, true)
}

func (s *SysPostServiceImpl) UpdateSysPostStatus(c *gin.Context, dto entity.UpdateSysPostStatusDto) {
	dao.UpdateSysPostStatus(dto)
	result.Success(c, true)
}

func (s *SysPostServiceImpl) QuerySysPostList(c *gin.Context) {
	res := dao.QuerySysPostList()
	result.Success(c, res)
}

var sysPostService = SysPostServiceImpl{}

func SysPostService() ISysPostService {
	return &sysPostService
}
