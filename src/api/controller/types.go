package controller

import "app/router"

type AbstractController struct {
}

// InjectController
// @RootPath(path="/api")
type InjectController struct {
	AbstractController
	router.Inject
}
