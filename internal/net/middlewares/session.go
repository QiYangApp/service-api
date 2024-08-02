package middlewares

import (
	"frame/modules/session"
	"github.com/gin-gonic/gin"
	"service-api/conf"
)

func Session() gin.HandlerFunc {
	return session.Middleware(conf.NewSessionStorage())
}
