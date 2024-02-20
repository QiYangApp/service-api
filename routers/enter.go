package routers

import (
	"framework"
	"framework/router"
	"github.com/gin-gonic/gin"
	"service-api/api/http/middlewares"
)

var RouterGroup = []router.Router{
	&CaptchaRouter{},
}

func Register(app *framework.App) {
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
