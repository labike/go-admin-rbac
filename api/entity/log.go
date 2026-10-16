package entity

import "go-admin/common/util"

type SysLoginInfo struct {
	ID            uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`
	Username      string     `gorm:"column:username;comment:'用户名'" json:"username"`
	IpAddress     string     `gorm:"column:ip_address;comment:'ip地址'" json:"ip_address"`
	LoginLocation string     `gorm:"column:login_location;comment:'登录地点'" json:"login_location"`
	Browser       string     `gorm:"column:browser;varchar(50);comment:'浏览器类型'" json:"browser"`
	Os            string     `gorm:"column:os;varchar(50);comment:'操作系统'" json:"os"`
	LoginStatus   int        `gorm:"column:login_status;comment:'登录状态:1->成功,2->失败'" json:"loginStatus"`
	Message       string     `gorm:"column:message;comment:'提示消息'" json:"message"`
	LoginTime     util.HTime `gorm:"column:login_time;comment:'访问时间'" json:"loginTime"`
}

func (SysLoginInfo) TableName() string {
	return "sys_login_info"
}

type SysLoginIdDto struct {
	Id uint `json:"id"`
}

type SysLoginDto struct {
	Ids []uint
}
