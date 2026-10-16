package entity

import (
	"go-admin/common/util"
)

type SysOperate struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`
	AdminId    int        `gorm:"column:post_id;comment:'管理员id';NOT NULL" json:"adminId"`
	Username   string     `gorm:"column:username;varchar(64);comment:'管理员账号';NOT NULL" json:"username"`
	Method     string     `gorm:"column:method;varchar(64);comment:'请求方式';NOT NULL" json:"method"`
	Ip         string     `gorm:"column:ip;varchar(64);comment:'ip'" json:"ip"`
	Url        string     `gorm:"column:url;varchar(500);comment:'url'" json:"url"`
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`
}

func (SysOperate) TableName() string {
	return "sys_operate"
}

type SysOperateLogDto struct {
	Id uint `json:"id"`
}

type SysOperateIdsDto struct {
	Ids []uint
}
