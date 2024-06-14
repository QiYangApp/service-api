// "github.com/1Panel-dev/1Panel/backend/buserr"

package errs

import (
	"runtime"
)

type Info struct {
	MessageId    string `json:"message"`
	Detail       any
	TemplateData map[string]any `json:"templateData"`
	Err          error          `json:"err"`
	Caller       []uintptr      `json:"caller"`
}

func (e Info) Error() string {
	return e.MessageId
}

func New(messageId string) *Info {
	return &Info{
		MessageId:    messageId,
		Caller:       nil,
		TemplateData: nil,
		Err:          nil,
	}
}

func NewCaller(messageId string) *Info {
	return &Info{
		MessageId:    messageId,
		Caller:       callers(),
		TemplateData: nil,
		Err:          nil,
	}
}

func WithDetail(messageId string, detail any, err error) *Info {
	paramMap := map[string]any{}
	if detail != nil {
		paramMap["detail"] = detail
	}

	return &Info{
		MessageId:    messageId,
		Detail:       detail,
		TemplateData: paramMap,
		Err:          err,
	}
}

func WithErr(messageId string, err error) *Info {
	paramMap := map[string]any{}
	if err != nil {
		paramMap["err"] = err
	}
	return &Info{
		MessageId:    messageId,
		TemplateData: paramMap,
		Err:          err,
	}
}

func WithMap(messageId string, templateData map[string]any, err error) *Info {
	return &Info{
		MessageId:    messageId,
		TemplateData: templateData,
		Err:          err,
	}
}

func callers() []uintptr {
	var pcs [32]uintptr
	l := runtime.Callers(3, pcs[:])
	return pcs[:l]
}
