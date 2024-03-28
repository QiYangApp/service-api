package captcha

import (
	"framework/utils"
)

type Store struct {
}

func genTokenString(t, token string) (string, error) {
	return utils.Md5(t + "-" + token), nil
}

func New() {
	//if !setting.ServiceSetting.EnableCaptcha {
	//	return "", CaptchaNotActived
	//}

}
