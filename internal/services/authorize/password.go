package authorize

import (
	"github.com/archine/ioc"
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

func (s *PasswordSingInService) CreateBean() ioc.Bean {
	return &PasswordSingInService{}
}

type PasswordSignUpService struct {
}

func (s *PasswordSignUpService) PreAuthorizationData(req entity.LoginRequest) *entity.LoginResponse {
	return &entity.LoginResponse{}
}

func (s *PasswordSignUpService) PreAuthorizationVerification(req entity.LoginRequest) *entity.LoginResponse {
	return &entity.LoginResponse{}
}

func (s *PasswordSignUpService) Authorizing(req entity.LoginRequest) *entity.LoginResponse {
	return &entity.LoginResponse{}
}

func (s *PasswordSignUpService) Authorized(req entity.LoginRequest) *entity.LoginResponse {
	return &entity.LoginResponse{}
}

func (s *PasswordSignUpService) CreateBean() ioc.Bean {
	return &PasswordSignUpService{}
}
