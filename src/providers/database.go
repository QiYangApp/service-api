package providers

import (
	"app"
	"service-api/src/repo"
)

type Database struct {
}

func (*Database) Register(app *app.App) {
	repo.New()
}
