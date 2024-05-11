package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Storage struct {
	Store    sessions.Store
	Security string
	Key      string
}

func Client(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}

func Middleware(storage Storage) gin.HandlerFunc {
	return sessions.Sessions(storage.Key, storage.Store)
}
