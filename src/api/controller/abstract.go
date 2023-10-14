package controller

import (
	"service-api/src/core/inject"
)

type AbstractMethod struct{}

// AbstractController
// @BasePath(path="/api)
type AbstractController struct {
	AbstractMethod
	inject.Controller
}
