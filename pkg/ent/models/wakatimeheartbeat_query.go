// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/wakatimeheartbeat"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WakatimeHeartBeatQuery is the builder for querying WakatimeHeartBeat entities.
type WakatimeHeartBeatQuery struct {
	config
	ctx        *QueryContext
	order      []wakatimeheartbeat.OrderOption
	inters     []Interceptor
	predicates []predicate.WakatimeHeartBeat
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WakatimeHeartBeatQuery builder.
func (whbq *WakatimeHeartBeatQuery) Where(ps ...predicate.WakatimeHeartBeat) *WakatimeHeartBeatQuery {
	whbq.predicates = append(whbq.predicates, ps...)
	return whbq
}

// Limit the number of records to be returned by this query.
func (whbq *WakatimeHeartBeatQuery) Limit(limit int) *WakatimeHeartBeatQuery {
	whbq.ctx.Limit = &limit
	return whbq
}

// Offset to start from.
func (whbq *WakatimeHeartBeatQuery) Offset(offset int) *WakatimeHeartBeatQuery {
	whbq.ctx.Offset = &offset
	return whbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (whbq *WakatimeHeartBeatQuery) Unique(unique bool) *WakatimeHeartBeatQuery {
	whbq.ctx.Unique = &unique
	return whbq
}

// Order specifies how the records should be ordered.
func (whbq *WakatimeHeartBeatQuery) Order(o ...wakatimeheartbeat.OrderOption) *WakatimeHeartBeatQuery {
	whbq.order = append(whbq.order, o...)
	return whbq
}

// First returns the first WakatimeHeartBeat entity from the query.
// Returns a *NotFoundError when no WakatimeHeartBeat was found.
func (whbq *WakatimeHeartBeatQuery) First(ctx context.Context) (*WakatimeHeartBeat, error) {
	nodes, err := whbq.Limit(1).All(setContextOp(ctx, whbq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{wakatimeheartbeat.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (whbq *WakatimeHeartBeatQuery) FirstX(ctx context.Context) *WakatimeHeartBeat {
	node, err := whbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WakatimeHeartBeat ID from the query.
// Returns a *NotFoundError when no WakatimeHeartBeat ID was found.
func (whbq *WakatimeHeartBeatQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = whbq.Limit(1).IDs(setContextOp(ctx, whbq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{wakatimeheartbeat.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (whbq *WakatimeHeartBeatQuery) FirstIDX(ctx context.Context) int {
	id, err := whbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WakatimeHeartBeat entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WakatimeHeartBeat entity is found.
// Returns a *NotFoundError when no WakatimeHeartBeat entities are found.
func (whbq *WakatimeHeartBeatQuery) Only(ctx context.Context) (*WakatimeHeartBeat, error) {
	nodes, err := whbq.Limit(2).All(setContextOp(ctx, whbq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{wakatimeheartbeat.Label}
	default:
		return nil, &NotSingularError{wakatimeheartbeat.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (whbq *WakatimeHeartBeatQuery) OnlyX(ctx context.Context) *WakatimeHeartBeat {
	node, err := whbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WakatimeHeartBeat ID in the query.
// Returns a *NotSingularError when more than one WakatimeHeartBeat ID is found.
// Returns a *NotFoundError when no entities are found.
func (whbq *WakatimeHeartBeatQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = whbq.Limit(2).IDs(setContextOp(ctx, whbq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{wakatimeheartbeat.Label}
	default:
		err = &NotSingularError{wakatimeheartbeat.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (whbq *WakatimeHeartBeatQuery) OnlyIDX(ctx context.Context) int {
	id, err := whbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WakatimeHeartBeats.
func (whbq *WakatimeHeartBeatQuery) All(ctx context.Context) ([]*WakatimeHeartBeat, error) {
	ctx = setContextOp(ctx, whbq.ctx, "All")
	if err := whbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WakatimeHeartBeat, *WakatimeHeartBeatQuery]()
	return withInterceptors[[]*WakatimeHeartBeat](ctx, whbq, qr, whbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (whbq *WakatimeHeartBeatQuery) AllX(ctx context.Context) []*WakatimeHeartBeat {
	nodes, err := whbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WakatimeHeartBeat IDs.
func (whbq *WakatimeHeartBeatQuery) IDs(ctx context.Context) (ids []int, err error) {
	if whbq.ctx.Unique == nil && whbq.path != nil {
		whbq.Unique(true)
	}
	ctx = setContextOp(ctx, whbq.ctx, "IDs")
	if err = whbq.Select(wakatimeheartbeat.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (whbq *WakatimeHeartBeatQuery) IDsX(ctx context.Context) []int {
	ids, err := whbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (whbq *WakatimeHeartBeatQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, whbq.ctx, "Count")
	if err := whbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, whbq, querierCount[*WakatimeHeartBeatQuery](), whbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (whbq *WakatimeHeartBeatQuery) CountX(ctx context.Context) int {
	count, err := whbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (whbq *WakatimeHeartBeatQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, whbq.ctx, "Exist")
	switch _, err := whbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("models: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (whbq *WakatimeHeartBeatQuery) ExistX(ctx context.Context) bool {
	exist, err := whbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WakatimeHeartBeatQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (whbq *WakatimeHeartBeatQuery) Clone() *WakatimeHeartBeatQuery {
	if whbq == nil {
		return nil
	}
	return &WakatimeHeartBeatQuery{
		config:     whbq.config,
		ctx:        whbq.ctx.Clone(),
		order:      append([]wakatimeheartbeat.OrderOption{}, whbq.order...),
		inters:     append([]Interceptor{}, whbq.inters...),
		predicates: append([]predicate.WakatimeHeartBeat{}, whbq.predicates...),
		// clone intermediate query.
		sql:  whbq.sql.Clone(),
		path: whbq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (whbq *WakatimeHeartBeatQuery) GroupBy(field string, fields ...string) *WakatimeHeartBeatGroupBy {
	whbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WakatimeHeartBeatGroupBy{build: whbq}
	grbuild.flds = &whbq.ctx.Fields
	grbuild.label = wakatimeheartbeat.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (whbq *WakatimeHeartBeatQuery) Select(fields ...string) *WakatimeHeartBeatSelect {
	whbq.ctx.Fields = append(whbq.ctx.Fields, fields...)
	sbuild := &WakatimeHeartBeatSelect{WakatimeHeartBeatQuery: whbq}
	sbuild.label = wakatimeheartbeat.Label
	sbuild.flds, sbuild.scan = &whbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WakatimeHeartBeatSelect configured with the given aggregations.
func (whbq *WakatimeHeartBeatQuery) Aggregate(fns ...AggregateFunc) *WakatimeHeartBeatSelect {
	return whbq.Select().Aggregate(fns...)
}

func (whbq *WakatimeHeartBeatQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range whbq.inters {
		if inter == nil {
			return fmt.Errorf("models: uninitialized interceptor (forgotten import models/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, whbq); err != nil {
				return err
			}
		}
	}
	for _, f := range whbq.ctx.Fields {
		if !wakatimeheartbeat.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if whbq.path != nil {
		prev, err := whbq.path(ctx)
		if err != nil {
			return err
		}
		whbq.sql = prev
	}
	return nil
}

func (whbq *WakatimeHeartBeatQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WakatimeHeartBeat, error) {
	var (
		nodes = []*WakatimeHeartBeat{}
		_spec = whbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WakatimeHeartBeat).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WakatimeHeartBeat{config: whbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, whbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (whbq *WakatimeHeartBeatQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := whbq.querySpec()
	_spec.Node.Columns = whbq.ctx.Fields
	if len(whbq.ctx.Fields) > 0 {
		_spec.Unique = whbq.ctx.Unique != nil && *whbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, whbq.driver, _spec)
}

func (whbq *WakatimeHeartBeatQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(wakatimeheartbeat.Table, wakatimeheartbeat.Columns, sqlgraph.NewFieldSpec(wakatimeheartbeat.FieldID, field.TypeInt))
	_spec.From = whbq.sql
	if unique := whbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if whbq.path != nil {
		_spec.Unique = true
	}
	if fields := whbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wakatimeheartbeat.FieldID)
		for i := range fields {
			if fields[i] != wakatimeheartbeat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := whbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := whbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := whbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := whbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (whbq *WakatimeHeartBeatQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(whbq.driver.Dialect())
	t1 := builder.Table(wakatimeheartbeat.Table)
	columns := whbq.ctx.Fields
	if len(columns) == 0 {
		columns = wakatimeheartbeat.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if whbq.sql != nil {
		selector = whbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if whbq.ctx.Unique != nil && *whbq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range whbq.predicates {
		p(selector)
	}
	for _, p := range whbq.order {
		p(selector)
	}
	if offset := whbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := whbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WakatimeHeartBeatGroupBy is the group-by builder for WakatimeHeartBeat entities.
type WakatimeHeartBeatGroupBy struct {
	selector
	build *WakatimeHeartBeatQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (whbgb *WakatimeHeartBeatGroupBy) Aggregate(fns ...AggregateFunc) *WakatimeHeartBeatGroupBy {
	whbgb.fns = append(whbgb.fns, fns...)
	return whbgb
}

// Scan applies the selector query and scans the result into the given value.
func (whbgb *WakatimeHeartBeatGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, whbgb.build.ctx, "GroupBy")
	if err := whbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WakatimeHeartBeatQuery, *WakatimeHeartBeatGroupBy](ctx, whbgb.build, whbgb, whbgb.build.inters, v)
}

func (whbgb *WakatimeHeartBeatGroupBy) sqlScan(ctx context.Context, root *WakatimeHeartBeatQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(whbgb.fns))
	for _, fn := range whbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*whbgb.flds)+len(whbgb.fns))
		for _, f := range *whbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*whbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := whbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WakatimeHeartBeatSelect is the builder for selecting fields of WakatimeHeartBeat entities.
type WakatimeHeartBeatSelect struct {
	*WakatimeHeartBeatQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (whbs *WakatimeHeartBeatSelect) Aggregate(fns ...AggregateFunc) *WakatimeHeartBeatSelect {
	whbs.fns = append(whbs.fns, fns...)
	return whbs
}

// Scan applies the selector query and scans the result into the given value.
func (whbs *WakatimeHeartBeatSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, whbs.ctx, "Select")
	if err := whbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WakatimeHeartBeatQuery, *WakatimeHeartBeatSelect](ctx, whbs.WakatimeHeartBeatQuery, whbs, whbs.inters, v)
}

func (whbs *WakatimeHeartBeatSelect) sqlScan(ctx context.Context, root *WakatimeHeartBeatQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(whbs.fns))
	for _, fn := range whbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*whbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := whbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
