package providers

import (
	"frame/cmd"
	"service-api/internal/routers"
)

func RouterRegister() {
	routers.Register(cmd.WebCli())
}
