package authorize

import (
	"github.com/gin-gonic/gin"
	"service-api/src/api/controller/authorize/login"
	"service-api/src/core/helpers/routes"
)

type loginRouteHandle struct{}

func (a loginRouteHandle) Handle(r *gin.RouterGroup) {
	route := r.Group("/login")

	a.password(route)
}

func (a loginRouteHandle) password(r *gin.RouterGroup) {
	route := r.Group("/password")
	controller := login.PasswordController{}

	route.GET("/mobile", routes.Bind(controller.Edit))
}
