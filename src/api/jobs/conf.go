package jobs

import (
	"app/cron"
)

func Conf() {
	//_, _ = c.AddFun(cron.TaskTime().EverySecond(1), func() {
	//	logger.D().Info("test", zap.String("test", "test"))
	//})

	//_, _ = cron.Instance().AddJob(cron.TaskTime().EverySecond(10), &command.WakatimeSourceData{})
	_, _ = cron.Instance().AddFun(cron.TaskTime().EverySecond(1), func() {
		//log.Client().Debug("test", zap.String("TEST", "test"))
	})

}
