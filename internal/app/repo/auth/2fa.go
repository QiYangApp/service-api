package auth

import (
	"ent/models/twofactor"
	"ent/models/webauthncredential"
	"github.com/gin-gonic/gin"
	"service-api/internal/app/repo"
)

func HasTwoFactorByUID(ctx *gin.Context, userId int64) (bool, error) {
	return repo.Client().TwoFactor.Query().Where(twofactor.UserIDEQ(userId)).Exist(ctx)
}

func HasWebAuthnRegistrationsByUID(ctx *gin.Context, userId int64) (bool, error) {
	return repo.Client().WebAuthnCredential.Query().Where(webauthncredential.UserIDEQ(userId)).Exist(ctx)
}
