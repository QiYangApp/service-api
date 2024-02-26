package providers

import (
	"service-api/internal/app/jobs"
)

type Cron struct {
}

func (*Cron) Register() {
	jobs.Conf()
}
