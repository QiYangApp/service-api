package session

import (
	"framework/config"
	"framework/log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func NewStore() sessions.Store {
	secret := []byte(config.Client.GetString("session.secret"))
	storeKey := config.Client.GetString("session.store")
	if storeKey == "" {
		log.Client.Sugar().Error("session store key is empty or not exist")
		return nil
	}

	var store sessions.Store
	var err error
	switch storeKey {
	case "redis":
		conf := config.Client.GetStringMap("conns." + storeKey)
		store, err = redis.NewStore(conf["database"].(int), "tcp", conf["add"].(string), conf["password"].(string), secret)
		break
	case "cookie":
		store = cookie.NewStore(secret)
		err = nil
		break
	case "memstore":
		store = memstore.NewStore(secret)
		break
	}

	if err != nil {
		log.Client.Sugar().Error("session store key create fail")
		return nil
	}

	return store
}

func Middleware() gin.HandlerFunc {
	key := config.Client.GetString("session.key")
	return sessions.Sessions(key, NewStore())
}
