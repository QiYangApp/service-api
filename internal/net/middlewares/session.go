package middlewares

import (
	"frame/modules/session"
	"github.com/gin-gonic/gin"
	"service-api/configs"
)

func Session() gin.HandlerFunc {
	return session.Middleware(configs.NewSessionStorage())
}
