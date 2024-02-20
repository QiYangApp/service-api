package providers

import (
	"framework"
	"service-api/routers"
)

type Router struct {
}

func (*Router) Register(app *framework.App) {
	routers.Register(app)
}
