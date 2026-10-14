package dao

import (
	"go-admin/api/entity"
	"go-admin/common/util"
	. "go-admin/pkg/db"
	"time"
)

func CreateSysRole(dto entity.AddSysRole) bool {
	sysRole := GetSysRoleByName(dto.RoleName)
	if sysRole.Id > 0 {
		return false
	}
	sysRoleKey := GetSysRoleByKey(dto.RoleKey)
	if sysRoleKey.Id > 0 {
		return false
	}
	addSysRole := &entity.SysRole{
		RoleName:    dto.RoleName,
		RoleKey:     dto.RoleKey,
		Status:      dto.Status,
		Description: dto.Description,
		CreateTime:  util.HTime{Time: time.Now()},
	}
	tx := Db.Create(&addSysRole)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

func GetSysRoleByName(name string) (sysRole entity.SysRole) {
	Db.Where("role_name = ?", name).First(&sysRole)
	return sysRole
}

func GetSysRoleByKey(key string) (sysRole entity.SysRole) {
	Db.Where("role_key = ?", key).First(&sysRole)
	return sysRole
}

func GetSysRoleById(Id int) (sysRole entity.SysRole) {
	Db.Where("id = ?", Id).First(&sysRole)
	return sysRole
}

func UpdateSysRole(dto entity.UpdateSysRole) (sysRole entity.SysRole) {
	Db.First(&sysRole, dto.Id)
	sysRole.RoleName = dto.RoleName
	sysRole.RoleKey = dto.RoleKey
	sysRole.Status = dto.Status
	if dto.Description != "" {
		sysRole.Description = dto.Description
	}
	Db.Save(&sysRole)
	return sysRole
}

func UpdateSysRoleStatus(dto entity.UpdateSysRoleStatus) bool {
	var sysRole entity.SysRole
	Db.First(&sysRole, dto.Id)
	sysRole.Status = dto.Status
	tx := Db.Save(&sysRole)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

func DeleteSysRoleById(dto entity.SysRoleDto) {
	Db.Table("sys_role").Delete(&entity.SysRole{}, dto.Id)
	Db.Table("sys_role_menu").Where("role_id = ?", dto.Id).Delete(&entity.SysRoleMenu{})
}

func GetSysRoleList(PageNum, PageSize int, RoleName, Status, BeginTime, EndTime string) (sysRole []*entity.SysRole, count int64) {
	curDb := Db.Table("sys_role")
	if RoleName != "" {
		curDb = curDb.Where("role_name = ?", RoleName)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	if Status != "" {
		curDb = curDb.Where("status = ?", Status)
	}
	curDb.Count(&count)
	curDb.Limit(PageNum).Offset(PageSize * (PageNum - 1)).Order("create_time DESC").Find(&sysRole)
	return sysRole, count
}

func GetSysRoleVoList() (sysRole []entity.SysRoleVo) {
	Db.Table("sys_role").Select("id, role_name").Scan(&sysRole)
	return sysRole
}

// 根据角色id查询菜单权限 select sm.id from sys_menu sm left join sys_role_menu srm on srm.menu_id = sm.id
// left join sys_role sr on sr.id = srm.role_id where sm.menu_type = 3 AND sr.id = Id
func QueryRoleMenuIdList(Id int) (idVo []entity.IdVo) {
	const menuType int = 3 // 菜单
	Db.Table("sys_menu sm").
		Select("sm.id").
		Joins("LEFT JOIN sys_role_menu srm ON srm.menu_id = sm.id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		Where("sm.menu_type = ?", menuType).
		Where("sr.id = ?", Id).
		Scan(&idVo)

	return idVo
}

func AssignPermission(menu entity.RoleMenu) (err error) {
	err = Db.Table("sys_role_menu").Where("role_id = ?", menu.Id).Delete(&entity.SysRoleMenu{}).Error
	if err != nil {
		return err
	}
	for _, v := range menu.MenuIds {
		var entityMenu entity.SysRoleMenu
		entityMenu.RoleId = menu.Id
		entityMenu.MenuId = v
		Db.Create(&entityMenu)
	}
	return err
}
