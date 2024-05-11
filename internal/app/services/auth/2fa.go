package auth

import (
	"ent/models"
	"errors"
	"frame/modules/cache"
	"frame/util/secret"
	"service-api/internal/app/repo/auth"
	"service-api/internal/modules/setting"

	"github.com/gin-gonic/gin"
)

func HasUser2FA(ctx *gin.Context, u *models.User) (bool, error) {

	// If this user is enrolled in 2FA TOTP, we can't sign the user in just yet.
	// Instead, redirect them to the 2FA authentication page.
	hasTOTPtwofa, err := auth.HasTwoFactorByUID(ctx, u.ID)
	if err != nil {
		return false, err
	}

	// Check if the user has webauthn registration
	hasWebAuthnTwofa, err := auth.HasWebAuthnRegistrationsByUID(ctx, u.ID)
	if err != nil {
		return false, err
	}

	return hasTOTPtwofa && hasWebAuthnTwofa, nil
}

func HasUserTwoFactor(ctx *gin.Context, userId int64) (bool, error) {
	return auth.HasTwoFactorByUID(ctx, userId)
}

func HasUserWebAuthn(ctx *gin.Context, userId int64) (bool, error) {
	return auth.HasWebAuthnRegistrationsByUID(ctx, userId)
}

func gen2FATmpToken(u *models.User) (string, error) {
	token := secret.Sha3Sum512Md5(u.String(), u.PasswdSalt)

	if !cache.SetEx(token, u, setting.AuthSetting.TwoFA.Expires) {
		return "", errors.New("SIGN_IN.2FA_GEN_FAIL")
	}

	return token, nil

}
