package scripts

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func state() {
	translator.LocalizeMessage(
		&i18n.Message{
			ID:    "STATE.SUCCESS",
			Other: "Success",
		},
	)

	translator.LocalizeMessage(
		&i18n.Message{
			ID:    "STATE.Error",
			Other: "Success",
		},
	)

	translator.LocalizeMessage(
		&i18n.Message{
			ID:    "STATE.Fail",
			Other: "Fail",
		},
	)
}
