package admin

import (
	"service-api/src/app/helpers"

	"github.com/gin-gonic/gin"
)

type testRouteHandle struct{}

func (s testRouteHandle) Handle(r *gin.RouterGroup) {
	group := r.Group("/test")

	group.GET("/welcome", func(r *gin.Context) {
		r.JSON(200, helpers.NewResponse[string]().Single("welcome"))
	})

}
