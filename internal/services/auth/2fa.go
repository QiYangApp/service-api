package auth

import (
	"ent/models"
	"github.com/gin-gonic/gin"
	"service-api/internal/repo/auth"
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

func HasUserWebAuthn(ctx *gin.Context, userId int64) bool {
	hasWebAuthnTwofa, err := auth.HasWebAuthnRegistrationsByUID(ctx, userId)
	if err != nil {
		return false
	}

	return hasWebAuthnTwofa
}
