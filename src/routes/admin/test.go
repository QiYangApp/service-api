package admin

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"service-api/src/app/helpers/response"
	"service-api/src/app/services/token"
)

type testRouteHandle struct{}

func (s testRouteHandle) Handle(r *gin.RouterGroup) {
	group := r.Group("/test")

	group.GET("/welcome", func(r *gin.Context) {
		r.JSON(200, response.SingleSuccess[string]("welcome"))
	})

	group.GET("/token/generate", func(r *gin.Context) {
		t, e := token.NewTokenService().Generate(&token.Claims{
			Context: "test",
		})

		if e != nil {
			zap.S().Info(e)
		}

		r.JSON(200, response.SingleSuccess[string](t))
	})

	group.GET("/token/parser", func(r *gin.Context) {
		t, e := token.NewTokenService().Parser(r.Query("token"))

		if e != nil {
			zap.S().Info(e)
		}

		r.JSON(200, response.SingleSuccess[*token.Claims](t))
	})

}
