package captcha

import "errors"

var TokenTypeNotExists = errors.New("captcha gen token string type not exists")
var CaptchaNotActived = errors.New("captcha is not activated")
