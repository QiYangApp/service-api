package response

import (
	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Response[T interface{}] struct {
	Context   *gin.Context      `json:"-"`
	Type      ResponseTypeEnum  `json:"-"`
	RequestId uuid.UUID         `json:"request_id"`
	Code      int               `json:"code"`
	State     ResponseStateEnum `json:"state"`
	Message   string            `json:"message"`
	Timestamp string            `json:"timestamp"`
	Data      T                 `json:"data,omitempty"`
	List      []T               `json:"list,omitempty"`
	Page      int               `json:"page,omitempty"`
	Total     int               `json:"total,omitempty"`
}

func (r *Response[T]) SetPage(page, total int) *Response[T] {
	r.Page = page
	r.Total = total

	return r
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
	return &Response[T]{
		Code:      r.GetCode(),
		State:     r.GetState(),
		Message:   r.GetMessage(),
		Data:      r.Data,
		Timestamp: r.Timestamp,
		RequestId: r.RequestId,
	}
}

func (r *Response[T]) ToJson() *gin.Context {
	r.Context.SecureJSON(r.GetCode(), r.ToStruct())

	return r.Context
}

func (r *Response[T]) ToStream(data []byte) *gin.Context {
	r.Context.File(string(data))

	return r.Context
}
