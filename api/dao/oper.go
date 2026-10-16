package dao

import (
	"go-admin/api/entity"
	. "go-admin/pkg/db"
)

func CreateSysOperLog(log entity.SysOperate) {
	Db.Create(&log)
}

func GetSysOperList(Username, BeginTime, EndTime string, PageSize, PageNum int) (sysOperate []entity.SysOperate, count int64) {
	curDb := Db.Table("sys_operate")
	if Username != "" {
		curDb = curDb.Where("username = ?", Username)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("begin_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time desc").Find(&sysOperate)
	return sysOperate, count
}

func BatchDeleteSysOperLog(dto entity.SysOperateIdsDto) {
	Db.Where("id in (?)", dto.Ids).Delete(entity.SysOperate{})
}

func DeleteSysOperLogById(dto entity.SysOperateLogDto) {
	Db.Where(entity.SysOperate{}, dto)
}
