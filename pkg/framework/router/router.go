package router

import (
	"fmt"
	"framework/log"
	"framework/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Router interface {
	Handle(private *gin.RouterGroup, public *gin.RouterGroup)
}

func Bind(fun any) gin.HandlerFunc {

	methodValueOf := reflect.ValueOf(fun)
	methodType := methodValueOf.Type()

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
				response.RFail(
					c,
					GetErrorMsg(c, tmp.Interface().(Validator), err),
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

func unmarshal(c *gin.Context, v interface{}) error {
	err := c.ShouldBind(v)
	if err != nil {
		log.Client.Sugar().Warnf("route handle fun error, method params bind, error method: %v", v)
	}

	err = c.ShouldBindUri(v)
	if err != nil {
		log.Client.Sugar().Warnf("route handle fun error, uri method params bind, error method: %v", v)
	}

	return err
}
