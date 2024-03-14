package providers

import "service-api/internal/modules/setting"

type Setting struct {
}

func (*Setting) Register() {
	setting.LoadCommonSettings()
}
