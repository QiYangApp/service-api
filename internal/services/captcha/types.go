package captcha

import (
	"errors"
	"framework/utils"
)

var TokenTypeNotExists = errors.New("captcha gen token string type not exists!")

func GenTokenString(t, token string) (string, error) {
	switch t {
	case "login":
		break
	default:
		return "", TokenTypeNotExists
	}

	return utils.Md5(t + "-" + token), nil
}
