package dao

import (
	"go-admin/api/entity"
	"go-admin/common/util"
	. "go-admin/pkg/db"
	"time"
)

// 根据菜单名称查询
func GetSysMenuByName(menuName string) (sysMenu entity.SysMenu) {
	Db.Where("menu_name = ?", menuName).First(&sysMenu)
	return sysMenu
}

// 新增菜单
func CreateSysMenu(addMenu entity.SysMenu) bool {
	sysMenuByName := GetSysMenuByName(addMenu.MenuName)
	if sysMenuByName.ID != 0 {
		return false
	}

	// 目录
	if addMenu.MenuType == 1 {
		sysMenu := entity.SysMenu{
			ParentId:   0,
			MenuName:   addMenu.MenuName,
			Icon:       addMenu.Icon,
			Sort:       addMenu.Sort,
			MenuType:   addMenu.MenuType,
			MenuStatus: addMenu.MenuStatus,
			CreateTime: util.HTime{Time: time.Now()},
			Url:        addMenu.Url,
		}
		Db.Create(&sysMenu)
		return true
	} else if addMenu.MenuType == 2 {
		sysMenu := entity.SysMenu{
			ParentId:   addMenu.ParentId,
			MenuName:   addMenu.MenuName,
			Icon:       addMenu.Icon,
			Sort:       addMenu.Sort,
			MenuType:   addMenu.MenuType,
			MenuStatus: addMenu.MenuStatus,
			CreateTime: util.HTime{Time: time.Now()},
			Url:        addMenu.Url,
			Value:      addMenu.Value,
		}
		Db.Create(&sysMenu)
		return true
	} else if addMenu.MenuType == 3 {
		sysMenu := entity.SysMenu{
			ParentId:   addMenu.ParentId,
			MenuName:   addMenu.MenuName,
			Sort:       addMenu.Sort,
			MenuType:   addMenu.MenuType,
			MenuStatus: addMenu.MenuStatus,
			CreateTime: util.HTime{Time: time.Now()},
			Value:      addMenu.Value,
		}
		Db.Create(&sysMenu)
		return true
	}
	return false
}

func QuerySysMenuVoList() (sysenuVo []entity.SysMenuVo) {
	Db.Table("sys_menu").Select("id, menu_name AS label, parent_id").Scan(&sysenuVo)
	return sysenuVo
}

func GetSysMenuById(Id int) (sysMenu entity.SysMenu) {
	Db.First(&sysMenu, Id)
	return sysMenu
}

func UpdateSysMenu(menu entity.SysMenu) (sysMenu entity.SysMenu) {
	Db.First(&sysMenu, menu.ID)
	sysMenu.Icon = menu.Icon
	sysMenu.MenuType = menu.MenuType
	sysMenu.MenuStatus = menu.MenuStatus
	sysMenu.MenuName = menu.MenuName
	sysMenu.ParentId = menu.ParentId
	sysMenu.Url = menu.Url
	sysMenu.Sort = menu.Sort
	sysMenu.Value = menu.Value
	Db.Save(&sysMenu)
	return sysMenu
}

func GetSysRoleMenu(id uint) (sysRoleMenu entity.SysRoleMenu) {
	Db.Where("id = ?", id).First(&sysRoleMenu)
	return sysRoleMenu
}

func DeleteSysMenu(dto entity.SysMenuIdDto) bool {
	sysRoleMenu := GetSysRoleMenu(dto.Id)
	if sysRoleMenu.MenuId > 0 {
		return false
	}
	Db.Where("id = ?", dto.Id).Delete(&entity.SysRoleMenu{})
	Db.Delete(&entity.SysMenu{}, dto.Id)
	return true
}

func GetSysMenuList(MenuName string, MenuStatus string) (sysMenu entity.SysMenu) {
	curDb := Db.Table("sys_menu").Order("sort")
	if MenuName != "" {
		curDb = curDb.Where("menu_name = ?", MenuName)
	}

	if MenuStatus != "" {
		curDb = curDb.Where("menu_status = ?", MenuStatus)
	}

	curDb.Find(&sysMenu)
	return sysMenu
}

func QueryMenuVoList(AdminId, MenuId uint) (menuVo []entity.MenuVo) {
	const status, menuStatus, menuType = 1, 2, 2
	Db.Table("sys_menu sm").
		Select("sm.menu_name, sm.icon, sm.url").
		Joins("LEFT JOIN sys_role_menu srm N sm.id = srm.menu_id").
		Joins("LEFT JOIN sys_role sr.id = srm.role_id").
		Joins("LEFT JOIN sys_admin_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN sys_admin sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type = ?", menuType).
		Where("sm.parent_id = ?", MenuId).
		Where("sa.id = ?", AdminId).
		Order("sm.sort").
		Scan(&menuVo)
	return menuVo
}

// 当前登录用户左侧菜单列表
func QueryLeftMenuList(Id uint) (leftMenuVo []entity.LeftMenuVo) {
	const status, menuStatus, menuType uint = 1, 2, 1
	Db.Table("sys_menu sm").
		Select("sm.id, sm.menu_name, sm.icon, sm.url").
		Joins("LEFT JOIN sys_role_menu srm ON srm.menu_id = sm.id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		Joins("LEFT JOIN sys_admin_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN sys_admin sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type = ?", menuType).
		Where("sa.id = ?", Id).
		Order("sm.sort").
		Scan(&leftMenuVo)
	return leftMenuVo
}

func QueryPermissionList(Id uint) (valueVo []entity.ValueVo) {
	const status, menuStatus, menuType uint = 1, 2, 1
	Db.Table("sys_menu sm").
		Select("sm.value").
		Joins("LEFT JOIN sys_role_menu srm ON srm.role_id = sm.id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		Joins("LEFT JOIN sys_admin_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN sys_admin sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type = ?", menuType).
		Where("sa.id = ?", Id).
		Scan(&valueVo)
	return valueVo
}
