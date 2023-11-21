// ⚠️⛔ Auto generate code by gin framework, Do not edit!!!
// All controller information and Api information for the current project is recorded here

package main

import (
	pack_71fc781208c148cecd8c184360762152 "service-api/src/api/controller"
	pack_850f5fd4b0bda50703b2d689c801ef40 "service-api/src/api/controller/authorize"
	"service-api/src/core/inject"
)

func init() {
	inject.Register(&pack_71fc781208c148cecd8c184360762152.InjectController{}, &pack_850f5fd4b0bda50703b2d689c801ef40.PasswordLoginController{}, &pack_850f5fd4b0bda50703b2d689c801ef40.PasswordRegisterController{})
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
			ApiPath:        "/api/authorize/password/authorized",
			PackMethodName: "PasswordLoginController",
			PackName:       "authorize",
			PackPath:       "service-api/src/api/controller/authorize",
		},
	}
}
