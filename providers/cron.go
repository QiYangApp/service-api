package providers

import (
	"service-api/internal/jobs"
)

func CronRegister() {
	jobs.Conf()
}
