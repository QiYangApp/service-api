package jobs

import (
	"service-api/src/api/command"
	"service-api/src/core/cron"
)

func Conf(c *cron.Instance) {
	//_, _ = c.AddFun(cron.TaskTime().EverySecond(1), func() {
	//	logger.D().Info("test", zap.String("test", "test"))
	//})

	_, _ = c.AddJob(cron.TaskTime().EverySecond(10), &command.WakatimeSourceData{})
}
