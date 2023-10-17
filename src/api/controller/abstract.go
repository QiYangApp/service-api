package controller

import "service-api/src/core/inject"

type AbstractControllerMethod interface {
}

// AbstractController
// @RootPath(path="/api")
type AbstractController struct {
	AbstractControllerMethod
}

type InjectController struct {
	AbstractController
	inject.Controller
}

func (*InjectController) PostConstruct() {

}
