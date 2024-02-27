package v1

import (
	"github.com/gin-gonic/gin"
	"service-api/internal/app/http/validator"
)

type AuthApi struct {
}

// PreAuthorizationData 申请注册数据
func (s *AuthApi) PreAuthorizationData(c *gin.Context, req validator.PreAuthDataRequest) *gin.Context {
	return c
}

// PreAuthorizationVerify PreAuthorizationVerification 账号存在校验
func (s *AuthApi) PreAuthorizationVerify(c *gin.Context, req validator.PreAuthVerifyRequest) *gin.Context {
	return c
}

// Authorizing 发送邮件校验
func (s *AuthApi) Authorizing(c *gin.Context, req validator.AuthorizingRequest) *gin.Context {
	return c
}

// Authorized 点击邮箱确认账号
func (s *AuthApi) Authorized(c *gin.Context, req validator.AuthorizedRequest) *gin.Context {
	return c
}
