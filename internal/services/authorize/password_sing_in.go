package authorize

import (
	"service-api/internal/entity"
)

type PasswordSingInService struct {
}

func (s *PasswordSingInService) PreAuthorizationData(req entity.LoginRequest) *entity.LoginResponse {
	return &entity.LoginResponse{}
}

func (s *PasswordSingInService) PreAuthorizationVerification(req entity.LoginRequest) *entity.LoginResponse {
	return &entity.LoginResponse{}
}

func (s *PasswordSingInService) Authorizing(req entity.LoginRequest) *entity.LoginResponse {
	return &entity.LoginResponse{}
}

func (s *PasswordSingInService) Authorized(req entity.LoginRequest) *entity.LoginResponse {
	return &entity.LoginResponse{}
}
