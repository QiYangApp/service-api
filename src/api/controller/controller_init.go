// ⚠️⛔ Auto generate code by gin-plus framework, Do not edit!!!
// All controller information and Api information for the current project is recorded here

package controller

import "service-api/src/core/inject"

func init() {
	inject.Register()
	inject.Apis = map[string]*inject.MethodInfo{
		"2bf6b51395f94da1c7905daf27970cbc": {

			Annotations: map[string]string{},
			ApiPath:     "path=/api/authorize/password/check",
			Method:      "GET",
			PackName:    "authorize",
			PackPath:    "service-api/src/api/controller/authorize",
		},
		"a56c6e931fca8b4154cf7169d79c0c15": {

			Annotations: map[string]string{},
			ApiPath:     "test",
			Method:      "GET",
			PackName:    "controller",
			PackPath:    "service-api/src/api/controller",
		},
	}
}