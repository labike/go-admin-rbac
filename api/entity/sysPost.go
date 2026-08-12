package entity

import "go-admin/common/util"

type SysPost struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`
	PostCode   string     `gorm:"column:post_code;varchar(64);comment:'岗位编码';NOT NULL" json:"postCode"`
	PostName   string     `gorm:"column:post_name;varchar(50);comment:'岗位名称';NOT NULL" json:"postName"`
	PostStatus int        `gorm:"column:post_status;default:1;comment:'岗位状态(1正常;2停用)';NOT NULL" json:"postStatus"`
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';primaryKey;NOT NULL" json:"createTime"`
	Remark     string     `gorm:"column:remark;varchar(500);comment:'备注'" json:"remark"`
}

func (SysPost) TableName() string {
	return "sys_post"
}

type SysPostIdDto struct {
	Id int `json:"id"`
}

type DelSysPostDto struct {
	Ids []int
}

// 修改状态
type UpdateSysPostStatusDto struct {
	Id         int
	PostStatus int
}

// 岗位下拉列表
type SysPostVo struct {
	Id       int    `json:"id"`
	PostName string `json:"postName"`
}
