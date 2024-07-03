package resp

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
	"time"
)

func New(c *gin.Context) *Response {
	return &Response{
		Context:   c,
		Timestamp: time.Now().Format("20060102150405.0000"),
		RequestId: uuid.New(),
	}
}

func R[T interface{}](
	c *gin.Context,
	data T,
	code int,
	t TypeEnum,
	m string,
	s StateEnum,
) *Response {
	return New(c).SetType(t).SetState(s).SetMessage(&i18n.Message{ID: m}).SetCode(code).SetData(data)
}

func Success(c *gin.Context, data any) *Response {
	return New(c).SetState(SuccessState).SetCode(http.StatusOK).SetType(JSON).SetMessage(SuccessMessage).SetData(data)
}

func SuccessWithMsg(c *gin.Context, data any, msg string) *Response {
	return New(c).SetState(SuccessState).SetCode(http.StatusOK).SetType(JSON).SetMessage(&i18n.Message{ID: msg}).SetData(data)
}

func Fail(c *gin.Context, data any, code int, msg string) *Response {
	return New(c).SetState(FailState).SetType(JSON).SetCode(code).SetMessage(&i18n.Message{ID: msg}).SetData(data)
}

func Error(
	c *gin.Context,
	err error,
	code int,
	data any,
) *Response {
	return New(c).SetType(JSON).SetState(ErrorState).SetMessage(&i18n.Message{ID: err.Error()}).SetCode(code).SetData(data)
}

func Jump(c *gin.Context, data any, msg string) *Response {
	return New(c).SetType(JSON).SetState(ErrorState).SetMessage(&i18n.Message{ID: msg}).SetCode(http.StatusSeeOther).SetData(data)
}
