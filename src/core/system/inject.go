package system

import (
	"github.com/gin-gonic/gin"
	"service-api/src/core/config"
	"service-api/src/core/inject"
)

type InjectService struct {
	service
}

func (d *InjectService) Handle(r *gin.Engine, cfg *config.ConfigService) {
	inject.Apply(r, true)
}
