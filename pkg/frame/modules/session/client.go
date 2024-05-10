package session

import (
	"frame/conf"
	"frame/modules/log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"sync"
)

var st = cookie.NewStore([]byte("secret"))

type FactoryStruct struct {
	Store    sessions.Store
	Security string
	Key      string
}

var factory = &FactoryStruct{}
var once = sync.Once{}

func Factory() *FactoryStruct {
	once.Do(func() {
		sessCfg := conf.Client().Sub("session")
		factory.Security = conf.Client().GetString("security.key")
		factory.Key = sessCfg.GetString("key")

		if st, err := NewStore(factory.Key, []byte(factory.Security), sessCfg); err != nil {
			log.Sugar().Error("session init store fail, err: ", err)
		} else {
			factory.Store = st
		}
	})

	return factory
}

func Client(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}
