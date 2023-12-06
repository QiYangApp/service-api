package response

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func NewResponse[T interface{}](c *gin.Context) *Response[T] {
	return &Response[T]{
		Context:   c,
		Timestamp: time.Now().Format("20060102150405.0000"),
		RequestId: uuid.New(),
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

func RError(
	c *gin.Context,
	err error,
	code int,
	data any,
) *Response[any] {
	return NewResponse[any](c).SetType(JSON).SetState(Error).SetMessage(err.Error()).SetCode(code).SetData(data)
}

func RImage[T interface{}](c *gin.Context, data []byte) *gin.Context {
	return NewResponse[T](c).SetType(IMAGE).ToStream(data)
}

func RFail[T interface{}](c *gin.Context, data T, code int, mes string) *Response[T] {
	return NewResponse[T](c).SetState(Fail).SetType(JSON).SetCode(code).SetMessage(mes).SetData(data)
}
