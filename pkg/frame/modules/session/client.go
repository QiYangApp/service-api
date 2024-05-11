package session

import (
	"frame/modules/log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Storage struct {
	Store sessions.Store
	Key   string
}

func Client(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}

func Middleware(storage *Storage) gin.HandlerFunc {
	if storage == nil {
		log.Client().Panic("session middleware init fail")
		return nil
	}

	return sessions.Sessions(storage.Key, storage.Store)
}
