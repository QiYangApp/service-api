package captcha

import (
	"frame/modules/resp"
	"github.com/gin-gonic/gin"
	"net/http"
	"service-api/conf"
	"service-api/resources/translate/messages"
)

func verify(t conf.CaptchaFeature, token, key string, answer any, clear bool) (bool, error) {
	if err := conf.IsCaptchaFeature(t); err != nil {
		return false, err
	}

	return New().Verify(token, key, answer, clear), nil
}

// Verify validates a captcha.
// This function receives a context, configuration, and a captcha verification form as parameters,
// and returns a response object or nil.
// If the captcha feature is enabled and validation fails, it returns an appropriate error response.
// If the captcha feature is not enabled, the function returns nil.
func Verify(ctx *gin.Context, t conf.CaptchaFeature, token, key string, answer any, clear bool) *resp.Response {
	if err := conf.IsCaptchaFeature(t); err != nil {
		return resp.Error(ctx, err, http.StatusPreconditionFailed, nil)
	}
	
	// Check if the captcha feature is enabled
	if conf.IsCaptchaFeatureEnable(t) {
		// Call the captcha service to perform validation
		st, err := verify(
			t,
			token,
			key,
			answer,
			clear,
		)

		// If there is an error during validation, return an error response
		if err != nil {
			return resp.Error(ctx, err, http.StatusPreconditionFailed, nil)
		}

		// If validation fails, return a failure message response
		if !st {
			return resp.ErrorWithMsg(ctx, messages.CaptchaValidationFAILED, http.StatusPreconditionFailed, nil)
		}
	}

	// If the captcha feature is not enabled or validation passes, return nil
	return nil
}
