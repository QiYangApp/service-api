package routers

import (
	"framework/cmd"
	"framework/router"
	"github.com/gin-gonic/gin"
	"service-api/internal/app/api/middlewares"
)

var RouterGroup = []router.Router{
	&SwaggerRouter{},
	&CaptchaRouter{},
	&AuthRouter{},
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
