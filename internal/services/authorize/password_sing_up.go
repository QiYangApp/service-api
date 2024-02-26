package authorize

import (
	"service-api/internal/entity"
)

type PasswordSignUpService struct {
}

// PreAuthorizationData 申请注册数据
func (s *PasswordSignUpService) PreAuthorizationData(req entity.LoginRequest) *entity.LoginResponse {
	return &entity.LoginResponse{}
}

// PreAuthorizationVerification 账号存在校验
func (s *PasswordSignUpService) PreAuthorizationVerification(req entity.LoginRequest) *entity.LoginResponse {
	return &entity.LoginResponse{}
}

// Authorizing 发送邮件校验
func (s *PasswordSignUpService) Authorizing(req entity.LoginRequest) *entity.LoginResponse {
	return &entity.LoginResponse{}
}

// Authorized 点击邮箱确认账号
func (s *PasswordSignUpService) Authorized(req entity.LoginRequest) *entity.LoginResponse {
	return &entity.LoginResponse{}
}
