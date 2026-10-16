package service

import (
	"go-admin/api/dao"
	"go-admin/api/entity"
	"go-admin/common/result"
	"go-admin/common/util"
	"go-admin/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 接口只声明方法签名，不关心具体实现
type ISysAdminService interface {
	Login(c *gin.Context, dto entity.LoginDto)
	CreateSysAdmin(c *gin.Context, dto entity.AddSysAdminDto)
	GetSysAdminInfo(c *gin.Context, Id int)
	UpdateSysAdmin(c *gin.Context, dto entity.UpdateSysAdminDto)
	DeleteSysAdmin(c *gin.Context, dto entity.SysAdminIdDto)
	UpdateSysAdminStatus(c *gin.Context, dto entity.UpdateSysAdminStatusDto)
	UpdateSysAdminPwd(c *gin.Context, dto entity.ResetSysAdminPasswordDto)
	GetSysAdminList(c *gin.Context, PageSize, PageNum int, Username, Status, BeginTime, EndTime string)
	UpdateUser(c *gin.Context, dto entity.UpdateUserDto)
	UpdatePersonPwd(c *gin.Context, dto entity.UpdatePersonPwdDto)
	//QueryMenuVoList(c *gin.Context, MenuId int)
	//QueryLeftMenuList(c *gin.Context, Id uint)
}

// 怎么做
type SysAdminServiceImpl struct{}

// 登录
func (s SysAdminServiceImpl) Login(c *gin.Context, dto entity.LoginDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissingLoginParams), result.ApiCode.GetMessage(result.ApiCode.MissingLoginParams))
		return
	}
	ip := c.ClientIP()
	code := util.RedisStore{}.Get(dto.IdKey, true)
	if len(code) == 0 {
		dao.CreateLoginLog(dto.Username, ip, util.GetRealAddressByIP(ip), util.GetBrowser(c), util.GetOs(c), "验证码已过期", 2)
		result.Failed(c, int(result.ApiCode.VerificationCodeHasExpired), result.ApiCode.GetMessage(result.ApiCode.VerificationCodeHasExpired))
		return
	}
	verifyRes := CaptVerify(dto.IdKey, dto.Image)
	if !verifyRes {
		dao.CreateLoginLog(dto.Username, ip, util.GetRealAddressByIP(ip), util.GetBrowser(c), util.GetOs(c), "验证码错误", 2)
		result.Failed(c, int(result.ApiCode.CAPTCHANOTTRUE), result.ApiCode.GetMessage(result.ApiCode.CAPTCHANOTTRUE))
		return
	}
	sysAdmin := dao.SysAdminDetail(dto)
	if sysAdmin.Password != util.EncryptionMd5(dto.Password) {
		dao.CreateLoginLog(dto.Username, ip, util.GetRealAddressByIP(ip), util.GetBrowser(c), util.GetOs(c), "账号或密码错误", 2)
		result.Failed(c, int(result.ApiCode.PASSWORDNOTRUE), result.ApiCode.GetMessage(result.ApiCode.PASSWORDNOTRUE))
		return
	}
	const status int = 2
	if sysAdmin.Status == status {
		dao.CreateLoginLog(dto.Username, ip, util.GetRealAddressByIP(ip), util.GetBrowser(c), util.GetOs(c), "账号已停用", 2)
		result.Failed(c, int(result.ApiCode.STATUSISENABLE), result.ApiCode.GetMessage(result.ApiCode.STATUSISENABLE))
		return
	}

	tokenString, _ := jwt.GenerateTokenByAdmin(sysAdmin)
	dao.CreateLoginLog(dto.Username, ip, util.GetRealAddressByIP(ip), util.GetBrowser(c), util.GetOs(c), "登录成功", 1)
	// 生成左侧菜单列表
	var leftMenuVo []entity.LeftMenuVo
	leftMenuList := dao.QueryLeftMenuList(sysAdmin.ID)
	for _, v := range leftMenuList {
		menuVoList := dao.QueryMenuVoList(sysAdmin.ID, v.Id)
		item := entity.LeftMenuVo{}
		item.MenuVoList = menuVoList
		item.Id = v.Id
		item.MenuName = v.MenuName
		item.Icon = v.Icon
		item.Url = v.Url
		leftMenuVo = append(leftMenuVo, item)
	}

	// 权限列表
	permissionsList := dao.QueryPermissionList(sysAdmin.ID)
	var stringList = make([]string, 0)
	for _, v := range permissionsList {
		stringList = append(stringList, v.Value)
	}
	result.Success(c, map[string]interface{}{
		"token":        tokenString,
		"sysAdmin":     sysAdmin,
		"leftMenuList": leftMenuVo,
		"permissions":  permissionsList,
	})
}

