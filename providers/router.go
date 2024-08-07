package providers

import (
	"frame/cmd"
	"service-api/router"
)

func RouterRegister() {
	router.Register(cmd.WebCli())
}
