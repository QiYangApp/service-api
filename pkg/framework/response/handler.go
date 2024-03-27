package response

import (
	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Response struct {
	Context   *gin.Context `json:"-"`
	Type      TypeEnum     `json:"-"`
	RequestId uuid.UUID    `json:"request_id"`
	Code      int          `json:"code"`
	State     StateEnum    `json:"state"`
	Message   string       `json:"message"`
	Timestamp string       `json:"timestamp"`
	Data      any          `json:"data,omitempty"`
}

func (r *Response) SetData(data any) *Response {
	r.Data = data

	return r
}

func (r *Response) SetType(t TypeEnum) *Response {
	r.Type = t

	return r
}

func (r *Response) GetType() TypeEnum {
	return r.Type
}

func (r *Response) SetCode(code int) *Response {
	r.Code = code

	return r
}

func (r *Response) GetCode() int {
	if r.Code != 0 {
		return r.Code
	}

	return http.StatusOK
}

func (r *Response) SetMessage(message string) *Response {
	r.Message = message

	return r
}

func (r *Response) GetMessage() string {
	if r.Message == "" {
		return ""
	}

	msg := i18n.MustGetMessage(r.Context, r.Message)
	if msg == "" {
		return r.Message
	}

	return msg
}

func (r *Response) SetState(state StateEnum) *Response {
	r.State = state

	return r
}

func (r *Response) GetState() StateEnum {
	if r.State == "" {
		return Success
	}

	return r.State
}

func (r *Response) RContext() *gin.Context {
	return r.Context
}

func (r *Response) ToSelf() *Response {
	return r
}

func (r *Response) ToStruct() *Response {
	return &Response{
		Code:      r.GetCode(),
		State:     r.GetState(),
		Message:   r.GetMessage(),
		Data:      r.Data,
		Timestamp: r.Timestamp,
		RequestId: r.RequestId,
	}
}

func (r *Response) Output() *Response {

	switch r.Type {
	default:
	case JSON:
		toJson(r)
		break

	}

	return r
}

func toJson(r *Response) {
	r.Context.JSON(r.GetCode(), r.ToStruct())
	r.Context.Abort()
}
