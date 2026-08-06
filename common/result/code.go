package result

type Codes struct {
	SUCCESS                    uint
	FAILED                     uint
	Message                    map[uint]string
	NOAUTH                     uint
	AUTHFORMATERROR            uint
	INVALIDTOKEN               uint
	MissingLoginParams         uint
	VerificationCodeHasExpired uint
	CAPTCHANOTTRUE             uint
	PASSWORDNOTRUE             uint
	STATUSISENABLE             uint
}

var ApiCode = &Codes{
	SUCCESS:                    200,
	FAILED:                     501,
	NOAUTH:                     403,
	AUTHFORMATERROR:            405,
	INVALIDTOKEN:               406,
	MissingLoginParams:         407,
	VerificationCodeHasExpired: 408,
	CAPTCHANOTTRUE:             409,
	PASSWORDNOTRUE:             410,
	STATUSISENABLE:             411,
}

func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:                    "成功",
		ApiCode.FAILED:                     "失败",
		ApiCode.NOAUTH:                     "token不能为空",
		ApiCode.AUTHFORMATERROR:            "请求头错误",
		ApiCode.MissingLoginParams:         "缺少登录参数",
		ApiCode.VerificationCodeHasExpired: "token过期",
		ApiCode.INVALIDTOKEN:               "token错误",
		ApiCode.CAPTCHANOTTRUE:             "验证码错误",
		ApiCode.PASSWORDNOTRUE:             "密码错误",
		ApiCode.STATUSISENABLE:             "该用户已被禁用",
	}
}

func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
