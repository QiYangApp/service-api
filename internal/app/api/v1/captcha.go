package v1

import (
	"errors"
	"framework/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"service-api/internal/app/api/validator"
	"service-api/internal/services/captcha"
	"service-api/resources/lang"
)

type CaptchaApi struct {
}

// Index 获取验证码
// @Summary 获取验证码
// @Description 获取验证码
// @Accept json
// @Produce  json
// @Param type path string true "验证码授权类型"
// @Param token query string true "token"
// @Success 200 {object} response.Response{Data=validator.CaptchaResponse} "请求成功"
// @Router /captcha/{type} [get]
func (*CaptchaApi) Index(c *gin.Context, req *validator.CaptchaRequest) *gin.Context {
	token := captcha.GenTokenString(req.Type, req.Token)
	img, err := captcha.NewImage().Generate(token, nil)
	if err != nil {
		return response.RError(c, errors.New(lang.CaptchaGenerate), http.StatusInternalServerError, nil).ToJson()
	}

	return response.RSuccess(c, &validator.CaptchaResponse{Id: img.Id, Captcha: img.Captcha, Token: token}).ToJson()
}

// Verify 校验验证码
// @Summary 校验验证码
// @Description 校验验证码
// @Accept json
// @Produce  json
// @Param type path string true "验证码授权类型"
// @Param req body validator.CaptchaVerifyRequest true "data"
// @Success 200 {object} response.Response{Data=boolean} "请求成功"
// @Router /captcha/{type} [post]
func (*CaptchaApi) Verify(c *gin.Context, req *validator.CaptchaVerifyRequest) *gin.Context {
	st := captcha.NewImage().Verify(req.Token, req.Id, req.Answer, false)

	return response.RSuccess(c, st).ToJson()
}
