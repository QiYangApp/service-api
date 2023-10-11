package helpers

import (
	"github.com/archine/ioc"
	"github.com/gin-gonic/gin"
	"reflect"
)

//https://github.com/archine/gin-plus/blob/v2/mvc/controller.go

// controller Top-level interface used to declare a structure as a controller.
type abstractInject interface {
	// Construct Triggered after dependency injection is completed. You can continue to decorate the controller here.
	Construct()
}

type Inject struct{}

func (c *Inject) Construct() {}

// Global  cache
var caches = make(map[string]abstractInject)

// IsIoc IsController Determine whether it is controller
func IsIoc(v interface{}) bool {
	ct := reflect.TypeOf(v)
	if ct.Kind() != reflect.Ptr {
		return false
	}
	return ct.Implements(reflect.TypeOf((*abstractInject)(nil)).Elem())
}

func Autowired(inject abstractInject) {
	if IsIoc(inject) == false {
		return
	}

	typeof := reflect.TypeOf(inject)

	if caches[typeof.Elem().String()] != nil {
		inject = caches[typeof.Elem().String()]
	} else {
		ioc.Inject(inject)
		caches[typeof.Elem().String()] = inject
	}

	inject.Construct()
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
