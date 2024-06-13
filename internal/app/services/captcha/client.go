package captcha

import (
	"errors"
	"frame/util/secret"
	"service-api/internal/modules/captcha"
	"service-api/internal/modules/setting"
	"service-api/resources/translate/messages"
)

func New() captcha.Captcha {
	var client, _ = setting.GetCaptchaClient()

	return client
}

func genToken(t setting.CaptchaFeature, token string) string {
	return secret.Sha1Sum(t.ToString() + " - " + setting.SecretSetting.Key + "-" + token)
}

func Gen(t setting.CaptchaFeature, token string) (*captcha.Resp, error) {
	if err := CheckTokenType(t); err != nil {
		return nil, err
	}

	resp, err := New().Generate(genToken(t, token))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Verify(t setting.CaptchaFeature, token, key string, answer any, clear bool) (bool, error) {
	if err := CheckTokenType(t); err != nil {
		return false, err
	}

	return New().Verify(token, key, answer, clear), nil
}

func CheckTokenType(t setting.CaptchaFeature) error {
	if !setting.CaptchaSetting.Enable || !setting.CheckCaptchaFeatureEnable(t) {
		return errors.New(messages.CaptchaNotActivated.ID)
	}

	switch t {
	case setting.CaptchaFeatureSignIn:
	case setting.CaptchaFeatureSignUp:
		break
	default:
		return errors.New(messages.CaptchaTokenMissing.ID)
	}

	return nil
}
