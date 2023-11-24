package password

import (
	"github.com/archine/ioc"
	"service-api/src/app/entity/authorize/request"
	"service-api/src/app/services/authorize"
	"service-api/src/app/services/captcha"
	"service-api/src/app/services/token"
	"service-api/src/enums/i18n"
	"service-api/src/errors"
	"service-api/src/models/repo/authoize"
)

type RegisterService struct {
}

// Check 判断账号是否存在
func (s *RegisterService) Check(req request.PasswordRegisterCheckReq) (*request.PasswordRegisterCheckRsp, error) {
	if req.Account == "" {
		return nil, errors.WithMes(i18n.EmptyAccount)
	}

	if req.Email == "" {
		return nil, errors.WithMes(i18n.EmptyEmail)
	}

	if authoize.AccountAndEmailExists(req.Account, req.Email) {
		return nil, errors.WithMes(i18n.ExistsAccount)
	}

	return &request.PasswordRegisterCheckRsp{
		State: true,
	}, nil
}

func (s *RegisterService) Authorizing(p request.PasswordLoggingReq) (*request.PasswordLoggingRsp, error) {
	client := captcha.NewImage()
	body, err := client.Generate(nil)
	if err != nil {
		return nil, err
	}

	return &request.PasswordLoggingRsp{
		Id:      body.Id,
		Captcha: body.Captcha,
	}, nil
}

func (s *RegisterService) Authorized(req request.PasswordLoggedReq) (*request.PasswordLoggedRsp, error) {
	client := captcha.NewImage()
	if client.Check(req.CaptchaId, req.Captcha) {
		return nil, errors.WithMes(i18n.CaptchaErrorCheck)
	}

	member, err := authoize.FirstMemberByAccount(req.Account)
	if err != nil {
		return nil, errors.WithErr(i18n.NotExistsAccount, err)
	}

	var password string
	password, err = authorize.Encrypt(req.Password, member.PasswordSing)
	if err != nil {
		return nil, errors.WithErr(i18n.ErrorSingPassword, err)
	}

	if !authoize.ExistsBYPassword(req.Account, password) {
		return nil, errors.WithMes(i18n.ErrorPassword)
	}

	sing, err := token.Instance.Generate(&token.Claims{
		Context: member,
	})

	if err != nil {
		return nil, err
	}

	return &request.PasswordLoggedRsp{
		MemberId: member.ID,
		Nickname: member.Nickname,
		Avatar:   member.Avatar,
		Account:  member.Account,
		Token:    sing,
	}, nil
}

func (s *RegisterService) CreateBean() ioc.Bean {
	return &RegisterService{}
}
