package providers

import (
	"app"
	"app/db"
	"service-api/src/repo"
	"service-api/src/repo/models"
)

type Database struct {
}

func (*Database) Register(app *app.App) {
	client := db.DB{}
	client.Init()

	repo.New([]models.Option{
		models.Driver(&db.MultiDriver{R: client.Read(), W: client.Write()}),
	})
}
