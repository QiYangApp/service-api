package controller

import (
	"github.com/gin-gonic/gin"
	"service-api/api/http/validator"
)

type CaptchaApi struct {
}

func (*CaptchaApi) Index(c *gin.Context, req *validator.CaptchaForm) *gin.Context {

	return c
}

func (*CaptchaApi) Check(c *gin.Context, req *validator.CaptchaForm) *gin.Context {
	return c
}
