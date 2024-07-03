package resp

import "github.com/nicksnyder/go-i18n/v2/i18n"

var (
	SuccessMessage = &i18n.Message{
		ID:    "STATE_SUCCESS",
		Other: "Success",
	}

	ErrorMessage = &i18n.Message{
		ID:    "STATE_ERROR",
		Other: "Error !",
	}

	WrongMessage = &i18n.Message{
		ID:    "STATE_WRONG",
		Other: "Wrong !",
	}

	FailMessage = &i18n.Message{
		ID:    "STATE.Fail",
		Other: "Fail !",
	}
)

func SetSuccessMessage(message *i18n.Message) {
	SuccessMessage = message
}

func SetErrorMessage(message *i18n.Message) {
	ErrorMessage = message
}

func SetWrongMessage(message *i18n.Message) {
	WrongMessage = message
}

func SetFailMessage(message *i18n.Message) {
	FailMessage = message
}
