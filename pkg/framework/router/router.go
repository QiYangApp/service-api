package router

import (
	"fmt"
	"frame/modules/log"
	"framework/response"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Handle(r *gin.RouterGroup)
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
				var d map[string]string
				if p, ok := tmp.Interface().(Validator); ok {
					d = GetErrorMsg(c, p, err)
				}
				response.RFail(
					c,
					d,
					http.StatusBadRequest,
					fmt.Sprintf("params error"),
				)

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
	if err := c.ShouldBind(v); err != nil {
		log.Client.Sugar().Warnf("route handle fun error, method params bind, error method: %v", v)
		return err
	}

	if err := c.ShouldBindUri(v); err != nil {
		log.Client.Sugar().Warnf("route handle fun error, uri method params bind, error method: %v", v)
		return err
	}

	return nil
}
