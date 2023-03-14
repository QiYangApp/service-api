package helpers

import (
	"net/http"
	"time"
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
	Single(data T) SingleResponse[T]
	// Pagination 分页
	Pagination(list []T) PaginationResponse[T]
	// Multiple 多个返回
	Multiple(data []T) SingleResponse[T]
}

type Response[T interface{}] struct {
	Code      int               `json:"code"`
	State     ResponseStateEnum `json:"state"`
	Message   string            `json:"message"`
	Timestamp time.Time         `json:"timestamp"`
}

type SingleResponse[T interface{}] struct {
	Response Response[T]
	Data     T   `json:"data,omitempty"`
	List     []T `json:"list,omitempty"`
}

type PaginationResponse[T interface{}] struct {
	Response Response[T]
	List     []T   `json:"list,omitempty"`
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
	if r.Message != "" {
		return r.Message
	}

	return ""
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

func (r *Response[T]) Single(data T) SingleResponse[T] {
	return SingleResponse[T]{
		Response: r.getDefaultResponse(),
		Data:     data,
	}
}

func (r *Response[T]) Pagination(list []T) PaginationResponse[T] {
	return PaginationResponse[T]{
		Response: r.getDefaultResponse(),
		List:     list,
	}
}

func (r *Response[T]) Multiple(list []T) SingleResponse[T] {
	return SingleResponse[T]{
		Response: r.getDefaultResponse(),
		List:     list,
	}
}

func (r *Response[T]) getDefaultResponse() Response[T] {
	return Response[T]{
		Message: r.GetMessage(),
		Code:    r.GetCode(),
		State:   r.GetState(),
	}
}

func NewResponse[T interface{}]() *Response[T] {
	return &Response[T]{
		Timestamp: time.Now(),
	}
}
