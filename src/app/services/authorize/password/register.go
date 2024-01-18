package password

import (
	"github.com/archine/ioc"
	"service-api/src/app/entity/authorize/password"
	member2 "service-api/src/app/entity/member"
	i18n2 "service-api/src/app/enums/i18n"
	"service-api/src/app/errors"
	"service-api/src/app/services/authorize"
	"service-api/src/app/services/captcha"
	"service-api/src/app/services/common"
	"service-api/src/app/services/token"
	"service-api/src/models/ent"
	"service-api/src/models/repo/member"
	"time"
)

type RegisterService struct {
}

func (s *RegisterService) Authorizing(req password.RegisteringReq) (*password.RegisteringRsp, error) {
	client := captcha.NewImage()
	body, err := client.Generate("register"+req.Email, nil)
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
	if client.Verify("register"+req.Email, req.CaptchaId, req.Captcha) {
		return nil, errors.WithMes(i18n2.CaptchaErrorCheck)
	}

	_, err := member.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.WithErr(i18n2.ExistsEmail, err)
	}

	sing, err := authorize.GenPasswordSing(req.Password, time.Now().String())
	if err != nil {
		return nil, errors.WithErr(i18n2.ErrorSingPassword, err)
	}

	req.Password, err = authorize.Encrypt(req.Password, sing)
	if err != nil {
		return nil, errors.WithErr(i18n2.ErrorSingPassword, err)
	}

	entity := &ent.Member{
		Email:        req.Email,
		Mobile:       "",
		Nickname:     req.Nickname,
		Avatar:       common.GenMemberAvatar(req.Email),
		PasswordSing: sing,
		Password:     req.Password,
		State:        member2.MemberStateUnActive,
		CreateTime:   time.Now(),
	}

	entity = member.AddSingleByRegister(entity)

	sing, err = token.Instance().Generate(&token.Claims{
		Context: entity,
	})

	if err != nil {
		return nil, err
	}

	return &password.RegisteredRsp{
		MemberId: entity.ID,
		Nickname: entity.Nickname,
		Avatar:   entity.Avatar,
		Email:    entity.Email,
		Token:    sing,
	}, nil
}

func (s *RegisterService) CreateBean() ioc.Bean {
	return &RegisterService{}
}
