package providers

import (
	"frame/cmd"
	"service-api/internal/api/route"
)

func RouterRegister() {
	route.Register(cmd.WebCli())
}
