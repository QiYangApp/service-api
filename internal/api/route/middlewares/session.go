package middlewares

import (
	"frame/modules/session"
	"github.com/gin-gonic/gin"
	"service-api/internal/modules/setting"
)

func Session() gin.HandlerFunc {
	return session.Middleware(setting.NewSessionStorage())
}
