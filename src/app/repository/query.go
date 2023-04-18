package repository

import "entgo.io/ent"

type QueryInterfaceMethods interface {
	ent() *ent.Query
}

type QueryMethods struct {
	QueryInterfaceMethods
}

func (q *QueryMethods) Query() *ent.Query {
	return q.ent()
}
