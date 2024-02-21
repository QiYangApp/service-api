package router

import (
	"crypto/md5"
	"fmt"
	"framework/exceptions"
	"framework/log"
	"framework/response"
	"github.com/archine/ioc"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

// https://github.com/xxjwxc/ginrpc/blob/master/common.go#L57
// https://github.com/archine/gin-plus/blob/v2/ast/mvc2/main.go

// Apis  Apis Project global API
var Apis map[string]*MethodInfo

var Cache []inject

// Register controllers
func Register(containers ...inject) {
	Cache = append(Cache, containers...)
}

// Apply all apis to the gin engine
// @param e: gin.Engine
// @param autowired: whether enable autowired properties
func Apply(e *gin.Engine, autowired bool) {
	for _, di := range Cache {
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

			if info, ok := Apis[key]; ok {
				if method, ok := diTypeOf.MethodByName(methodProxy.Name); ok {
					RegisterRoute(e, info.ApiMethodName, info.ApiPath, method, diTypeOf, diProxy)
				}

			}
		}
	}

	Apis = nil
}

func RegisterRoute(e *gin.Engine, method string, path string, fn reflect.Method, typeOf reflect.Type, valueOf reflect.Value) {
	RouteBind(e, method, path, fn, typeOf, valueOf)
}

func RouteBind(e *gin.Engine, method string, path string, fn reflect.Method, typeOf reflect.Type, valueOf reflect.Value) {
	callFn := RouteHandle(fn, valueOf, typeOf)

	switch strings.ToUpper(method) {
	case "ANY":
		e.Any(path, callFn)
		break
	case "POST":
		e.POST(path, callFn)
		break
	case "GET":
		e.GET(path, callFn)
		break
	case "PATCH":
		e.PATCH(path, callFn)
		break
	case "PUT":
		e.PUT(path, callFn)
		break
	case "DELETE":
		e.DELETE(path, callFn)
		break
	case "OPTIONS":
		e.OPTIONS(path, callFn)
		break
	default:
		e.Handle(method, path, callFn)
	}
}

func RouteHandle(fn reflect.Method, valueOf reflect.Value, typeOf reflect.Type) gin.HandlerFunc {
	if fn.Func.Type().NumIn() <= 1 {
		log.Client().Sugar().Fatal("route handle fun error, method name: %v", fn.Name)
	}

	reqType := fn.Func.Type().In(1)
	if reqType != reflect.TypeOf(&gin.Context{}) {
		log.Client().Sugar().Fatal("route handle fun error, method first params require *gin.Context{},  error method: %v", fn.Name)
	}

	var reqLen = fn.Func.Type().NumIn()
	return func(c *gin.Context) {
		var tmp reflect.Value
		var reqs []reflect.Value

		for i := 2; i < reqLen; i++ {
			reqType := fn.Func.Type().In(i)

			if reqType.Kind() != reflect.Ptr {
				tmp = reflect.New(reqType)
			} else {
				tmp = reflect.New(reqType.Elem())
			}

			if err := unmarshal(c, tmp.Interface()); err != nil { // Return error message.返回错误信息
				response.RError(
					c,
					exceptions.New(err.Error()),
					http.StatusBadRequest,
					fmt.Sprintf("params error"),
				).ToJson().Abort()
				return
			}

			if reqType.Kind() == reflect.Ptr {
				reqs = append(reqs, tmp)
			} else {
				reqs = append(reqs, tmp.Elem())
			}
		}

		fn.Func.Call(append([]reflect.Value{valueOf, reflect.ValueOf(c)}, reqs...))
		if !c.IsAborted() {
			c.Abort()
		}
	}

}

func unmarshal(c *gin.Context, v interface{}) error {
	err := c.ShouldBind(v)
	if err != nil {
		log.Client().Sugar().Error("route handle fun error, method params bind, error method: %v", v)
	}

	err = c.ShouldBindUri(v)
	if err != nil {
		log.Client().Sugar().Error("route handle fun error, uri method params bind, error method: %v", v)
	}

	return err
}

// MethodInterceptor API method interceptor
// You can do logical processing before and after method calls
type MethodInterceptor interface {
	// Predicate true means intercept.
	Predicate(ctx *gin.Context) bool

	// PreHandle triggered before method invocation.
	// If you want to abort the current request, just call abort() and response inside the method
	PreHandle(ctx *gin.Context)

	// PostHandle triggered after method invocation.
	// If you want to abort the current request, just call abort() and response inside the method
	PostHandle(ctx *gin.Context)
}
