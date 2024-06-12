package auth

import (
	"ent/models"
	"errors"
	"frame/modules/log"
	"frame/modules/middlewares"
	"frame/modules/session"
	"frame/util/types"
	"github.com/gin-gonic/gin"
	usermodel "service-api/internal/app/repo/user"
	"service-api/internal/modules/setting"
)

type UserSession struct {
	Token      string
	UserId     int64
	RoleId     int64
	IsSigned   bool
	IsTwoFa    bool // 双重
	IsRegister bool
	Language   string
	Theme      string
}

func GetUserSession(c *gin.Context) *UserSession {
	cli := session.Client(c)

	tmp := cli.Get("user")
	if tmp == nil {
		return nil
	}

	u := &UserSession{}
	if err := types.UnmarshalJson(tmp.([]byte), u); err != nil {
		log.Sugar().Warnf("get user session fail, err: %v", err)
		return nil
	}

	return u
}

func SaveUserSession(c *gin.Context, u *UserSession) error {
	if u == nil {
		e := errors.New("save user session data empty")
		log.Sugar().Warn(e)
		return e
	}

	if str, err := types.MarshalJson(u); err == nil {
		cli := session.Client(c)
		cli.Set("user", str)
		return cli.Save()
	} else {
		log.Sugar().Warnf("save user session fail, err: %v", err)

		return err
	}
}

func SignInUserSession(ctx *gin.Context, u *models.User, remember bool) *UserSession {
	// TODO remember 记住用户待开发
	userSession := &UserSession{
		Token:      "",
		UserId:     0,
		RoleId:     0,
		IsSigned:   false,
		IsTwoFa:    false,
		IsRegister: false,
		Language:   "",
		Theme:      "",
	}

	return userSession
}

func SignInInitLanguage(ctx *gin.Context, u *models.User) error {
	if u.Language == "" {
		u.Language = setting.AppSetting.Language.String()
		err := usermodel.SetUserLanguage(ctx, u.ID, u.Language)
		if err != nil {
			return err
		}
	}

	middlewares.SetLocaleCookie(ctx, u.Language, 0)

	return nil
}
