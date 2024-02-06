package providers

import (
	"app"
	"service-api/src/api/jobs"
)

type Cron struct {
}

func (*Cron) Register(app *app.App) {
	jobs.Conf()
}
