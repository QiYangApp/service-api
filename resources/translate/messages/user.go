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

	UserNameCheckFailed = &i18n.Message{
		ID:    "USER_NAME_CHECK_FAILED",
		Other: "The name does not meet the requirements.",
	}

	UserNameEmpty = &i18n.Message{
		ID:    "USER_NAME_EMPTY",
		Other: "Please provide an userName.",
	}

	UserNameMinLength = &i18n.Message{
		ID:    "USER_NAME_MIN_LENGTH",
		Other: " Minimum allowed length for a username is {{.Len}}.",
	}

	UserNameMaxLength = &i18n.Message{
		ID:    "USER_NAME_MAX_LENGTH",
		Other: " Maximum allowed length for a username is {{.Len}}.",
	}

	UserEmailCheckFailed = &i18n.Message{
		ID:    "USER_NAME_CHECK_FAILED",
		Other: "The  email does not meet the requirements.",
	}

	UserEmailMaxLength = &i18n.Message{
		ID:    "USER_EMAIL_MAX_LENGTH",
		Other: " Maximum allowed length for a email is {{.Len}}.",
	}

	UserEmailMinLength = &i18n.Message{
		ID:    "USER_EMAIL_MAX_LENGTH",
		Other: " Maximum allowed length for a email is {{.Len}}.",
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
