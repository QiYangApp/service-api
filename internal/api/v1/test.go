package v1

import (
	"frame/modules/log"
	"github.com/gin-gonic/gin"
	"service-api/internal/services/auth"
)

func TestSession(c *gin.Context) {
	u := &auth.UserSession{UserId: 1}
	log.Sugar().Info(auth.GetUserSession(c))
	log.Sugar().Info(auth.SaveUserSession(c, u))
	log.Sugar().Info(auth.GetUserSession(c))
}
