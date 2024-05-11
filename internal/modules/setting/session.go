package setting

import (
	"errors"
	"frame/modules/log"
	"frame/modules/session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
	"github.com/spf13/viper"
)

var SessionSetting = &struct {
	Store string `mapstructure:"store"`
	Key   string `mapstructure:"key"`
}{}

func loadSessionSetting(cfg *viper.Viper) {
	if err := cfg.Unmarshal(SessionSetting); err != nil {
		log.Sugar().Error("load service setting")
	}
}

func NewSessionStorage() *session.Storage {
	store, err := newSessionStore()
	if err != nil {
		log.Sugar().Error("new session storage fail")
		return nil
	}

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
