package admin

import (
	"service-api/src/app/services/token"
	"service-api/src/core/helpers/response"
	"service-api/src/core/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type testRouteHandle struct{}

func (s testRouteHandle) Handle(r *gin.RouterGroup) {
	group := r.Group("/test")

	group.GET("/welcome", func(r *gin.Context) {
		r.JSON(200, response.SingleSuccess[string](r, "welcome"))
	})

	group.GET("/token/generate", func(r *gin.Context) {
		t, e := token.NewTokenService().Generate(&token.Claims{
			Context: "test",
		})

		if e != nil {
			logger.R().Info("admin", zap.Error(e))
		}

		r.JSON(200, response.SingleSuccess[string](r, t))
	})

	group.GET("/token/parser", func(r *gin.Context) {
		t, e := token.NewTokenService().Parser(r.Query("token"))

		if e != nil {
			logger.R().Info("fail", zap.Error(e))
		}

		r.JSON(200, response.SingleSuccess[*token.Claims](r, t))
	})

}
