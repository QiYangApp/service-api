package auth

import (
	"ent/models"
	"errors"
	"frame/errs"
	"frame/modules/log"
	"frame/modules/middlewares"
	"frame/modules/session"
	"frame/util/types"
	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	usermodel "service-api/internal/app/repo/user"
	"service-api/internal/modules/setting"
	"service-api/resources/translate/messages"
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

func SignInUserSession(ctx *gin.Context, u *models.User, remember bool) (*UserSession, error) {

	err := SignInInitLanguage(ctx, u)
	if err != nil {
		log.Sugar().Errorf("sign in init user default lang fail, err: %v", err)
		return nil, errs.New(messages.UserSignInInitLanguageFail.ID)
	}

	err = SignInInitTheme(ctx, u)
	if err != nil {
		log.Sugar().Errorf("sign in init user default theme fail, err: %v", err)
		return nil, errs.New(messages.UserSignInInitLanguageFail.ID)
	}

	if i18n.GetLngHandler() != u.Language {
		ctx.Locale = middleware.Locale(ctx.Resp, ctx.Req)
	}

	// TODO remember 记住用户待开发
	userSession := &UserSession{
		Token:      "",
		UserId:     u.ID,
		RoleId:     0,
		IsSigned:   true,
		IsTwoFa:    false,
		IsRegister: false,
		Language:   u.Language,
		Theme:      u.Theme,
	}

	return userSession, nil
}

func SignInInitTheme(ctx *gin.Context, u *models.User) error {
	if u.Theme == "" {
		u.Theme = setting.AppSetting.Theme
		err := usermodel.SetUserLanguage(ctx, u.ID, u.Theme)
		if err != nil {
			return err
		}
	}

	return nil
}

func SignInInitLanguage(ctx *gin.Context, u *models.User) error {
	if u.Language == "" {
		u.Language = setting.AppSetting.Language
		err := usermodel.SetUserLanguage(ctx, u.ID, u.Language)
		if err != nil {
			return err
		}
	}

	middlewares.SetLocaleCookie(ctx, u.Language, 0)

	return nil
}
