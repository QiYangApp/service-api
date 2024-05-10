package middlewares

import (
	"frame/modules/session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Session() gin.HandlerFunc {
	return sessions.Sessions(session.Factory().Key, session.Factory().Store)
}
