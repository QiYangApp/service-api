package captcha

import (
	"framework/log"
	"framework/utils"
)

func GenTokenString(t, token string) string {
	switch t {
	case "login":
		break
	default:
		log.Client().Error("captcha gen token string type not exists!")
		return ""
	}

	return utils.Md5(t + "-" + token)
}
