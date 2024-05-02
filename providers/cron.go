package providers

import (
	"service-api/internal/app/jobs"
)

func CronRegister() {
	jobs.Conf()
}
