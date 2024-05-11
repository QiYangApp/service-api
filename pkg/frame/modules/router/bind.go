package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MethodType string

func (m MethodType) ToString() string {
	return string(m)
}

const (
	ANY     MethodType = "ANY"
	GET     MethodType = http.MethodGet
	POST    MethodType = http.MethodPost
	HEAD    MethodType = http.MethodHead
	PUT     MethodType = http.MethodPut
	PATCH   MethodType = http.MethodPatch
	DELETE  MethodType = http.MethodDelete
	CONNECT MethodType = http.MethodConnect
	OPTIONS MethodType = http.MethodOptions
	TRACE   MethodType = http.MethodTrace
)

func AnyBind(c *gin.RouterGroup, path string, fun any) gin.IRoutes {
	return AnyMethodBind(c, ANY, path, fun)
}

func GetBind(c *gin.RouterGroup, path string, fun any) gin.IRoutes {
	return AnyMethodBind(c, GET, path, fun)
}

func PostBind(c *gin.RouterGroup, path string, fun any) gin.IRoutes {
	return AnyMethodBind(c, POST, path, fun)
}

func HeadBind(c *gin.RouterGroup, path string, fun any) gin.IRoutes {
	return AnyMethodBind(c, HEAD, path, fun)
}

func PutBind(c *gin.RouterGroup, path string, fun any) gin.IRoutes {
	return AnyMethodBind(c, PUT, path, fun)
}

func PatchBind(c *gin.RouterGroup, path string, fun any) gin.IRoutes {
	return AnyMethodBind(c, PATCH, path, fun)
}

func DeleteBind(c *gin.RouterGroup, path string, fun any) gin.IRoutes {
	return AnyMethodBind(c, DELETE, path, fun)
}

func AnyMethodBind(c *gin.RouterGroup, method MethodType, path string, fun any) gin.IRoutes {
	bindFun := bind(fun)

	switch method {
	case ANY:
		return c.Any(path, bindFun)
	case GET:
		return c.GET(path, bindFun)
	case POST:
		return c.POST(path, bindFun)
	case HEAD:
		return c.HEAD(path, bindFun)
	case PATCH:
		return c.PATCH(path, bindFun)
	case PUT:
		return c.PUT(path, bindFun)
	case DELETE:
		return c.DELETE(path, bindFun)
	case OPTIONS:
		return c.OPTIONS(path, bindFun)
	}

	return nil
}
