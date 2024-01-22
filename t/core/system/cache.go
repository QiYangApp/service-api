package system

import (
	"github.com/gin-gonic/gin"
	"service-api/src/core/cache"
	"service-api/src/core/config"
)

type CacheService struct {
	service
}

func (d *CacheService) Handle(r *gin.Engine, cfg *config.ConfigService) {
	cache.NewInstance(cfg.GetCache().Driver)
}
