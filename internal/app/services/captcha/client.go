package captcha

import (
	"errors"
	"frame/util/secret"
	setting2 "service-api/conf"
	"service-api/modules/captcha"
	"service-api/resources/translate/messages"
)

func New() captcha.Captcha {
	var client, _ = setting2.GetCaptchaClient()

	return client
}

func genToken(t setting2.CaptchaFeature, token string) string {
	return secret.Sha1Sum(t.ToString() + " - " + setting2.SecretSetting.Key + "-" + token)
}

func Gen(t setting2.CaptchaFeature, token string) (*captcha.Resp, error) {
	if err := CheckTokenType(t); err != nil {
		return nil, err
	}

	resp, err := New().Generate(genToken(t, token))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Verify(t setting2.CaptchaFeature, token, key string, answer any, clear bool) (bool, error) {
	if err := CheckTokenType(t); err != nil {
		return false, err
	}

	return New().Verify(token, key, answer, clear), nil
}

func CheckTokenType(t setting2.CaptchaFeature) error {
	if !setting2.CaptchaSetting.Enable || !setting2.CheckCaptchaFeatureEnable(t) {
		return errors.New(messages.CaptchaNotActivated.ID)
	}

	switch t {
	case setting2.CaptchaFeatureSignIn:
	case setting2.CaptchaFeatureSignUp:
		break
	default:
		return errors.New(messages.CaptchaTokenMissing.ID)
	}

	return nil
}
