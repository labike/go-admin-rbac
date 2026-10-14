package entity

// 用户与角色关联

type SysAdminRole struct {
	AdminId uint `gorm:"column:admin_id;comment:'用户id';NOT NULL" json:"admin_id"`
	RoleId  uint `gorm:"column:role_id;comment:'角色id';NOT NULL" json:"role_id"`
}

func (SysAdminRole) TableName() string {
	return "sys_admin_role"
}
