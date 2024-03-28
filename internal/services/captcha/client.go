package captcha

import (
	"framework/utils"
	"service-api/internal/modules/captcha"
	"service-api/internal/modules/setting"
)

func New() captcha.Captcha {
	var client, _ = setting.GetCaptchaClient()

	return client
}

func genTokenString(t, token string) string {
	return utils.Md5(t + "-" + token)
}

func Gen(t, token string) (string, any, error) {
	resp, err := New().Generate(genTokenString(t, token))
	if err != nil {
		return "", nil, err
	}

	return resp.GetKey(), resp.GetBody(), nil
}

func Verify(token, key string, answer any, clear bool) bool {
	return New().Verify(token, key, answer, clear)
}
