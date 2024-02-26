package providers

import (
	"framework/cmd"
	"service-api/internal/routers"
)

type Router struct {
}

func (*Router) Register() {
	routers.Register(cmd.WebServerClient())
}
