package captcha

import (
	"errors"
	"frame/util/secret"
	"service-api/configs"
	"service-api/modules/captcha"
	"service-api/resources/translate/messages"
)

func New() captcha.Captcha {
	var client, _ = configs.GetCaptchaClient()

	return client
}

func genToken(t configs.CaptchaFeature, token string) string {
	return secret.Sha1Sum(t.ToString() + " - " + configs.SecretSetting.Key + "-" + token)
}

func Gen(t configs.CaptchaFeature, token string) (*captcha.Resp, error) {
	if err := configs.IsCaptchaFeature(t); err != nil {
		return nil, err
	}

	if st := configs.IsCaptchaFeatureEnable(t); st == false {
		return nil, errors.New(messages.CaptchaNotActivated.ID)
	}

	resp, err := New().Generate(genToken(t, token))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
