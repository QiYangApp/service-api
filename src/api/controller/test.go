package controller

import "service-api/src/core/inject"

type TestController struct {
	inject.Controller
}

// Test
// @GET(path="test")
func (t *TestController) Test() string {
	return ""
}
