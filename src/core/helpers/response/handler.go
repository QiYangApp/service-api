package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/gin-contrib/i18n"
)

type Response[T interface{}] struct {
	Context   *gin.Context      `json:"-"`
	Type      ResponseTypeEnum  `json:"-"`
	Code      int               `json:"code"`
	State     ResponseStateEnum `json:"state"`
	Message   string            `json:"message"`
	Timestamp time.Time         `json:"timestamp"`
	Data      T                 `json:"data,omitempty"`
}

func (r *Response[T]) SetData(data T) *Response[T] {
	r.Data = data

	return r
}

func (r *Response[T]) SetType(t ResponseTypeEnum) *Response[T] {
	r.Type = t

	return r
}

func (r *Response[T]) GetType() ResponseTypeEnum {
	return r.Type
}

func (r *Response[T]) SetCode(code int) *Response[T] {
	r.Code = code

	return r
}

func (r *Response[T]) GetCode() int {
	if r.Code != 0 {
		return r.Code
	}

	return http.StatusOK
}

func (r *Response[T]) SetMessage(message string) *Response[T] {
	r.Message = message

	return r
}

func (r *Response[T]) GetMessage() string {
	if r.Message == "" {
		return ""
	}

	msg := i18n.MustGetMessage(r.Context, r.Message)
	if msg == "" {
		return r.Message
	}

	return msg
}

func (r *Response[T]) SetState(state ResponseStateEnum) *Response[T] {
	r.State = state

	return r
}

func (r *Response[T]) GetState() ResponseStateEnum {
	if r.State == "" {
		return Success
	}

	return r.State
}

func (r *Response[T]) RContext() *gin.Context {
	return r.Context
}

func (r *Response[T]) ToStruct() *Response[T] {
	r.Message = r.GetMessage()
	r.Code = r.GetCode()
	r.State = r.GetState()

	return r
}

func (r *Response[T]) ToJson() *gin.Context {
	r.Context.SecureJSON(r.GetCode(), r.ToStruct())

	return r.Context
}

func (r *Response[T]) ToStream(data []byte) *gin.Context {
	r.Context.File(string(data))

	return r.Context
}
