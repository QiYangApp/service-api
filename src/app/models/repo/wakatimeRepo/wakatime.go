package wakatimeRepo

import (
	"service-api/src/models/ent"
	"service-api/src/models/ent/wakatime"
	"service-api/src/models/repo"
)

func AllWakatimeByState(state string) []*ent.Wakatime {
	return repo.Query().Wakatime.Query().Where(wakatime.StateEQ(state)).AllX(repo.Ctx())
}
