package providers

import (
	"framework"
	"service-api/api/jobs"
)

type Cron struct {
}

func (*Cron) Register(app *framework.App) {
	jobs.Conf()
}
