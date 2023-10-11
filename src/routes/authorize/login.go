package authorize

import (
	"github.com/gin-gonic/gin"
	"service-api/src/api/controller/authorize"
	"service-api/src/core/helpers/routes"
)

type loginRouteHandle struct{}

func (a loginRouteHandle) Handle(r *gin.RouterGroup) {
	route := r.Group("/login")

	a.password(route)
}

func (a loginRouteHandle) password(r *gin.RouterGroup) {
	route := r.Group("/password")
	c := authorize.NewPasswordLogin()

	route.GET("/check", routes.Bind(c.Check))

	route.GET("/authoring", routes.Bind(c.Authorized))

	route.GET("/authorized", routes.Bind(c.Authorized))
}
