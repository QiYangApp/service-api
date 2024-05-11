package auth

import "github.com/gin-gonic/gin"

const IsSignedKey = "IsSigned"

func IsSigned(ctx *gin.Context) bool {
	return ctx.GetBool(IsSignedKey)
}

func SetIsSigned(ctx *gin.Context, token string, singed bool) {

	//Manually set up login
	if singed {
		ctx.Set(IsSignedKey, singed)
		return
	}

	return
}
