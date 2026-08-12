package dao

import (
	"go-admin/api/entity"
	"go-admin/common/util"
	. "go-admin/pkg/db"
	"time"
)

func GetSysDeptList(DeptName, DeptStatus string) (dept []entity.SysDept) {
	curDb := Db.Table("sys_dept")
	if DeptName != "" {
		curDb = curDb.Where("dept_name = ?", DeptName)
	}

	if DeptStatus != "" {
		curDb = curDb.Where("dept_status = ?", DeptStatus)
	}

	curDb.Find(&dept)
	return dept
}

func GetSysDeptByName(DeptName string) (dept entity.SysDept) {
	Db.Where("dept_name = ?", DeptName).First(&dept)
	return dept
}

func CreateSysDept(sysDept entity.SysDept) bool {
	deptName := GetSysDeptByName(sysDept.DeptName)
	if deptName.ID > 0 {
		return false
	}

	if sysDept.DeptType == 1 {
		dto := entity.SysDept{
			ParentId:   0,
			DeptStatus: sysDept.DeptStatus,
			DeptName:   sysDept.DeptName,
			DeptType:   sysDept.DeptType,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&dto)
		return true
	} else {
		dto := entity.SysDept{
			ParentId:   sysDept.ParentId,
			DeptStatus: sysDept.DeptStatus,
			DeptName:   sysDept.DeptName,
			DeptType:   sysDept.DeptType,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Save(&dto)
		return true
	}
}

func GetSysDeptById(Id int) (sysDept entity.SysDept) {
	Db.First(&sysDept, Id)
	return sysDept
}

func UpdateSysDept(dept entity.SysDept) (sysDept entity.SysDept) {
	Db.First(&sysDept, dept.ID)
	sysDept.DeptName = dept.DeptName
	sysDept.DeptType = dept.DeptType
	sysDept.DeptStatus = dept.DeptStatus
	sysDept.ParentId = dept.ParentId
	Db.Save(&sysDept)
	return sysDept
}

// 删除部门前查询是否友用户
func GetSysAdminInDept(id int) (sysAdmin entity.SysAdmin) {
	Db.Where("dept_id = ?", id).First(&sysAdmin)
	return sysAdmin
}

func DeleteSysDeptById(dto entity.SysDeptIdDto) bool {
	sysAdmin := GetSysAdminInDept(dto.Id)
	if sysAdmin.ID > 0 {
		return false
	}
	Db.Where("parent_id = ?", dto.Id).Delete(&entity.SysDept{})
	Db.Delete(&entity.SysDept{}, dto.Id)
	return true
}

func QueryDeptList() (sysDeptVo []entity.SysDeptVo) {
	Db.Table("sys_dept").Select("id, dept_name AS label, parent_id").Scan(&sysDeptVo)
	return sysDeptVo
}
