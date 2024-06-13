// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/wakatimesystem"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WakatimeSystemQuery is the builder for querying WakatimeSystem entities.
type WakatimeSystemQuery struct {
	config
	ctx        *QueryContext
	order      []wakatimesystem.OrderOption
	inters     []Interceptor
	predicates []predicate.WakatimeSystem
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*WakatimeSystem) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WakatimeSystemQuery builder.
func (wsq *WakatimeSystemQuery) Where(ps ...predicate.WakatimeSystem) *WakatimeSystemQuery {
	wsq.predicates = append(wsq.predicates, ps...)
	return wsq
}

// Limit the number of records to be returned by this query.
func (wsq *WakatimeSystemQuery) Limit(limit int) *WakatimeSystemQuery {
	wsq.ctx.Limit = &limit
	return wsq
}

// Offset to start from.
func (wsq *WakatimeSystemQuery) Offset(offset int) *WakatimeSystemQuery {
	wsq.ctx.Offset = &offset
	return wsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wsq *WakatimeSystemQuery) Unique(unique bool) *WakatimeSystemQuery {
	wsq.ctx.Unique = &unique
	return wsq
}

// Order specifies how the records should be ordered.
func (wsq *WakatimeSystemQuery) Order(o ...wakatimesystem.OrderOption) *WakatimeSystemQuery {
	wsq.order = append(wsq.order, o...)
	return wsq
}

