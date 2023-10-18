package controller

import "service-api/src/core/inject"

type AbstractControllerMethod interface {
}

type AbstractController struct {
}

// InjectController
// @RootPath(path="/api")
type InjectController struct {
	AbstractController
	inject.Inject
}
