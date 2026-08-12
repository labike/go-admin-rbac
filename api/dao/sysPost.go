package dao

import (
	"go-admin/api/entity"
	"go-admin/common/util"
	. "go-admin/pkg/db"
	"time"
)

// 根据code查询
func GetSysPostByCode(code string) (sysPost entity.SysPost) {
	Db.Where("post_code = ?", code).First(&sysPost)
	return sysPost
}

// 根据名称查询
func GetSysPostByName(name string) (sysPost entity.SysPost) {
	Db.Where("post_name = ?", name).First(&sysPost)
	return sysPost
}

// 新增岗位
func CreateSysPost(sysPost entity.SysPost) bool {
	sysPostByCode := GetSysPostByCode(sysPost.PostCode)
	if sysPostByCode.ID > 0 {
		return false
	}

	sysPostByName := GetSysPostByName(sysPost.PostName)
	if sysPostByName.ID > 0 {
		return false
	}

	addSysPost := entity.SysPost{
		PostCode:   sysPost.PostCode,
		PostName:   sysPost.PostName,
		PostStatus: sysPost.PostStatus,
		CreateTime: util.HTime{
			Time: time.Now(),
		},
		Remark: sysPost.Remark,
	}
	tx := Db.Save(&addSysPost)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// 分页查询
func GetSysPostList(PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string) (sysPost []entity.SysPost, count int64) {
	curDb := Db.Table("sys_post")
	if PostName != "" {
		curDb = curDb.Where("post_name = ?", PostName)
	}

	if PostStatus != "" {
		curDb = curDb.Where("post_status = ?", PostStatus)
	}

	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("`create_time` BETWEEN ? AND ?", BeginTime, EndTime)
	}

	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time desc").Find(&sysPost)
	return sysPost, count
}

// 根据岗位id查询
func GetSysPostListById(Id int) (sysPost entity.SysPost) {
	Db.Where("id = ?", Id).First(&sysPost)
	return sysPost
}

// 编辑岗位
func UpdateSysPost(post entity.SysPost) (sysPost entity.SysPost) {
	Db.First(&sysPost, post.ID)
	sysPost.PostName = post.PostName
	sysPost.PostStatus = post.PostStatus
	sysPost.PostCode = post.PostCode
	if post.Remark != "" {
		sysPost.Remark = post.Remark
	}
	Db.Save(&sysPost)
	return sysPost
}

// 根据单个id删除
func DeleteSysPost(dto entity.SysPostIdDto) {
	Db.Delete(&entity.SysPost{}, dto.Id)
}

// 批量删除
func BatchDeleteSysPost(dto entity.DelSysPostDto) {
	Db.Where("id in (?)", dto.Ids).Delete(&entity.SysPost{})
}

// 修改状态
func UpdateSysPostStatus(dto entity.UpdateSysPostStatusDto) {
	var sysPost entity.SysPost
	Db.First(&sysPost, dto.Id)
	sysPost.PostStatus = dto.PostStatus
	Db.Save(&sysPost)
}

// 下拉列表
func QuerySysPostList() (vo []entity.SysPostVo) {
	Db.Table("sys_post").Select("id, post_name").Scan(&vo)
	return vo
}
