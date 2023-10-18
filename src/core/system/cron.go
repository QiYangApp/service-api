package system

import (
	"github.com/gin-gonic/gin"
	"service-api/src/api/jobs"
	"service-api/src/core/config"
	"service-api/src/core/cron"
)

type CronService struct {
	service
}

func (d *CronService) Handle(r *gin.Engine, cfg *config.ConfigService) {
	s := cron.S()

	jobs.Conf(s)

	s.Start()
}
