package entity

import "go-admin/common/util"

type SysAdmin struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`
	PostId     int        `gorm:"column:post_id;comment:'岗位id'" json:"postId"`
	DeptId     int        `gorm:"column:dept_id;comment:'部门id'" json:"deptId"`
	Username   string     `gorm:"column:username;varchar(64);comment:'用户名';NOT NULL" json:"username"`
	Password   string     `gorm:"column:password;varchar(64);comment:'密码';NOT NULL" json:"password"`
	Nickname   string     `gorm:"column:nickname;varchar(64);comment:'昵称'" json:"nickname"`
	Status     int        `gorm:"column:status;default:1;comment:'状态:1->启用,2->禁用';NOT NULL" json:"status"`
	Icon       string     `gorm:"column:icon;varchar(500);comment:'头像'" json:"icon"`
	Email      string     `gorm:"column:email;varchar(64);comment:'邮箱'" json:"email"`
	Phone      string     `gorm:"column:phone;varchar(64);comment:'手机号'" json:"phone"`
	Note       string     `gorm:"column:note;varchar(500);comment:'备注'" json:"note"`
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`
}

func (SysAdmin) TableName() string {
	return "sys_admin"
}

// 鉴权用户结构体, 作为jwt的payload
type JwtAdmin struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Note     string `json:"note"`
}

// 登录请求体
type LoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Image    string `json:"image" validate:"required,min=4,max=6"` // 验证码
	IdKey    string `json:"idKey" validate:"required"`
}

type AddSysAdminDto struct {
	PostId   int    `validate:"required"`
	RoleId   uint   `validate:"required"`
	DeptId   int    `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required"`
	Nickname string `validate:"required"`
	Phone    string `validate:"required"`
	Email    string `validate:"required"`
	Note     string
	Status   int `validate:"required"`
}

type SysAdminInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Status   int    `json:"status"`
	PostId   int    `json:"postId"`
	DeptId   int    `json:"deptId"`
	RoleId   int    `json:"roleId"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Note     string `json:"note"`
}

type UpdateSysAdminDto struct {
	Id       uint
	Username string
	Nickname string
	Status   int
	PostId   int
	DeptId   int
	RoleId   uint
	Email    string
	Phone    string
	Note     string
}

type SysAdminIdDto struct {
	Id uint `json:"id"`
}

type UpdateSysAdminStatusDto struct {
	Id     uint
	Status int
}

type ResetSysAdminPasswordDto struct {
	Id       uint
	Password string
}
type UpdateUserDto struct {
	Id       uint
	Icon     string `validate:"required"`
	Username string `validate:"required"`
	Nickname string `validate:"required"`
	Phone    string `validate:"required"`
	Email    string `validate:"required"`
	Note     string `validate:"required"`
}

type UpdateUserPwdDto struct {
	Id       uint
	Password string
}

type UpdatePersonPwdDto struct {
	Id            uint
	Password      string `validate:"required"`
	NewPassword   string `validate:"required"`
	ResetPassword string `validate:"required"`
}

type SysAdminVo struct {
	ID         uint       `json:"id"`
	Username   string     `json:"username"`
	Nickname   string     `json:"nickname"`
	Status     int        `json:"status"`
	PostId     int        `json:"postId"`
	PostName   string     `json:"postName"`
	DeptId     int        `json:"deptId"`
	DeptName   string     `json:"deptName"`
	RoleId     int        `json:"roleId"`
	RoleName   string     `json:"roleName"`
	Email      string     `json:"email"`
	Phone      string     `json:"phone"`
	Note       string     `json:"note"`
	CreateTime util.HTime `json:"createTime"`
}
