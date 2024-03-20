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
}

func Register(app *cmd.WebServer) {
	r := app.Engine.Group("api")

	var funcs []gin.HandlerFunc
	for _, middleware := range app.Middlewares {
		funcs = append(funcs, middleware())
	}

	r.Use(funcs...)

	private := r.Group("")
	private.Use(middlewares.Auth())

	public := r.Group("")

	for _, r := range RouterGroup {
		if r.IsPrivate() {
			r.Handle(private)
		} else {
			r.Handle(public)
		}
	}
}
