package controller

import (
	"service-api/src/core/inject"
)

type AbstractMethod struct{}

type AbstractController struct {
	AbstractMethod
	inject.Controller
}
