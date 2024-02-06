// ⚠️⛔ Auto generate code by gin framework, Do not edit!!!
// All controller information and Api information for the current project is recorded here

package main

import (
	"app/router"
	pack_71fc781208c148cecd8c184360762152 "service-api/src/api/controller"
)

func init() {
	router.Register(&pack_71fc781208c148cecd8c184360762152.LoginController{}, &pack_71fc781208c148cecd8c184360762152.InjectController{})
	router.Apis = map[string]*router.MethodInfo{
		"054850cbb8f2d5f96de13218da515d5c": {

			Annotations:    map[string]string{},
			ApiMethodName:  "GET",
			ApiPath:        "/api/authorize/password/login/check",
			PackMethodName: "LoginController",
			PackName:       "controller",
			PackPath:       "service-api/src/api/controller",
		},
		"1d4be837d5aa8509780b72e462832bc7": {

			Annotations:    map[string]string{},
			ApiMethodName:  "POST",
			ApiPath:        "/api/authorize/password/login/authorized",
			PackMethodName: "LoginController",
			PackName:       "controller",
			PackPath:       "service-api/src/api/controller",
		},
		"f19890d9a8f7d55974e16547b8a922c4": {

			Annotations:    map[string]string{},
			ApiMethodName:  "GET",
			ApiPath:        "/api/authorize/password/login/authorizing",
			PackMethodName: "LoginController",
			PackName:       "controller",
			PackPath:       "service-api/src/api/controller",
		},
	}
}
