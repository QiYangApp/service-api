// ⚠️⛔ Auto generate code by gin framework, Do not edit!!!
// All controller information and Api information for the current project is recorded here

package main

import (
	"app/router"
	pack_71fc781208c148cecd8c184360762152 "service-api/src/api/controller"
	"service-api/src/api/controller/authorize"
)

func init() {
	router.Register(&authorize.LoginController{}, &pack_71fc781208c148cecd8c184360762152.InjectController{})
	router.Apis = map[string]*router.MethodInfo{
		"691e02a00fde316bb5e9f560404484d7": {

			Annotations:    map[string]string{},
			ApiMethodName:  "GET",
			ApiPath:        "/api/authorizing/login",
			PackMethodName: "LoginController",
			PackName:       "controller",
			PackPath:       "service-api/src/api/controller",
		},
		"a6d9cdb0e19ca62c6d2503f11c0c2daa": {

			Annotations:    map[string]string{},
			ApiMethodName:  "GET",
			ApiPath:        "/api/authorizing/register",
			PackMethodName: "LoginController",
			PackName:       "controller",
			PackPath:       "service-api/src/api/controller",
		},
	}
}
