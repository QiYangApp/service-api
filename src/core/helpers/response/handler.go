package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/gin-contrib/i18n"
)

type Response[T interface{}] struct {
	Context   *gin.Context      `json:"-"`
	Code      int               `json:"code"`
	State     ResponseStateEnum `json:"state"`
	Message   string            `json:"message"`
	Timestamp time.Time         `json:"timestamp"`
	Data      T                 `json:"data,omitempty"`
	List      []T               `json:"list,omitempty"`
	Page      int32             `json:"page,omitempty"`
	PageSize  int32             `json:"page_size,omitempty"`
	Total     int32             `json:"total,omitempty"`
	LastPage  int32             `json:"last_page,omitempty"`
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

	return i18n.MustGetMessage(r.Context, r.Message)
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

func (r *Response[T]) Single(data T) *Response[T] {
	r.Data = data

	return r.toStruct()
}

func (r *Response[T]) Pagination(list []T) *Response[T] {
	r.List = list

	return r.toStruct()
}

func (r *Response[T]) Multiple(list []T) *Response[T] {
	r.List = list

	return r.toStruct()
}

func (r *Response[T]) toStruct() *Response[T] {
	r.Message = r.GetMessage()
	r.Code = r.GetCode()
	r.State = r.GetState()

	return r
}
