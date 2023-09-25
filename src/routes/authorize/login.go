package authorize

import (
	"github.com/gin-gonic/gin"
	"service-api/src/api/controller/authorize"
	"service-api/src/app/entity/http/authorize/request"

	"service-api/src/core/helpers/routes"
)

type loginRouteHandle struct{}

func (a loginRouteHandle) Handle(r *gin.RouterGroup) {
	route := r.Group("/login")

	a.password(route)
}

func (a loginRouteHandle) password(r *gin.RouterGroup) {
	route := r.Group("/password")
	controller := authorize.PasswordController[request.PasswordLoginVerify]{}

	route.GET("/mobile", routes.Bind(controller.Authorized))
}
