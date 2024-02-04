package providers

import (
	"app"
	"service-api/src/api/jobs"
)

type Cron struct {
}

func (*Cron) Register(appCaron *app.App) {
	jobs.Conf()
}
