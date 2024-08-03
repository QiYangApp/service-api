package route

import (
	"frame/cmd"
	"frame/modules/router"
	"service-api/internal/net/middlewares"
	"service-api/internal/net/route/doc"
	v1 "service-api/internal/net/route/group/v1"

	"github.com/gin-gonic/gin"
)

var RouterGroup = []router.Router{
	&doc.SwaggerRouter{},
	&v1.CaptchaRouter{},
	&v1.AuthRouter{},
}

func Register(app *cmd.WebServer) {
	r := app.Engine.Group("api")

	var funcs = []gin.HandlerFunc{
		middlewares.Session(),
		middlewares.Auth(),
	}

	for _, middleware := range app.Middlewares {
		funcs = append(funcs, middleware())
	}

	r.Use(funcs...)

	Routes(r)
}

func Routes(r *gin.RouterGroup) {
	for _, g := range RouterGroup {
		g.Handle(r)
	}
}
