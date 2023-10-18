package authorize

import (
	"github.com/archine/ioc"
	"service-api/src/app/entity/authorize/request"
)

type PasswordLoginService struct {
}

func (s *PasswordLoginService) Check(p request.PasswordLoginCheckReq) request.PasswordLoginCheckRsp {
	return request.PasswordLoginCheckRsp{}
}

func (s *PasswordLoginService) Authorizing(p request.PasswordLoggingReq) request.PasswordLoggingRsp {
	return request.PasswordLoggingRsp{}
}

func (s *PasswordLoginService) Authorized(p request.PasswordLoggedReq) request.PasswordLoggedRsp {
	return request.PasswordLoggedRsp{}
}

func (s *PasswordLoginService) CreateBean() ioc.Bean {
	return &PasswordLoginService{}
}
