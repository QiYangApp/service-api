package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func NewResponse[T interface{}](c *gin.Context) *Response[T] {
	return &Response[T]{
		Context:   c,
		Timestamp: time.Now(),
	}
}

func R[T interface{}](
	c *gin.Context,
	data T,
	code int,
	t ResponseTypeEnum,
	m string,
	s ResponseStateEnum,
) *Response[T] {
	return NewResponse[T](c).SetType(t).SetState(s).SetMessage(m).SetCode(code).SetData(data)
}

func RSuccess[T interface{}](c *gin.Context, data T) *Response[T] {
	return NewResponse[T](c).SetState(Success).SetCode(http.StatusOK).SetType(JSON).SetMessage("success").SetData(data)
}

func RError[T interface{}](
	c *gin.Context,
	data T,
	code int,
	m string,
	s ResponseStateEnum,
) *Response[T] {
	return NewResponse[T](c).SetType(JSON).SetState(s).SetMessage(m).SetCode(code).SetData(data)
}

func RImage[T interface{}](c *gin.Context, data []byte) *gin.Context {
	return NewResponse[T](c).SetType(IMAGE).ToStream(data)
}

func RFail[T interface{}](c *gin.Context, data T, code int) *Response[T] {
	return NewResponse[T](c).SetType(IMAGE).SetCode(code).SetMessage("success").SetData(data)
}
