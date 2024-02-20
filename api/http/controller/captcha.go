package controller

import (
	"github.com/gin-gonic/gin"
	"service-api/api/http/validator"
)

type CaptchaApi struct {
}

func (*CaptchaApi) Index(c *gin.Context, req *validator.CaptchaFrom) *gin.Context {
	return c
}
