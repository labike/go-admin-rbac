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

type ISysAdminService interface {
	Login(c *gin.Context, dto entity.LoginDto)
}

type SysAdminServiceImpl struct{}

// 登录
func (s SysAdminServiceImpl) Login(c *gin.Context, dto entity.LoginDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissingLoginParams), result.ApiCode.GetMessage(result.ApiCode.MissingLoginParams))
		return
	}
	code := util.RedisStore{}.Get(dto.IdKey, true)
	if len(code) == 0 {
		result.Failed(c, int(result.ApiCode.VerificationCodeHasExpired), result.ApiCode.GetMessage(result.ApiCode.VerificationCodeHasExpired))
		return
	}
	verifyRes := CaptVerify(dto.IdKey, dto.Image)
	if !verifyRes {
		result.Failed(c, int(result.ApiCode.CAPTCHANOTTRUE), result.ApiCode.GetMessage(result.ApiCode.CAPTCHANOTTRUE))
		return
	}
	sysAdmin := dao.SysAdminDetail(dto)
	if sysAdmin.Password != util.EncryptionMd5(dto.Password) {
		result.Failed(c, int(result.ApiCode.PASSWORDNOTRUE), result.ApiCode.GetMessage(result.ApiCode.PASSWORDNOTRUE))
		return
	}
	const status int = 2
	if sysAdmin.Status == status {
		result.Failed(c, int(result.ApiCode.STATUSISENABLE), result.ApiCode.GetMessage(result.ApiCode.STATUSISENABLE))
		return
	}

	tokenString, _ := jwt.GenerateTokenByAdmin(sysAdmin)
	result.Success(c, map[string]interface{}{
		"token":    tokenString,
		"sysAdmin": sysAdmin,
	})
}

var sysAdminService = SysAdminServiceImpl{}

func SysAdminService() ISysAdminService {
	return &sysAdminService
}