// First returns the first WakatimeSystem entity from the query.
// Returns a *NotFoundError when no WakatimeSystem was found.
func (wsq *WakatimeSystemQuery) First(ctx context.Context) (*WakatimeSystem, error) {
	nodes, err := wsq.Limit(1).All(setContextOp(ctx, wsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{wakatimesystem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wsq *WakatimeSystemQuery) FirstX(ctx context.Context) *WakatimeSystem {
	node, err := wsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WakatimeSystem ID from the query.
// Returns a *NotFoundError when no WakatimeSystem ID was found.
func (wsq *WakatimeSystemQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = wsq.Limit(1).IDs(setContextOp(ctx, wsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{wakatimesystem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wsq *WakatimeSystemQuery) FirstIDX(ctx context.Context) int64 {
	id, err := wsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WakatimeSystem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WakatimeSystem entity is found.
// Returns a *NotFoundError when no WakatimeSystem entities are found.
func (wsq *WakatimeSystemQuery) Only(ctx context.Context) (*WakatimeSystem, error) {
	nodes, err := wsq.Limit(2).All(setContextOp(ctx, wsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{wakatimesystem.Label}
	default:
		return nil, &NotSingularError{wakatimesystem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wsq *WakatimeSystemQuery) OnlyX(ctx context.Context) *WakatimeSystem {
	node, err := wsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WakatimeSystem ID in the query.
// Returns a *NotSingularError when more than one WakatimeSystem ID is found.
// Returns a *NotFoundError when no entities are found.
func (wsq *WakatimeSystemQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = wsq.Limit(2).IDs(setContextOp(ctx, wsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{wakatimesystem.Label}
	default:
		err = &NotSingularError{wakatimesystem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wsq *WakatimeSystemQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := wsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WakatimeSystems.
func (wsq *WakatimeSystemQuery) All(ctx context.Context) ([]*WakatimeSystem, error) {
	ctx = setContextOp(ctx, wsq.ctx, "All")
	if err := wsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WakatimeSystem, *WakatimeSystemQuery]()
	return withInterceptors[[]*WakatimeSystem](ctx, wsq, qr, wsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wsq *WakatimeSystemQuery) AllX(ctx context.Context) []*WakatimeSystem {
	nodes, err := wsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WakatimeSystem IDs.
func (wsq *WakatimeSystemQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if wsq.ctx.Unique == nil && wsq.path != nil {
		wsq.Unique(true)
	}
	ctx = setContextOp(ctx, wsq.ctx, "IDs")
	if err = wsq.Select(wakatimesystem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wsq *WakatimeSystemQuery) IDsX(ctx context.Context) []int64 {
	ids, err := wsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wsq *WakatimeSystemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wsq.ctx, "Count")
	if err := wsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wsq, querierCount[*WakatimeSystemQuery](), wsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wsq *WakatimeSystemQuery) CountX(ctx context.Context) int {
	count, err := wsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wsq *WakatimeSystemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wsq.ctx, "Exist")
	switch _, err := wsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("models: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wsq *WakatimeSystemQuery) ExistX(ctx context.Context) bool {
	exist, err := wsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WakatimeSystemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wsq *WakatimeSystemQuery) Clone() *WakatimeSystemQuery {
	if wsq == nil {
		return nil
	}
	return &WakatimeSystemQuery{
		config:     wsq.config,
		ctx:        wsq.ctx.Clone(),
		order:      append([]wakatimesystem.OrderOption{}, wsq.order...),
		inters:     append([]Interceptor{}, wsq.inters...),
		predicates: append([]predicate.WakatimeSystem{}, wsq.predicates...),
		// clone intermediate query.
		sql:  wsq.sql.Clone(),
		path: wsq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (wsq *WakatimeSystemQuery) GroupBy(field string, fields ...string) *WakatimeSystemGroupBy {
	wsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WakatimeSystemGroupBy{build: wsq}
	grbuild.flds = &wsq.ctx.Fields
	grbuild.label = wakatimesystem.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (wsq *WakatimeSystemQuery) Select(fields ...string) *WakatimeSystemSelect {
	wsq.ctx.Fields = append(wsq.ctx.Fields, fields...)
	sbuild := &WakatimeSystemSelect{WakatimeSystemQuery: wsq}
	sbuild.label = wakatimesystem.Label
	sbuild.flds, sbuild.scan = &wsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WakatimeSystemSelect configured with the given aggregations.
func (wsq *WakatimeSystemQuery) Aggregate(fns ...AggregateFunc) *WakatimeSystemSelect {
	return wsq.Select().Aggregate(fns...)
}

func (wsq *WakatimeSystemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wsq.inters {
		if inter == nil {
			return fmt.Errorf("models: uninitialized interceptor (forgotten import models/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wsq); err != nil {
				return err
			}
		}
	}
	for _, f := range wsq.ctx.Fields {
		if !wakatimesystem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if wsq.path != nil {
		prev, err := wsq.path(ctx)
		if err != nil {
			return err
		}
		wsq.sql = prev
	}
	return nil
}

func (wsq *WakatimeSystemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WakatimeSystem, error) {
	var (
		nodes = []*WakatimeSystem{}
		_spec = wsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WakatimeSystem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WakatimeSystem{config: wsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(wsq.modifiers) > 0 {
		_spec.Modifiers = wsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range wsq.loadTotal {
		if err := wsq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wsq *WakatimeSystemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wsq.querySpec()
	if len(wsq.modifiers) > 0 {
		_spec.Modifiers = wsq.modifiers
	}
	_spec.Node.Columns = wsq.ctx.Fields
	if len(wsq.ctx.Fields) > 0 {
		_spec.Unique = wsq.ctx.Unique != nil && *wsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wsq.driver, _spec)
}

func (wsq *WakatimeSystemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(wakatimesystem.Table, wakatimesystem.Columns, sqlgraph.NewFieldSpec(wakatimesystem.FieldID, field.TypeInt64))
	_spec.From = wsq.sql
	if unique := wsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wsq.path != nil {
		_spec.Unique = true
	}
	if fields := wsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wakatimesystem.FieldID)
		for i := range fields {
			if fields[i] != wakatimesystem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wsq *WakatimeSystemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wsq.driver.Dialect())
	t1 := builder.Table(wakatimesystem.Table)
	columns := wsq.ctx.Fields
	if len(columns) == 0 {
		columns = wakatimesystem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wsq.sql != nil {
		selector = wsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wsq.ctx.Unique != nil && *wsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wsq.predicates {
		p(selector)
	}
	for _, p := range wsq.order {
		p(selector)
	}
	if offset := wsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WakatimeSystemGroupBy is the group-by builder for WakatimeSystem entities.
type WakatimeSystemGroupBy struct {
	selector
	build *WakatimeSystemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wsgb *WakatimeSystemGroupBy) Aggregate(fns ...AggregateFunc) *WakatimeSystemGroupBy {
	wsgb.fns = append(wsgb.fns, fns...)
	return wsgb
}

// Scan applies the selector query and scans the result into the given value.
func (wsgb *WakatimeSystemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wsgb.build.ctx, "GroupBy")
	if err := wsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WakatimeSystemQuery, *WakatimeSystemGroupBy](ctx, wsgb.build, wsgb, wsgb.build.inters, v)
}

func (wsgb *WakatimeSystemGroupBy) sqlScan(ctx context.Context, root *WakatimeSystemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wsgb.fns))
	for _, fn := range wsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wsgb.flds)+len(wsgb.fns))
		for _, f := range *wsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WakatimeSystemSelect is the builder for selecting fields of WakatimeSystem entities.
type WakatimeSystemSelect struct {
	*WakatimeSystemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wss *WakatimeSystemSelect) Aggregate(fns ...AggregateFunc) *WakatimeSystemSelect {
	wss.fns = append(wss.fns, fns...)
	return wss
}

// Scan applies the selector query and scans the result into the given value.
func (wss *WakatimeSystemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wss.ctx, "Select")
	if err := wss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WakatimeSystemQuery, *WakatimeSystemSelect](ctx, wss.WakatimeSystemQuery, wss, wss.inters, v)
}

func (wss *WakatimeSystemSelect) sqlScan(ctx context.Context, root *WakatimeSystemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wss.fns))
	for _, fn := range wss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
