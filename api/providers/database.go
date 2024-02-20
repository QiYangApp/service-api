package providers

import (
	"framework"
	"service-api/repo"
)

type Database struct {
}

func (*Database) Register(app *framework.App) {
	repo.New()
}
