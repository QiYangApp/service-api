package messages

import "github.com/nicksnyder/go-i18n/v2/i18n"

var (
	CaptchaValidationFAILED = &i18n.Message{
		ID:    "CAPTCHA_VALIDATION_FAILED",
		Other: "Captcha validation failed",
	}

	CaptchaStorageFAILED = &i18n.Message{
		ID:    "CAPTCHA_STORAGE_FAILURE",
		Other: "Failed to store generated captcha",
	}

	CaptchaCreationFAILED = &i18n.Message{
		ID:    "CAPTCHA_CREATION_FAILED",
		Other: "Error generating captcha",
	}

	CaptchaGenerationFailed = &i18n.Message{
		ID:    "CAPTCHA_GENERATION_FAILED",
		Other: "Captcha generation failed",
	}

	CaptchaInputEmpty = &i18n.Message{
		ID:    "CAPTCHA_INPUT_EMPTY",
		Other: "Captcha value not provided",
	}

	CaptchaIdEmpty = &i18n.Message{
		ID:    "CAPTCHA_ID_EMPTY",
		Other: "Captcha id missing",
	}
	CaptchaTokenMissing = &i18n.Message{
		ID:    "CAPTCHA_TOKEN_MISSING",
		Other: "Captcha token is missing.",
	}
	CaptchaNotActivated = &i18n.Message{
		ID:    "CAPTCHA_NOT_ACTIVATED",
		Other: "Captcha is not activated.",
	}
)
