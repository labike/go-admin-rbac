package dao

import (
	"go-admin/api/entity"
	. "go-admin/pkg/db"
)

// 用户数据层, 只管数据存取，不含业务判断

func SysAdminDetail(dto entity.LoginDto) (sysAdmin entity.SysAdmin) {
	username := dto.Username
	Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}
