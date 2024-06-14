package messages

import "github.com/nicksnyder/go-i18n/v2/i18n"

var (
	UserEmailInUse = &i18n.Message{
		ID:    "EMAIL_IN_USE",
		Other: "This email address is already in use.",
	}
	UserEmailNotFound = &i18n.Message{
		ID:    "EMAIL_NOT_FOUND",
		Other: "The provided email address could not be found.",
	}

	UserEmailEmpty = &i18n.Message{
		ID:    "USER_EMAIL_EMPTY",
		Other: "Please provide an email address.",
	}

	UserEmailFormatInvalid = &i18n.Message{
		ID:    "USER_EMAIL_FORMAT_INVALID",
		Other: "The email address format is invalid.",
	}

	UserPasswordEmpty = &i18n.Message{
		ID:    "USER_PASSWORD_EMPTY",
		Other: "Please enter a password.",
	}

	UserPasswordCheckFailed = &i18n.Message{
		ID:    "USER_PASSWORD_CHECK_FAILED",
		Other: "The password does not meet the requirements.",
	}

	UserSignInInitLanguageFail = &i18n.Message{
		ID:    "USER_SIGN_IN_INIT_LANGUAGE_FAIL",
		Other: "Failed to initialize user's sign-in language.",
	}
)
