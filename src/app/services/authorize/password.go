package authorize

import (
	"crypto/md5"
	"fmt"
	"github.com/archine/ioc"
	"golang.org/x/crypto/bcrypt"
	"service-api/src/app/entity/authorize/request"
	"service-api/src/app/services/captcha"
	"service-api/src/app/services/token"
	"service-api/src/enums/i18n"
	"service-api/src/errors"
	"service-api/src/models/repo/authoize"
)

const bcryptDefaultCost = 31

func Encrypt(password, sing string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcryptDefaultCost)
	if err != nil {
		return "", err
	}

	password = fmt.Sprintf("%x-%s", p, sing)

	return fmt.Sprintf("%x", md5.Sum([]byte(password))), nil
}

func GenPasswordSing(password, sing string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcryptDefaultCost)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", p), nil
}

type PasswordLoginService struct {
}

// Check 判断账号是否存在
func (s *PasswordLoginService) Check(req request.PasswordLoginCheckReq) (*request.PasswordLoginCheckRsp, error) {
	if req.Account == "" {
		return nil, errors.WithMes(i18n.EmptyAccount)
	}

	if !authoize.AccountExists(req.Account) {
		return nil, errors.WithMes(i18n.NotExistsAccount)
	}

	return &request.PasswordLoginCheckRsp{
		State: true,
	}, nil
}

func (s *PasswordLoginService) Authorizing(p request.PasswordLoggingReq) (*request.PasswordLoggingRsp, error) {
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

func (s *PasswordLoginService) Authorized(req request.PasswordLoggedReq) (*request.PasswordLoggedRsp, error) {
	client := captcha.NewImage()
	if client.Check(req.CaptchaId, req.Captcha) {
		return nil, errors.WithMes(i18n.CaptchaErrorCheck)
	}

	member, err := authoize.FirstMemberByAccount(req.Account)
	if err != nil {
		return nil, errors.WithErr(i18n.NotExistsAccount, err)
	}

	var password string
	password, err = Encrypt(req.Password, member.PasswordSing)
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

func (s *PasswordLoginService) CreateBean() ioc.Bean {
	return &PasswordLoginService{}
}

type PasswordRegisterService struct {
}

// Check 判断账号是否存在
func (s *PasswordRegisterService) Check(req request.PasswordLoginCheckReq) (*request.PasswordLoginCheckRsp, error) {
	if req.Account == "" {
		return nil, errors.WithMes(i18n.EmptyAccount)
	}

	if !authoize.AccountExists(req.Account) {
		return nil, errors.WithMes(i18n.NotExistsAccount)
	}

	return &request.PasswordLoginCheckRsp{
		State: true,
	}, nil
}

func (s *PasswordRegisterService) Authorizing(p request.PasswordLoggingReq) (*request.PasswordLoggingRsp, error) {
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

func (s *PasswordRegisterService) Authorized(req request.PasswordLoggedReq) (*request.PasswordLoggedRsp, error) {
	client := captcha.NewImage()
	if client.Check(req.CaptchaId, req.Captcha) {
		return nil, errors.WithMes(i18n.CaptchaErrorCheck)
	}

	member, err := authoize.FirstMemberByAccount(req.Account)
	if err != nil {
		return nil, errors.WithErr(i18n.NotExistsAccount, err)
	}

	var password string
	password, err = Encrypt(req.Password, member.PasswordSing)
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

func (s *PasswordRegisterService) CreateBean() ioc.Bean {
	return &PasswordRegisterService{}
}
