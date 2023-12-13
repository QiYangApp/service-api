package password

import (
	"github.com/archine/ioc"
	"service-api/src/app/entity/authorize/password"
	"service-api/src/app/services/authorize"
	"service-api/src/app/services/captcha"
	"service-api/src/app/services/token"
	"service-api/src/enums/i18n"
	"service-api/src/errors"
	"service-api/src/models/repo/member"
)

type LoginService struct {
}

// Check 判断账号是否存在
func (s *LoginService) Check(req password.LoginCheckReq) (*password.LoginCheckRsp, error) {
	if req.Email == "" {
		return nil, errors.WithMes(i18n.EmptyEmail)
	}

	if !member.EmailByExists(req.Email) {
		return nil, errors.WithMes(i18n.NotExistsEmail)
	}

	return &password.LoginCheckRsp{
		State: true,
	}, nil
}

func (s *LoginService) Authorizing(req password.LoggingReq) (*password.LoggingRsp, error) {
	client := captcha.NewImage()
	body, err := client.Generate("login"+req.Email, nil)
	if err != nil {
		return nil, err
	}

	return &password.LoggingRsp{
		Id:      body.Id,
		Captcha: body.Captcha,
	}, nil
}

func (s *LoginService) Authorized(req password.LoggedReq) (*password.LoggedRsp, error) {
	client := captcha.NewImage()
	if client.Check("login"+req.Account, req.CaptchaId, req.Captcha) {
		return nil, errors.WithMes(i18n.CaptchaErrorCheck)
	}

	mb, err := member.FindByEmail(req.Account)
	if err != nil {
		return nil, errors.WithErr(i18n.NotExistsEmail, err)
	}

	var pwd string
	pwd, err = authorize.Encrypt(req.Password, mb.PasswordSing)
	if err != nil {
		return nil, errors.WithErr(i18n.ErrorSingPassword, err)
	}

	if !member.ExistsBYPassword(req.Account, pwd) {
		return nil, errors.WithMes(i18n.ErrorPassword)
	}

	sing, err := token.Instance().Generate(&token.Claims{
		Context: mb,
	})

	if err != nil {
		return nil, err
	}

	return &password.LoggedRsp{
		MemberId: mb.ID,
		Nickname: mb.Nickname,
		Avatar:   mb.Avatar,
		Email:    mb.Email,
		Token:    sing,
	}, nil
}

func (s *LoginService) CreateBean() ioc.Bean {
	return &LoginService{}
}
