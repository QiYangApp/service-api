package main

import (
	"app/helpers"
	"app/router"
)

func main() {
	scan := &router.Scan{
		Apis:       make(map[string]*router.MethodInfo),
		BasePaths:  make(map[string]string),
		Names:      make(map[string]router.Names),
		Controller: make([]router.Inject, 0),
	}

	scan.RecursionPkgDir(helpers.Path.ControllerPath, "service-api/src/api/controller", "controller")

	router.GenerateApi(scan, helpers.Path.RootPath)
}
