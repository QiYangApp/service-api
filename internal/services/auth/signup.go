package auth

import "service-api/configs"

func IsSignUpEnable() bool {

	if configs.AppSetting.SignUpEnable {

	}

	return false
}
