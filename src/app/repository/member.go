package repository

import (
	"context"
	"service-api/src/ent"
)

type MemberRepository struct {
	QueryMethods
}

func (a *MemberRepository) findSingleById(id int64) *ent.Member {
	return new(ent.MemberClient).Query().FirstX(context.Background())
}
