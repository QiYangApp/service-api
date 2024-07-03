package providers

import (
	"frame/modules/resp"
	"service-api/resources/translate/messages"
)

func MessageRegister() {
	resp.SetErrorMessage(messages.Error)
	resp.SetSuccessMessage(messages.Success)
	resp.SetFailMessage(messages.Fail)
	resp.SetWrongMessage(messages.Wrong)
}
