package authorize

import (
	"github.com/archine/ioc"
	"service-api/src/entity"
)

type PasswordService struct {
}

func (s *PasswordService) Login(req entity.LoginRequest) *entity.LoginResponse {
	return &entity.LoginResponse{}
}

func (s *PasswordService) CreateBean() ioc.Bean {
	return &PasswordService{}
}
