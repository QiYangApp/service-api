package db

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

type MultiDriver struct {
	R, W dialect.Driver
}

var _ dialect.Driver = (*MultiDriver)(nil)

func (d *MultiDriver) Query(ctx context.Context, query string, args, v any) error {
	e := d.R
	// Mutation statements that use the RETURNING clause.
	if ent.QueryFromContext(ctx) == nil {
		e = d.W
	}
	return e.Query(ctx, query, args, v)
}

func (d *MultiDriver) Exec(ctx context.Context, query string, args, v any) error {
	return d.W.Exec(ctx, query, args, v)
}

func (d *MultiDriver) Tx(ctx context.Context) (dialect.Tx, error) {
	return d.W.Tx(ctx)
}

func (d *MultiDriver) BeginTx(ctx context.Context, opts *sql.TxOptions) (dialect.Tx, error) {
	return d.W.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
}

func (d *MultiDriver) Close() error {
	rerr := d.R.Close()
	werr := d.W.Close()
	if rerr != nil {
		return rerr
	}
	if werr != nil {
		return werr
	}
	return nil
}

func (d *MultiDriver) Dialect() string {
	return d.R.Dialect()
}
