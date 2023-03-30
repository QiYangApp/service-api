package helpers

import (
	"net/http"
	"time"

	"github.com/gin-contrib/i18n"
)

type ResponseStateEnum string

const (
	Success ResponseStateEnum = "success"
	Error   ResponseStateEnum = "error"
	Warn    ResponseStateEnum = "warn"
	Fail    ResponseStateEnum = "fail"
)

type ResponseMethods[T interface{}] interface {
	// Single 单个返回
	Single(data T) Response[T]
	// Pagination 分页
	Pagination(list []T) Response[T]
	// Multiple 多个返回
	Multiple(data []T) Response[T]
}

type Response[T interface{}] struct {
	Code      int               `json:"code"`
	State     ResponseStateEnum `json:"state"`
	Message   string            `json:"message"`
	Timestamp time.Time         `json:"timestamp"`
	Data     T   `json:"data,omitempty"`
	List     []T `json:"list,omitempty"`
	Page     int32 `json:"page,omitempty"`
	PageSize int32 `json:"page_size,omitempty"`
	Total    int32 `json:"total,omitempty"`
	LastPage int32 `json:"last_page,omitempty"`
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

	return i18n.MustGetMessage(r.Message)
}

func (r *Response[T]) SetState(state ResponseStateEnum) *Response[T] {
	r.State = state

	return r
}

func (r *Response[T]) GetState() ResponseStateEnum {
	switch r.State {
	case Success:
	case Fail:
	case Warn:
	case Error:
		return r.State
	default:
		return Success
	}

	return Success
}

func (r *Response[T]) Single(data T) Response[T] {
	r.Data = data;

	return r.toStruct();
}

func (r *Response[T]) Pagination(list []T) Response[T] {
	r.List = list;

	return r.toStruct();
}

func (r *Response[T]) Multiple(list []T) Response[T] {
	r.List = list;

	return r.toStruct();
}

func (r *Response[T]) toStruct() Response[T] {
	resp := *r
	resp.Message = r.GetMessage()
	resp.Code = r.GetCode()
	resp.State = r.GetState()

	return resp
}

func NewResponse[T interface{}]() *Response[T] {
	return &Response[T]{
		Timestamp: time.Now(),
	}
}
