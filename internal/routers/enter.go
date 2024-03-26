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

	privateRoute := r.Group("")
	privateRoute.Use(middlewares.Auth())

	publicRoute := r.Group("")

	for _, r := range RouterGroup {

		r.Handle(privateRoute, publicRoute)

	}
}
