package inject

// https://github.com/xxjwxc/ginrpc/blob/master/common.go#L57
// https://github.com/archine/gin-plus/blob/v2/ast/mvc2/main.go

import (
	"crypto/md5"
	"github.com/archine/ioc"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
)

// controller Top-level interface used to declare a structure as a controller.
type abstractController interface {
	// PostConstruct Triggered after dependency injection is completed. You can continue to decorate the controller here.
	PostConstruct()
}

/*
Controller Default abstract controller implementation.

	Simply integrate the default controller into your structure.

Example:

	type YourController struct {
	   mvc.Controller
	}

	// Hello
	// @GET(path="/hello") this is api method
	func (y *YourController) Hello(ctx *gin.Context) {
	   resp.Json(ctx, "Hello World")
	}

	// Access the API
	curl http://localhost:4006/hello
*/
type Controller struct{}

func (c *Controller) PostConstruct() {}

// Annotations the annotation of Api method
type Annotations map[string]string

// Global controller cache
var controllerCache []abstractController

// Annotations of each API
var annotationCache map[string]Annotations

// Register controllers
func Register(controller ...abstractController) {
	controllerCache = append(controllerCache, controller...)
}

// IsController Determine whether it is controller
func IsController(v interface{}) bool {
	ct := reflect.TypeOf(v)
	if ct.Kind() != reflect.Ptr {
		return false
	}
	return ct.Implements(reflect.TypeOf((*abstractController)(nil)).Elem())
}

// Apply all apis to the gin engine
// @param e: gin.Engine
// @param autowired: whether enable autowired properties
func Apply(e *gin.Engine, autowired bool) {
	if Apis == nil {
		for _, controller := range controllerCache {
			if autowired {
				ioc.Inject(controller)
			}
		}
		return
	}

	haser := md5.New()
	ginProxy := reflect.ValueOf(e)
	annotationCache = make(map[string]Annotations)
	for _, controller := range controllerCache {
		if autowired {
			ioc.Inject(controller)
		}
		controller.PostConstruct()
		controllerTypeOf := reflect.TypeOf(controller)
		controllerProxy := reflect.ValueOf(controller)
		for i := 0; i < controllerTypeOf.NumMethod(); i++ {
			methodProxy := controllerTypeOf.Method(i)

			key := controllerTypeOf.Elem().PkgPath() + "-" + controllerTypeOf.Elem().Name() + "-" + methodProxy.Name
			haser.Write([]byte(key))
			//key = hex.EncodeToString(haser.Sum(nil))

			if info, ok := Apis[key]; ok {
				//ginMethod := ginProxy.MethodByName(info.MethodName)
				//args := []reflect.Value{reflect.ValueOf(info.ApiPath)}
				//args = append(args, controllerProxy.MethodByName(methodProxy.Name))
				//ginMethod.Call(args)
				//annotationCache[info.ApiPath] = info.Annotations
				//
				//fn, ok := controllerTypeOf.MethodByName(methodProxy.Name)
				//if !ok {
				//	continue
				//}
				//
				//e.Match([]string{info.MethodName}, info.ApiPath, fn)

				if method, ok := controllerTypeOf.MethodByName(methodProxy.Name); ok {
					RegisterRoute(e, info.MethodName, info.ApiPath, method, controllerTypeOf, controllerProxy)
				}

			}
		}
		if len(controllerCache) == 1 {
			controllerCache = nil
			return
		}
		controllerCache = controllerCache[1:]
	}

	Apis = nil
}

func RegisterRoute(e *gin.Engine, method string, path string, fn reflect.Method, TypeOf reflect.Type, ValueOf reflect.Value) {

}

func RouteBind(e *gin.Engine, method string, path string, fn reflect.Value, TypeOf reflect.Type, ValueOf reflect.Value) {

}

func RouteHandle(fn reflect.Method, valueOf reflect.Value, typeOf reflect.Type) gin.HandlerFunc {
	if valueOf.Type().NumIn() <= 1 {
		log.Fatalf("route handle fun error, method name: %v", fn.Name)
	}

	t := valueOf.Type().In(1)

	defFun := func(c *gin.Context) interface{} { return c }
	if t == defFun {

	}
}

// GetAnnotation Gets the specified annotation
// Returns the value of this annotation, when the has is false mine this val is empty
func GetAnnotation(ctx *gin.Context, annotationName string) (val string, has bool) {
	anno, has := annotationCache[ctx.FullPath()]
	if !has || len(anno) == 0 {
		return "", false
	}
	val, has = anno[annotationName]
	return
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
