package response

import (
	"github.com/gin-gonic/gin"
	"time"
)

func NewResponse[T interface{}](c *gin.Context) *Response[T] {
	return &Response[T]{
		Context:   c,
		Timestamp: time.Now(),
	}
}

func Single[T interface{}](c *gin.Context, data T) *Response[T] {
	return NewResponse[T](c).Single(data)
}

func SingleSuccess[T interface{}](c *gin.Context, data T) *Response[T] {
	return NewResponse[T](c).SetCode(200).SetMessage("STATE.SUCCESS").Single(data)
}

func SingleCustom[T interface{}](c *gin.Context, data T, code int, message string) *Response[T] {
	return NewResponse[T](c).SetCode(code).SetMessage(message).Single(data)
}
