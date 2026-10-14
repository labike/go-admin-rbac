package entity

import "go-admin/common/util"

type SysRole struct {
	Id          uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`
	RoleName    string     `gorm:"column:role_name;varchar(64);comment:'角色名称';NOT NULL" json:"roleName"`
	RoleKey     string     `gorm:"column:role_key;varchar(64);comment:'权限';NOT NULL" json:"roleKey"`
	Status      int        `gorm:"column:status;default:1;comment:'状态:1->启用,2->禁用';NOT NULL" json:"status"`
	Description string     `gorm:"column:description;varchar(500);comment:'描述'" json:"description"`
	CreateTime  util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

type AddSysRole struct {
	RoleName    string
	RoleKey     string
	Status      int
	Description string
}

type UpdateSysRole struct {
	Id          uint
	RoleName    string
	RoleKey     string
	Status      int
	Description string
}

type UpdateSysRoleStatus struct {
	Id     uint
	Status int // 1启用 2禁用
}

type SysRoleVo struct {
	Id       int    `json:"id"`
	RoleName string `json:"roleName"`
}

type SysRoleDto struct {
	Id int `json:"id"`
}

type IdVo struct {
	Id int `json:"id"`
}

type RoleMenu struct {
	Id      uint   `json:"id" binding:"required"`
	MenuIds []uint `json:"menuIds" binding:"required"`
}
