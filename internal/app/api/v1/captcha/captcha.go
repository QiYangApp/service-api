package captcha

import (
	"framework/log"
	"framework/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"service-api/internal/app/api/validator"
	"service-api/internal/services/captcha"
)

// Index 获取验证码
// @Summary 获取验证码
// @Description 获取验证码
// @Accept json
// @Produce  json
// @Param type path string true "验证码授权类型"
// @Param token query string true "token"
// @Success 200 {object} response.Response{Data=validator.CaptchaResponse} "请求成功"
// @Router /captcha/{type} [get]
func Index(c *gin.Context, req *validator.CaptchaRequest) {

	key, body, err := captcha.Gen(req.Type, req.Token)
	if err != nil {
		log.Client.Sugar().Info(zap.String("captcha gen err", err.Error()))
		response.RError(c, err, http.StatusNotFound, nil)
		return
	}
	response.RSuccess(c, &validator.CaptchaResponse{Id: key, Captcha: body, Token: req.Token})
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
func Verify(c *gin.Context, req *validator.CaptchaVerifyRequest) {
	st := captcha.Verify(req.Token, req.Id, req.Answer, false)
	response.RSuccess(c, st)
}
