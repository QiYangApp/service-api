package jobs

import (
	"app/cron"
	"app/log"
	"go.uber.org/zap"
)

func Conf() {
	//_, _ = c.AddFun(cron.TaskTime().EverySecond(1), func() {
	//	logger.D().Info("test", zap.String("test", "test"))
	//})

	//_, _ = cron.Instance().AddJob(cron.TaskTime().EverySecond(10), &command.WakatimeSourceData{})
	a, err := cron.Instance().AddFun(cron.TaskTime().EverySecond(1), func() {
		log.Client().Debug("test", zap.String("TEST", "test"))
	})

	log.Client().Debug("test", zap.Any("a", a), zap.Error(err))
}
