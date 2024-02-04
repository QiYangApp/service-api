package controller

import "app/router"

type AbstractControllerMethod interface {
}

type AbstractController struct {
}

// InjectController
// @RootPath(path="/api")
type InjectController struct {
	AbstractController
	router.Inject
}
