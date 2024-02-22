package controller

import (
	"errors"
	"framework/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"service-api/api/http/validator"
	"service-api/resources/lang"
	"service-api/services/captcha"
)

type CaptchaApi struct {
}

func (*CaptchaApi) Index(c *gin.Context, req *validator.CaptchaRequest) *gin.Context {
	img, err := captcha.NewImage().Generate(req.Token, nil)
	if err != nil {
		return response.RError(c, errors.New(lang.CaptchaGenerate), http.StatusInternalServerError, nil).ToJson()
	}

	return response.RSuccess(c, &validator.CaptchaResponse{Id: img.Id, Captcha: img.Captcha}).ToJson()
}

func (*CaptchaApi) Verify(c *gin.Context, req *validator.CaptchaVerifyRequest) *gin.Context {
	st := captcha.NewImage().Verify(req.Token, req.Id, req.Answer)

	return response.RSuccess(c, st).ToJson()
}
