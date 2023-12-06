// ⚠️⛔ Auto generate code by gin framework, Do not edit!!!
// All controller information and Api information for the current project is recorded here

package main

import (
	pack_71fc781208c148cecd8c184360762152 "service-api/src/api/controller"
	pack_f9cd26517a85a763fe7acac8691e4de8 "service-api/src/api/controller/authorize/password"
	"service-api/src/core/inject"
)

func init() {
	inject.Register(&pack_71fc781208c148cecd8c184360762152.InjectController{}, &pack_f9cd26517a85a763fe7acac8691e4de8.LoginController{}, &pack_f9cd26517a85a763fe7acac8691e4de8.RegisterController{})
	inject.DI = map[string]*inject.MethodInfo{
		"3f855e20f2fea9e7cd8932bee26b60a6": {

			Annotations:    map[string]string{},
			ApiMethodName:  "GET",
			ApiPath:        "/api/authorize/password/register/authorizing",
			PackMethodName: "RegisterController",
			PackName:       "password",
			PackPath:       "service-api/src/api/controller/authorize/password",
		},
		"556902bab00c142a01230816f883559a": {

			Annotations:    map[string]string{},
			ApiMethodName:  "GET",
			ApiPath:        "/api/authorize/password/login/authorizing",
			PackMethodName: "LoginController",
			PackName:       "password",
			PackPath:       "service-api/src/api/controller/authorize/password",
		},
		"7ff40593b4b22ed5dba9052c470c53af": {

			Annotations:    map[string]string{},
			ApiMethodName:  "GET",
			ApiPath:        "/api/authorize/password/login/authorized",
			PackMethodName: "LoginController",
			PackName:       "password",
			PackPath:       "service-api/src/api/controller/authorize/password",
		},
		"d5fa1995f02dd382744519e85df30b62": {

			Annotations:    map[string]string{},
			ApiMethodName:  "GET",
			ApiPath:        "/api/authorize/password/register/authorized",
			PackMethodName: "RegisterController",
			PackName:       "password",
			PackPath:       "service-api/src/api/controller/authorize/password",
		},
		"e5c48c31f71c60ebe119cafdb0a627db": {

			Annotations:    map[string]string{},
			ApiMethodName:  "GET",
			ApiPath:        "/api/authorize/password/register/check",
			PackMethodName: "RegisterController",
			PackName:       "password",
			PackPath:       "service-api/src/api/controller/authorize/password",
		},
		"ff1bd80654ed43f7fcc68f1ae721283d": {

			Annotations:    map[string]string{},
			ApiMethodName:  "GET",
			ApiPath:        "/api/authorize/password/login/check",
			PackMethodName: "LoginController",
			PackName:       "password",
			PackPath:       "service-api/src/api/controller/authorize/password",
		},
	}
}
