package messages

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	Success = &i18n.Message{
		ID:    "STATE_SUCCESS",
		Other: "Success",
	}

	Error = &i18n.Message{
		ID:    "STATE_ERROR",
		Other: "Error !",
	}

	Fail = &i18n.Message{
		ID:    "STATE.Fail",
		Other: "Fail !",
	}
)
