package repo

import (
	"ent"
	"ent/models"
	"framework/db"
	"framework/utils"
	"sync"
)

var once = sync.Once{}

var Client *models.Client = nil

func Init() {
	once.Do(func() {
		conns := db.DB{}

		conns.Init()

		Client = ent.NewClient([]models.Option{
			models.Driver(&db.MultiDriver{R: conns.Read(), W: conns.Write()}),
		}, utils.Path.Join(utils.Path.StoragePath, "migrate.sql"))
	})
}
