// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
	"errors"
	"frame/modules/log"
	"frame/modules/middlewares"
	"frame/modules/session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
	"github.com/spf13/viper"
	"net/http"
)

var SessionSetting = &struct {
	Key           string `mapstructure:"key"`
	OriginalStore string
	Store         string `mapstructure:"store"`
	// Cookie domain name. Default is false.
	EnableCookieDomain bool `mapstructure:"enable_cookie_domain"`
	// Cookie domain default is empty
	CookieDomain string
	CookieMaxAge int    `mapstructure:"cookie_max_age"`
	CookieName   string `mapstructure:"cookie_name"`
	// Cookie path to store. Default is "/".
	CookiePath string `mapstructure:"cookie_path"`
	// GC interval time in seconds. Default is 3600.
	Gclifetime int64
	// Max life time in seconds. Default is whatever GC interval time is.
	Maxlifetime int64
	// Use HTTPS only. Default is false.
	Secure   bool `mapstructure:"secure"`
	HttpOnly bool `mapstructure:"http_only"`
	// SameSite declares if your cookie should be restricted to a first-party or same-site context. Valid strings are "none", "lax", "strict". Default is "lax"
	SameSite http.SameSite `mapstructure:"same_site"`
}{
	HttpOnly:    true,
	CookieName:  "QiYang",
	Gclifetime:  86400,
	Maxlifetime: 86400,
	SameSite:    http.SameSiteLaxMode,
}

func loadSessionSetting(cfg *viper.Viper) {
	if err := cfg.Unmarshal(SessionSetting); err != nil {
		log.Sugar().Error("load service setting")
	}

	if SessionSetting.EnableCookieDomain {
		SessionSetting.CookieDomain = AppSetting.Domain
	}

	SessionSetting.OriginalStore = SessionSetting.Store

	loadCookieSetting()
}

func loadCookieSetting() {

	middlewares.SettingCookieConf(
		SessionSetting.CookiePath,
		SessionSetting.CookieDomain,
		SessionSetting.Secure,
		SessionSetting.HttpOnly,
	)
}

func NewSessionStorage() *session.Storage {
	store, err := newSessionStore()
	if err != nil {
		log.Sugar().Error("new session storage fail")
		return nil
	}

	store.Options(
		sessions.Options{
			Path:     SessionSetting.CookiePath,
			Domain:   SessionSetting.CookieDomain,
			MaxAge:   SessionSetting.CookieMaxAge,
			Secure:   SessionSetting.Secure,
			HttpOnly: SessionSetting.HttpOnly,
			SameSite: SessionSetting.SameSite,
		},
	)

	return &session.Storage{
		Store: store,
		Key:   SessionSetting.Key,
	}
}

func newSessionStore() (sessions.Store, error) {
	var conn = GetConnConf(SessionSetting.Store)
	if conn == nil {
		err := errors.New("load session store fail")
		log.Sugar().Error(err)
		return nil, err
	}

	drive := conn.GetString("driver")
	if drive == "" {
		err := errors.New("session store drive not setting")
		log.Sugar().Error(err)
		return nil, err
	}
	secret := []byte(SecretSetting.Key)

	switch drive {
	case "cookie":
		return cookie.NewStore(secret), nil
	case "redis":
		return redis.NewStore(
			conn.GetInt("database"),
			conn.GetString("net"),
			conn.GetString("addr"),
			conn.GetString("password"),
			secret,
		)
	case "memstore":
		return memstore.NewStore(secret), nil
	default:
		return memstore.NewStore(secret), nil
	}

}
