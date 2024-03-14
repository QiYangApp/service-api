package providers

import (
	"service-api/internal/models"
)

type Database struct {
}

func (*Database) Register() {
	models.New()
}
