package system

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	"service-api/src/core/config"
	"service-api/src/core/db"
)

type DatabaseService struct {
	service
	cfg config.Database
}

func (d *DatabaseService) Handle(r *gin.Engine, cfg *config.ConfigService) {
	_ = db.Init(cfg.GetDatabase())
}
