package repo

import (
	"context"
	"service-api/src/core/db"
	"service-api/src/models/ent"
)

func Query() *ent.Client {
	return db.Client()
}

func Ctx() context.Context {
	return context.TODO()
}

type BaseModel struct {
}

func (*BaseModel) Query() *ent.Client {
	return Query()
}
func (*BaseModel) Ctx() context.Context {
	return Ctx()
}
