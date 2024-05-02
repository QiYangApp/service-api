package conf

import "github.com/gin-gonic/gin"

func IsRunDebug() bool {
	return RunMode() == gin.DebugMode
}

func IsRunTest() bool {
	return RunMode() == gin.TestMode
}

func IsRunRelease() bool {
	return RunMode() == gin.ReleaseMode
}

func IsDebug() bool {
	return Client().GetBool("debug")
}

func RunMode() string {
	return Client().GetString("run_mode")
}
