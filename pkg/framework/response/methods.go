package response

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	return New(c).SetType(t).SetState(s).SetMessage(m).SetCode(code).SetData(data)
}

func RSuccess(c *gin.Context, data any) {
	New(c).SetState(Success).SetCode(http.StatusOK).SetType(JSON).SetMessage("STATE.SUCCESS").SetData(data).Output()
}

func RSuccessWithMsg(c *gin.Context, data any, msg string) {
	New(c).SetState(Success).SetCode(http.StatusOK).SetType(JSON).SetMessage(msg).SetData(data).Output()
}

func RFail(c *gin.Context, data any, code int, mes string) {
	New(c).SetState(Fail).SetType(JSON).SetCode(code).SetMessage(mes).SetData(data).Output()
}

func RError(
	c *gin.Context,
	err error,
	code int,
	data any,
) {
	New(c).SetType(JSON).SetState(Error).SetMessage(err.Error()).SetCode(code).SetData(data).Output()
}

func RJump(c *gin.Context, data any, msg string) {
	New(c).SetType(JSON).SetState(Error).SetMessage(msg).SetCode(http.StatusSeeOther).SetData(data).Output()
}
