package providers

import (
	"service-api/internal/repo"
)

type Database struct {
}

func (*Database) Register() {
	repo.New()
}
