package providers

import "app"

type Database struct {
}

func (*Database) Register(appCaron *app.App) {
	jobs.Conf()
}
