package router

import (
	"fmt"
	"framework/exceptions"
	"framework/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Router interface {
	Handle(r *gin.RouterGroup)
	IsPrivate() bool
}

type BindFunc[T any] func(c *gin.Context, req T) *gin.Context

func Bind[T any](fun BindFunc[T]) gin.HandlerFunc {

	methodValueOf := reflect.ValueOf(fun)
	methodType := methodValueOf.Type()
	fmt.Println(methodValueOf.String(), methodType.String())

	return func(c *gin.Context) {
		var tmp reflect.Value
		var reqs []reflect.Value

		for i := 1; i < methodType.NumIn(); i++ {
			reqType := methodType.In(i)

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

		methodValueOf.Call(append([]reflect.Value{reflect.ValueOf(c)}, reqs...))
		if !c.IsAborted() {
			c.Abort()
		}
	}
}
