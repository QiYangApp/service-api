package providers

import (
	"frame/cmd"
	"service-api/internal/net/route"
)

func RouterRegister() {
	route.Register(cmd.WebCli())
}
