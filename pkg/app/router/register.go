package router

// https://github.com/xxjwxc/ginrpc/blob/master/common.go#L57
// https://github.com/archine/gin-plus/blob/v2/ast/mvc2/main.go

import (
	"crypto/md5"
	"fmt"
	"github.com/archine/ioc"
	"github.com/gin-gonic/gin"
	"reflect"
)

type Register struct {
	Scan Scan
	[]MethodInfo
}

func (r *Register) Register(e *gin.Engine, autowired bool) {

	annotationCache = make(map[string]Annotations)
	for _, di := range r.Scan.Apis {
		if autowired {
			ioc.Inject(di)
		}
		di.PostConstruct()
		diTypeOf := reflect.TypeOf(di)
		diProxy := reflect.ValueOf(di)
		for i := 0; i < diTypeOf.NumMethod(); i++ {
			methodProxy := diTypeOf.Method(i)

			key := diTypeOf.Elem().PkgPath() + "-" + diTypeOf.Elem().Name() + "-" + methodProxy.Name
			key = fmt.Sprintf("%x", md5.Sum([]byte(key)))

			if info, ok := DI[key]; ok {
				if method, ok := diTypeOf.MethodByName(methodProxy.Name); ok {
					RegisterRoute(e, info.ApiMethodName, info.ApiPath, method, diTypeOf, diProxy)
				}

			}
		}

		if len(diCache) == 1 {
			diCache = nil
			return
		}

		diCache = diCache[1:]
	}

	DI = nil
}
