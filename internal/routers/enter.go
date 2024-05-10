package routers

import (
	"frame/cmd"
	"frame/modules/router"
	"service-api/internal/app/api/middlewares"

	"github.com/gin-gonic/gin"
)

var RouterGroup = []router.Router{
	&SwaggerRouter{},
	&CaptchaRouter{},
	&AuthRouter{},
	&TestRouter{},
}

func Register(app *cmd.WebServer) {
	r := app.Engine.Group("api")

	var funcs []gin.HandlerFunc = []gin.HandlerFunc{middlewares.Auth()}
	for _, middleware := range app.Middlewares {
		funcs = append(funcs, middleware())
	}

	r.Use(funcs...)

	for _, g := range RouterGroup {
		g.Handle(r)
	}
}
