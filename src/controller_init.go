// ⚠️⛔ Auto generate code by gin framework, Do not edit!!!
// All controller information and Api information for the current project is recorded here

package main

import (
	pack_71fc781208c148cecd8c184360762152 "service-api/src/api/controller"
	pack_44baa70fb0a0c7a5b3e85c805406f0bf "service-api/src/api/controller/authorize"
	"service-api/src/core/inject"
)

func init() {
	inject.Register(&pack_71fc781208c148cecd8c184360762152.InjectController{}, &pack_44baa70fb0a0c7a5b3e85c805406f0bf.PasswordLoginController{})
	inject.DI = map[string]*inject.MethodInfo{
		"5a95ba9c494eba401182399acc6c4d95": {

			Annotations:    map[string]string{},
			ApiMethodName:  "GET",
			ApiPath:        "/api/authorize/password/authorizing",
			PackMethodName: "PasswordLoginController",
			PackName:       "authorize",
			PackPath:       "service-api/src/api/controller/authorize",
		},
		"8da03d27bb2baa2ecbde4bff293d7668": {

			Annotations:    map[string]string{},
			ApiMethodName:  "GET",
			ApiPath:        "/api/authorize/password/check",
			PackMethodName: "PasswordLoginController",
			PackName:       "authorize",
			PackPath:       "service-api/src/api/controller/authorize",
		},
		"994daa151fc09281b473959f28371e38": {

			Annotations:    map[string]string{},
			ApiMethodName:  "GET",
			ApiPath:        "/api/authorize/password/Authorized",
			PackMethodName: "PasswordLoginController",
			PackName:       "authorize",
			PackPath:       "service-api/src/api/controller/authorize",
		},
	}
}
