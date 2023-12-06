package app

import "github.com/gin-gonic/gin"

type Provider interface {
	Register(r *gin.Engine, cfg *config.ConfigService)
}
