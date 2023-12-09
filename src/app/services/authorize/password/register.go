package password

import (
	"github.com/archine/ioc"
	"service-api/src/app/entity/authorize/password"
	member2 "service-api/src/app/entity/member"
	"service-api/src/app/services/authorize"
	"service-api/src/app/services/avatar"
	"service-api/src/app/services/captcha"
	"service-api/src/app/services/token"
	"service-api/src/enums/i18n"
	"service-api/src/errors"
	"service-api/src/models/ent"
	"service-api/src/models/repo/member"
	"time"
)

type RegisterService struct {
	memberModel member.Model
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
	if client.Check("register"+req.Email, req.CaptchaId, req.Captcha) {
		return nil, errors.WithMes(i18n.CaptchaErrorCheck)
	}

	_, err := member.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.WithErr(i18n.ExistsEmail, err)
	}

	sing, err := authorize.GenPasswordSing(req.Password, time.Now().String())
	if err != nil {
		return nil, errors.WithErr(i18n.ErrorSingPassword, err)
	}

	req.Password, err = authorize.Encrypt(req.Password, sing)
	if err != nil {
		return nil, errors.WithErr(i18n.ErrorSingPassword, err)
	}

	entity := &ent.Member{
		Email:        req.Email,
		Mobile:       "",
		Nickname:     req.Nickname,
		Avatar:       avatar.GenMemberAvatar(req.Email),
		PasswordSing: sing,
		Password:     req.Password,
		State:        member2.MemberStateUnActive,
		CreateTime:   time.Now(),
	}

	entity = s.memberModel.Create(entity)

	sing, err = token.Instance.Generate(&token.Claims{
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
	return &RegisterService{
		memberModel: member.Model{},
	}
}
