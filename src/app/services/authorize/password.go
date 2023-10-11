package authorize

import "github.com/archine/ioc"

type PasswordLoginService struct {
	//AuthorizedService[P]
}

func (s *PasswordLoginService) Check(p any) any {
	return p
}

func (s *PasswordLoginService) Authorizing(p any) any {
	return p
}

func (s *PasswordLoginService) Authorized(p any) any {
	return p
}

func (s *PasswordLoginService) CreateBean() ioc.Bean {
	return &PasswordLoginService{}
}
