package ioc

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

//https://github.com/archine/gin-plus/blob/v2/mvc/controller.go

// controller Top-level interface used to declare a structure as a controller.
type abstractIoc interface {
	// Construct Triggered after dependency injection is completed. You can continue to decorate the controller here.
	Construct()
}

type Ioc struct{}

func (c *Ioc) Construct() {}

// Annotations the annotation of Api method
type Annotations map[string]string

// Global controller cache
var caches []Ioc

// Annotations of each API
var annotationCache map[string]Annotations

// Register controllers
func Register(ioc ...Ioc) {
	caches = append(caches, ioc...)
}

// IsIoc IsController Determine whether it is controller
func IsIoc(v interface{}) bool {
	ct := reflect.TypeOf(v)
	if ct.Kind() != reflect.Ptr {
		return false
	}
	return ct.Implements(reflect.TypeOf((*abstractIoc)(nil)).Elem())
}

func Inject(ioc Ioc) {
	if IsIoc(ioc) == false {

	}
}

// Apply all apis to the gin engine
// @param e: gin.Engine
// @param autowired: whether enable autowired properties
//func Apply(e *gin.Engine, autowired bool) {
//	//if ast.Apis == nil {
//	//	for _, c := range caches {
//	//		if autowired {
//	//			ioc.Inject(c)
//	//		}
//	//	}
//	//	return
//	//}
//	ginProxy := reflect.ValueOf(e)
//	annotationCache = make(map[string]Annotations)
//	for _, c := range caches {
//		if autowired {
//			ioc.Inject(c)
//		}
//		c.Construct()
//
//		cTypeOf := reflect.TypeOf(c)
//		cProxy := reflect.ValueOf(c)
//		for i := 0; i < cTypeOf.NumMethod(); i++ {
//			methodProxy := cTypeOf.Method(i)
//			methodFullName := controllerTypeOf.Elem().Name() + "/" + methodProxy.Name
//			if info, ok := ast.Apis[methodFullName]; ok {
//				ginMethod := ginProxy.MethodByName(info.Method)
//				args := []reflect.Value{reflect.ValueOf(info.ApiPath)}
//				args = append(args, controllerProxy.MethodByName(methodProxy.Name))
//				ginMethod.Call(args)
//				annotationCache[info.ApiPath] = info.Annotations
//			}
//		}
//		if len(controllerCache) == 1 {
//			controllerCache = nil
//			return
//		}
//		controllerCache = controllerCache[1:]
//	}
//	ast.Apis = nil
//}

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
