package dao

import (
	"go-admin/api/entity"
	"go-admin/common/util"
	. "go-admin/pkg/db"
	"time"
)

// 用户数据层, 只管数据存取，不含业务判断

func SysAdminDetail(dto entity.LoginDto) (sysAdmin entity.SysAdmin) {
	username := dto.Username
	Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

func GetSysAdminByUsername(username string) (sysAdmin entity.SysAdmin) {
	Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

func CreateSysAdmin(dto entity.AddSysAdminDto) bool {
	sysAdminByUsername := GetSysAdminByUsername(dto.Username)
	if sysAdminByUsername.ID > 0 {
		return false
	}

	sysAdmin := entity.SysAdmin{
		Username:   dto.Username,
		Password:   util.EncryptionMd5(dto.Password),
		Nickname:   dto.Nickname,
		Phone:      dto.Phone,
		Email:      dto.Email,
		Status:     dto.Status,
		PostId:     dto.PostId,
		DeptId:     dto.DeptId,
		Note:       dto.Note,
		CreateTime: util.HTime{Time: time.Now()},
	}
	tx := Db.Create(&sysAdmin)
	sysAdminExists := GetSysAdminByUsername(dto.Username)
	// 用户和角色绑定
	var entityAdmin entity.SysAdminRole
	entityAdmin.AdminId = sysAdminExists.ID
	entityAdmin.RoleId = dto.RoleId
	Db.Create(&entityAdmin)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

func GetSysAdminInfo(Id int) (sysAdminInfo entity.SysAdminInfo) {
	Db.Table("sys_admin").
		Select("sys_admin.*, sys_admin_role.role_id").
		Joins("LEFT JOIN sys_admin_role ON sys_admin.id = sys_admin_role.admin_id").
		Joins("LEFT JOIN sys_role ON sys_admin_role.role_id = sys_admin.id").
		First(&sysAdminInfo, Id)
	return sysAdminInfo
}

func UpdateSysAdmin(dto entity.UpdateSysAdminDto) (sysAdmin entity.SysAdmin) {
	Db.First(&sysAdmin, dto.Id)
	if dto.Username != "" {
		sysAdmin.Username = dto.Username
	}
	sysAdmin.PostId = dto.PostId
	sysAdmin.Status = dto.Status
	sysAdmin.DeptId = dto.DeptId
	if dto.Nickname != "" {
		sysAdmin.Nickname = dto.Nickname
	}
	if dto.Phone != "" {
		sysAdmin.Phone = dto.Phone
	}
	if dto.Email != "" {
		sysAdmin.Email = dto.Email
	}
	if dto.Note != "" {
		sysAdmin.Note = dto.Note
	}
	Db.Save(&sysAdmin)
	// 删除之前的角色分配新的角色
	var sysAdminRole entity.SysAdminRole
	Db.Where("admin_id = ?", dto.Id).Delete(&entity.SysAdminRole{})
	sysAdminRole.RoleId = dto.RoleId
	sysAdminRole.AdminId = dto.Id
	Db.Create(&sysAdminRole)
	return sysAdmin
}

func DeleteSysAdmin(dto entity.SysAdminIdDto) {
	Db.First(&entity.SysAdmin{}, dto.Id)
	Db.Delete(&entity.SysAdminRole{}, dto.Id)
	Db.Where("admin_id = ?", dto.Id).Delete(&entity.SysAdminRole{})
}

func UpdateSysAdminStatus(dto entity.UpdateSysAdminStatusDto) {
	var sysAdmin entity.SysAdmin
	Db.First(&sysAdmin, dto.Id)
	sysAdmin.Status = dto.Status
	Db.Save(&sysAdmin)
}

func UpdateSysAdminPwd(dto entity.ResetSysAdminPasswordDto) {
	var sysAdmin entity.SysAdmin
	Db.First(&sysAdmin, dto.Id)
	sysAdmin.Password = util.EncryptionMd5(dto.Password)
	Db.Save(&sysAdmin)
}

// 分页查询用户
func GetSysAdminList(PageSize, PageNum int, Username, Status, BeginTime, EndTime string) (sysAdminVo []entity.SysAdminVo, count int64) {
	curDb := Db.Table("sys_admin").
		Select("sys_admin.*, sys_post.post_name, sys_role.role_name, sys_dept.dept_name").
		Joins("LEFT JOIN sys_post ON sys_admin.post_id = sys_post.id").
		Joins("LEFT JOIN sys_admin_role ON sys_admin.id = sys_admin_role.admin_id").
		Joins("LEFT JOIN sys_role ON sys_role.id = sys_admin_role.role_id").
		Joins("LEFT JOIN sys_dept ON sys_dept.id = sys_admin.dept_id")

	if Username != "" {
		curDb = curDb.Where("sys_admin.username = ?", Username)
	}

	if Status != "" {
		curDb = curDb.Where("sys_admin.status = ?", Status)
	}

	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("sys_admin.create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}

	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("sys_admin.create_time DESC").Find(&sysAdminVo)
	return sysAdminVo, count
}

func UpdateUser(dto entity.UpdateUserDto) (sysAdmin entity.SysAdmin) {
	Db.First(&sysAdmin, dto.Id)
	if dto.Icon != "" {
		sysAdmin.Icon = dto.Icon
	}
	if dto.Username != "" {
		sysAdmin.Username = dto.Username
	}
	if dto.Phone != "" {
		sysAdmin.Phone = dto.Phone
	}
	if dto.Email != "" {
		sysAdmin.Email = dto.Email
	}
	if dto.Nickname != "" {
		sysAdmin.Nickname = dto.Nickname
	}
	Db.Save(&sysAdmin)
	return sysAdmin
}

// 修改个人密码
func UpdatePersonPwd(dto entity.UpdatePersonPwdDto) (sysAdmin entity.SysAdmin) {
	Db.First(&sysAdmin, dto.Id)
	sysAdmin.Password = dto.NewPassword
	Db.Save(&sysAdmin)
	return sysAdmin
}
