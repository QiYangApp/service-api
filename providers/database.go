package providers

import (
	"service-api/internal/repo"
)

func DatabaseRegister() {
	repo.Init()
}
