package main

import (
	"framework/router"
	"framework/utils"
)

func main() {
	scan := &router.Scan{
		Apis:       make(map[string]*router.MethodInfo),
		BasePaths:  make(map[string]string),
		Names:      make(map[string]router.Names),
		Controller: make([]router.Inject, 0),
	}

	scan.RecursionPkgDir(utils.Path.ControllerPath, "service-api/api/http/controller", "controller")

	router.GenerateApi(scan, utils.Path.RootPath)
}
