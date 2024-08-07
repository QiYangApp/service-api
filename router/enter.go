package router

import (
	"frame/cmd"
	"frame/modules/router"
	middlewares2 "service-api/internal/net/middlewares"
	"service-api/router/doc"
	v12 "service-api/router/group/v1"

	"github.com/gin-gonic/gin"
)

var RouterGroup = []router.Router{
	&doc.SwaggerRouter{},
	&v12.CaptchaRouter{},
	&v12.AuthRouter{},
}

func Register(app *cmd.WebServer) {
	r := app.Engine.Group("api")

	var funcs = []gin.HandlerFunc{
		middlewares2.Session(),
		middlewares2.Auth(),
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
