package password

import (
	"github.com/archine/ioc"
	"service-api/src/app/entity/authorize/password"
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
func (s *RegisterService) Check(req password.RegisterCheckReq) (*password.RegisterCheckRsp, error) {
	if req.Account == "" {
		return nil, errors.WithMes(i18n.EmptyAccount)
	}

	if req.Email == "" {
		return nil, errors.WithMes(i18n.EmptyEmail)
	}

	if authoize.AccountAndEmailExists(req.Account, req.Email) {
		return nil, errors.WithMes(i18n.ExistsAccount)
	}

	return &password.RegisterCheckRsp{
		State: true,
	}, nil
}

func (s *RegisterService) Authorizing(p password.RegisteringReq) (*password.RegisteringRsp, error) {
	client := captcha.NewImage()
	body, err := client.Generate("register", nil)
	if err != nil {
		return nil, err
	}

	return &password.RegisteringRsp{
		Id:      body.Id,
		Captcha: body.Captcha,
	}, nil
}

func (s *RegisterService) Authorized(req password.RegisteredReq) (*password.RegisteredRsp, error) {
	client := captcha.NewImage()
	if client.Check("register", req.CaptchaId, req.Captcha) {
		return nil, errors.WithMes(i18n.CaptchaErrorCheck)
	}

	member, err := authoize.FirstMemberByAccount(req.Account)
	if err != nil {
		return nil, errors.WithErr(i18n.NotExistsAccount, err)
	}

	var pwd string
	pwd, err = authorize.Encrypt(req.Password, member.PasswordSing)
	if err != nil {
		return nil, errors.WithErr(i18n.ErrorSingPassword, err)
	}

	if !authoize.ExistsBYPassword(req.Account, pwd) {
		return nil, errors.WithMes(i18n.ErrorPassword)
	}

	sing, err := token.Instance.Generate(&token.Claims{
		Context: member,
	})

	if err != nil {
		return nil, err
	}

	return &password.RegisteredRsp{
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
