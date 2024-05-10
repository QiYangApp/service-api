package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
	"github.com/spf13/viper"
)

func NewStore(key string, secret []byte, cfg *viper.Viper) (sessions.Store, error) {
	switch key {
	case "cookie":
		return cookie.NewStore(secret), nil
	case "redis":
		return redis.NewStore(
			cfg.GetInt("database"),
			cfg.GetString("net"),
			cfg.GetString("addr"),
			cfg.GetString("password"),
			secret,
		)
	case "memstore":
		return memstore.NewStore(secret), nil
	default:
		return memstore.NewStore(secret), nil
	}
}
