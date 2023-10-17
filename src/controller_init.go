// ⚠️⛔ Auto generate code by gin-plus framework, Do not edit!!!
// All controller information and Api information for the current project is recorded here

package main

import (
	controller "service-api/src/api/controller"
	authorize "service-api/src/api/controller/authorize"
	"service-api/src/core/inject"
)

func init() {
	inject.Register(&authorize.PasswordLoginController{}, &controller.TestController{}, &controller.InjectController{})
	inject.Apis = map[string]*inject.MethodInfo{
		"service-api/src/api/controller-TestController-Test": {

			Annotations: map[string]string{},
			ApiPath:     "test",
			MethodName:  "GET",
			PackName:    "controller",
			PackPath:    "service-api/src/api/controller",
			ServiceName: "TestController",
		},
		"service-api/src/api/controller/authorize-PasswordLoginController-Check": {

			Annotations: map[string]string{},
			ApiPath:     "check",
			MethodName:  "GET",
			PackName:    "authorize",
			PackPath:    "service-api/src/api/controller/authorize",
			ServiceName: "PasswordLoginController",
		},
	}
}
