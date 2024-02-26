package captcha

import (
	"framework/helpers"
	"framework/log"
)

func GenTokenString(t, token string) string {
	switch t {
	case "login":
		break
	default:
		log.Client().Error("captcha gen token string type not exists!")
		return ""
	}

	return helpers.Md5(t + "-" + token)
}
