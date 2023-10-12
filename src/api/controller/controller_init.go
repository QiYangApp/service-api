// ⚠️⛔ Auto generate code by gin-plus framework, Do not edit!!!
// All controller information and Api information for the current project is recorded here

package controller

import (
	"service-api/src/core/inject"
)

func init() {
	inject.Register(&AbstractController{})
	inject.Apis = map[string]*inject.MethodInfo{"TestController/Test": {

		Annotations: map[string]string{},
		ApiPath:     "test",
		Method:      "GET",
	}}
}