func (s SysAdminServiceImpl) CreateSysAdmin(c *gin.Context, dto entity.AddSysAdminDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MISSINGADMINPARAMS), result.ApiCode.GetMessage(result.ApiCode.MISSINGADMINPARAMS))
		return
	}
	res := dao.CreateSysAdmin(dto)
	if !res {
		result.Failed(c, int(result.ApiCode.USERNAMEEXISTS), result.ApiCode.GetMessage(result.ApiCode.USERNAMEEXISTS))
		return
	}
	result.Success(c, res)
	return
}

func (s SysAdminServiceImpl) GetSysAdminInfo(c *gin.Context, Id int) {
	res := dao.GetSysAdminInfo(Id)
	result.Success(c, res)
}

func (s SysAdminServiceImpl) UpdateSysAdmin(c *gin.Context, dto entity.UpdateSysAdminDto) {
	//err := validator.New().Struct(dto)
	//if err != nil {
	//	result.Failed(c, int(result.ApiCode.MISSINGCHANGEUSERPARAMS), result.ApiCode.GetMessage(result.ApiCode.MISSINGCHANGEUSERPARAMS))
	//	return
	//}
	//id, _ := jwt.GetAdminId(c)
	//dto.Id = id
	res := dao.UpdateSysAdmin(dto)
	result.Success(c, res)
}

func (s SysAdminServiceImpl) DeleteSysAdmin(c *gin.Context, dto entity.SysAdminIdDto) {
	dao.DeleteSysAdmin(dto)
	result.Success(c, true)
}

func (s SysAdminServiceImpl) UpdateSysAdminStatus(c *gin.Context, dto entity.UpdateSysAdminStatusDto) {
	dao.UpdateSysAdminStatus(dto)
	result.Success(c, true)
}

func (s SysAdminServiceImpl) UpdateSysAdminPwd(c *gin.Context, dto entity.ResetSysAdminPasswordDto) {
	dao.UpdateSysAdminPwd(dto)
	result.Success(c, true)
}

func (s SysAdminServiceImpl) GetSysAdminList(c *gin.Context, PageSize, PageNum int, Username, Status, BeginTime, EndTime string) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysAdmin, count := dao.GetSysAdminList(PageSize, PageNum, Username, Status, BeginTime, EndTime)
	result.Success(c, map[string]interface{}{
		"list":     sysAdmin,
		"total":    count,
		"pageSize": PageSize,
		"pageNum":  PageNum,
	})
	return
}

func (s SysAdminServiceImpl) UpdateUser(c *gin.Context, dto entity.UpdateUserDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.UPDATEUSERFAILED), result.ApiCode.GetMessage(result.ApiCode.UPDATEUSERFAILED))
		return
	}
	id, _ := jwt.GetAdminId(c)
	dto.Id = id
	result.Success(c, dao.UpdateUser(dto))
}

func (s SysAdminServiceImpl) UpdatePersonPwd(c *gin.Context, dto entity.UpdatePersonPwdDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.UPDATEPWDFAILED), result.ApiCode.GetMessage(result.ApiCode.UPDATEPWDFAILED))
		return
	}
	sysAdmin, _ := jwt.GetAdmin(c)
	dto.Id = sysAdmin.ID
	sysAdminExist := dao.GetSysAdminByUsername(sysAdmin.Username)
	if sysAdminExist.Password != util.EncryptionMd5(dto.Password) {
		result.Failed(c, int(result.ApiCode.UPDATEPWDFAILED), result.ApiCode.GetMessage(result.ApiCode.UPDATEPWDFAILED))
		return
	}
	if dto.NewPassword != dto.ResetPassword {
		result.Failed(c, int(result.ApiCode.UPDATEPWDFAILED), result.ApiCode.GetMessage(result.ApiCode.UPDATEPWDFAILED))
		return
	}
	dto.NewPassword = util.EncryptionMd5(dto.NewPassword)
	sysAdminUpdatePwd := dao.UpdatePersonPwd(dto)
	tokenString, _ := jwt.GenerateTokenByAdmin(sysAdminUpdatePwd)
	result.Success(c, map[string]interface{}{
		"token":    tokenString,
		"sysAdmin": sysAdminUpdatePwd,
	})
	return
}

// 单例对象，接口的实现实例
var sysAdminService = SysAdminServiceImpl{}

// 单例的获取入口
func SysAdminService() ISysAdminService {
	return &sysAdminService
}
