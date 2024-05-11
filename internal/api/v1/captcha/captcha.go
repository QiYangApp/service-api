package captcha

import (
	"frame/modules/log"
	"frame/modules/resp"
	"net/http"
	"service-api/internal/app/services/captcha"
	"service-api/internal/app/validator"
	"service-api/internal/modules/setting"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

	body, err := captcha.Gen(setting.CaptchaFeature(req.Type), req.Token)
	if err != nil {
		log.Sugar().Info(zap.String("captcha gen err", err.Error()))
		resp.Error(c, err, http.StatusNotFound, nil)
		return
	}
	resp.Success(c, &validator.CaptchaResponse{Id: body.GetKey(), Captcha: body.GetBody(), Token: body.GetToken()})
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
	st, err := captcha.Verify(setting.CaptchaFeature(req.Type), req.Token, req.Id, req.Answer, false)
	if err != nil {
		log.Sugar().Info(zap.String("captcha verify err", err.Error()))
		resp.Error(c, err, http.StatusNotFound, nil)
		return
	}
	resp.Success(c, st)
}
