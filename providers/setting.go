package providers

import "service-api/internal/modules/setting"

func SettingRegister() {
	setting.LoadCommonSettings()
}
