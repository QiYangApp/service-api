package passwd

import (
	"service-api/internal/app/http/validator"
)

type PasswordSignUpService struct {
}

// PreAuthData 申请注册数据
func (s *PasswordSignUpService) PreAuthData(req validator.PreAuthDataRequest) any {
	return true
}

// PreAuthVerify 账号存在校验
func (s *PasswordSignUpService) PreAuthVerify(req validator.PreAuthVerifyRequest) any {
	return true
}

// Authorizing 发送邮件校验
func (s *PasswordSignUpService) Authorizing(req validator.AuthorizingRequest) any {
	return true
}

// Authorized 点击邮箱确认账号
func (s *PasswordSignUpService) Authorized(req validator.AuthorizedRequest) any {
	return true
}
