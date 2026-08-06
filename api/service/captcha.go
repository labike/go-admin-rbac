package service

import (
	"go-admin/common/util"
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// 验证码生成
var store = util.RedisStore{}

func CaptchaMake() (id, b64s string) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString

	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           240,
		NoiseCount:      1,
		ShowLineOptions: 2 | 4,
		Length:          6,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	lid, lb64s, _, _ := captcha.Generate()
	return lid, lb64s
}

func CaptVerify(id, capt string) bool {
	if store.Verify(id, capt, false) == true {
		return true
	} else {
		return false
	}
}
