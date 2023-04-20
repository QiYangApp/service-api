package repository

import "service-api/src/ent"

type AuthorizeRepository struct {
	Query
}

func (a *AuthorizeRepository) findSingleById(id int64) *ent.Member {
	return a.NewQuery()
}
