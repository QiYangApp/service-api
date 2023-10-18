package jobs

import (
	"service-api/src/core/cron"
)

func Conf(c *cron.Instance) {
	//_, _ = c.AddFun(cron.TaskTime().EverySecond(1), func() {
	//	logger.D().Info("test", zap.String("test", "test"))
	//})
}
