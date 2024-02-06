package providers

import (
	"app"
	"service-api/src/router"
)

type Router struct {
}

func (*Router) Register(app *app.App) {
	router.Entry(app.Engine)
}
