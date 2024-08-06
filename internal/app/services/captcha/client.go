package captcha

import (
	"errors"
	"frame/util/secret"
	"service-api/conf"
	"service-api/modules/captcha"
	"service-api/resources/translate/messages"
)

func New() captcha.Captcha {
	var client, _ = conf.GetCaptchaClient()

	return client
}

func genToken(t conf.CaptchaFeature, token string) string {
	return secret.Sha1Sum(t.ToString() + " - " + conf.SecretSetting.Key + "-" + token)
}

func Gen(t conf.CaptchaFeature, token string) (*captcha.Resp, error) {
	if err := conf.IsCaptchaFeature(t); err != nil {
		return nil, err
	}

	if st := conf.IsCaptchaFeatureEnable(t); st == false {
		return nil, errors.New(messages.CaptchaNotActivated.ID)
	}

	resp, err := New().Generate(genToken(t, token))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
