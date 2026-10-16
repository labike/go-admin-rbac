package dao

import (
	"go-admin/api/entity"
	"go-admin/common/util"
	. "go-admin/pkg/db"
	"time"
)

func CreateLoginLog(username, ipAddress, loginLocation, browser, os, message string, loginStatus int) {
	sysLoginInfo := entity.SysLoginInfo{
		Username:      username,
		IpAddress:     ipAddress,
		LoginLocation: loginLocation,
		Browser:       browser,
		Os:            os,
		Message:       message,
		LoginStatus:   loginStatus,
		LoginTime:     util.HTime{Time: time.Now()},
	}
	Db.Save(&sysLoginInfo)
}

func GetLoginLog(Username, LoginStatus, BeginTime, EndTime string, PageSize, PageNum int) (sysLoginInfo []entity.SysLoginInfo, count int64) {
	curDb := Db.Table("sys_login_info")
	if Username != "" {
		curDb = curDb.Where("username = ?", Username)
	}
	if LoginStatus != "" {
		curDb = curDb.Where("login_status = ?", LoginStatus)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("login_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("login_time desc").Find(&sysLoginInfo)
	return sysLoginInfo, count
}

func BatchDeleteLoginLog(dto entity.SysLoginDto) {
	Db.Where("id in (?)", dto.Ids).Delete(&entity.SysLoginInfo{})
}

func DeleteLoginById(dto entity.SysLoginIdDto) {
	Db.Delete(&entity.SysLoginInfo{}, dto.Id)
}

func ClearLoginLog() {
	Db.Exec("truncate table sys_login_info")
}
