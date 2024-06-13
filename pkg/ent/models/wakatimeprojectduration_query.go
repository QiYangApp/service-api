// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/wakatimeprojectduration"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WakatimeProjectDurationQuery is the builder for querying WakatimeProjectDuration entities.
type WakatimeProjectDurationQuery struct {
	config
	ctx        *QueryContext
	order      []wakatimeprojectduration.OrderOption
	inters     []Interceptor
	predicates []predicate.WakatimeProjectDuration
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*WakatimeProjectDuration) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WakatimeProjectDurationQuery builder.
func (wpdq *WakatimeProjectDurationQuery) Where(ps ...predicate.WakatimeProjectDuration) *WakatimeProjectDurationQuery {
	wpdq.predicates = append(wpdq.predicates, ps...)
	return wpdq
}

// Limit the number of records to be returned by this query.
func (wpdq *WakatimeProjectDurationQuery) Limit(limit int) *WakatimeProjectDurationQuery {
	wpdq.ctx.Limit = &limit
	return wpdq
}

// Offset to start from.
func (wpdq *WakatimeProjectDurationQuery) Offset(offset int) *WakatimeProjectDurationQuery {
	wpdq.ctx.Offset = &offset
	return wpdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wpdq *WakatimeProjectDurationQuery) Unique(unique bool) *WakatimeProjectDurationQuery {
	wpdq.ctx.Unique = &unique
	return wpdq
}

// Order specifies how the records should be ordered.
func (wpdq *WakatimeProjectDurationQuery) Order(o ...wakatimeprojectduration.OrderOption) *WakatimeProjectDurationQuery {
	wpdq.order = append(wpdq.order, o...)
	return wpdq
}

// First returns the first WakatimeProjectDuration entity from the query.
// Returns a *NotFoundError when no WakatimeProjectDuration was found.
func (wpdq *WakatimeProjectDurationQuery) First(ctx context.Context) (*WakatimeProjectDuration, error) {
	nodes, err := wpdq.Limit(1).All(setContextOp(ctx, wpdq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{wakatimeprojectduration.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wpdq *WakatimeProjectDurationQuery) FirstX(ctx context.Context) *WakatimeProjectDuration {
	node, err := wpdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WakatimeProjectDuration ID from the query.
// Returns a *NotFoundError when no WakatimeProjectDuration ID was found.
func (wpdq *WakatimeProjectDurationQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = wpdq.Limit(1).IDs(setContextOp(ctx, wpdq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{wakatimeprojectduration.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wpdq *WakatimeProjectDurationQuery) FirstIDX(ctx context.Context) int64 {
	id, err := wpdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WakatimeProjectDuration entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WakatimeProjectDuration entity is found.
// Returns a *NotFoundError when no WakatimeProjectDuration entities are found.
func (wpdq *WakatimeProjectDurationQuery) Only(ctx context.Context) (*WakatimeProjectDuration, error) {
	nodes, err := wpdq.Limit(2).All(setContextOp(ctx, wpdq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{wakatimeprojectduration.Label}
	default:
		return nil, &NotSingularError{wakatimeprojectduration.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wpdq *WakatimeProjectDurationQuery) OnlyX(ctx context.Context) *WakatimeProjectDuration {
	node, err := wpdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WakatimeProjectDuration ID in the query.
// Returns a *NotSingularError when more than one WakatimeProjectDuration ID is found.
// Returns a *NotFoundError when no entities are found.
func (wpdq *WakatimeProjectDurationQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = wpdq.Limit(2).IDs(setContextOp(ctx, wpdq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{wakatimeprojectduration.Label}
	default:
		err = &NotSingularError{wakatimeprojectduration.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wpdq *WakatimeProjectDurationQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := wpdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WakatimeProjectDurations.
func (wpdq *WakatimeProjectDurationQuery) All(ctx context.Context) ([]*WakatimeProjectDuration, error) {
	ctx = setContextOp(ctx, wpdq.ctx, "All")
	if err := wpdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WakatimeProjectDuration, *WakatimeProjectDurationQuery]()
	return withInterceptors[[]*WakatimeProjectDuration](ctx, wpdq, qr, wpdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wpdq *WakatimeProjectDurationQuery) AllX(ctx context.Context) []*WakatimeProjectDuration {
	nodes, err := wpdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WakatimeProjectDuration IDs.
func (wpdq *WakatimeProjectDurationQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if wpdq.ctx.Unique == nil && wpdq.path != nil {
		wpdq.Unique(true)
	}
	ctx = setContextOp(ctx, wpdq.ctx, "IDs")
	if err = wpdq.Select(wakatimeprojectduration.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wpdq *WakatimeProjectDurationQuery) IDsX(ctx context.Context) []int64 {
	ids, err := wpdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wpdq *WakatimeProjectDurationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wpdq.ctx, "Count")
	if err := wpdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wpdq, querierCount[*WakatimeProjectDurationQuery](), wpdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wpdq *WakatimeProjectDurationQuery) CountX(ctx context.Context) int {
	count, err := wpdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wpdq *WakatimeProjectDurationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wpdq.ctx, "Exist")
	switch _, err := wpdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("models: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wpdq *WakatimeProjectDurationQuery) ExistX(ctx context.Context) bool {
	exist, err := wpdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WakatimeProjectDurationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wpdq *WakatimeProjectDurationQuery) Clone() *WakatimeProjectDurationQuery {
	if wpdq == nil {
		return nil
	}
	return &WakatimeProjectDurationQuery{
		config:     wpdq.config,
		ctx:        wpdq.ctx.Clone(),
		order:      append([]wakatimeprojectduration.OrderOption{}, wpdq.order...),
		inters:     append([]Interceptor{}, wpdq.inters...),
		predicates: append([]predicate.WakatimeProjectDuration{}, wpdq.predicates...),
		// clone intermediate query.
		sql:  wpdq.sql.Clone(),
		path: wpdq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (wpdq *WakatimeProjectDurationQuery) GroupBy(field string, fields ...string) *WakatimeProjectDurationGroupBy {
	wpdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WakatimeProjectDurationGroupBy{build: wpdq}
	grbuild.flds = &wpdq.ctx.Fields
	grbuild.label = wakatimeprojectduration.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (wpdq *WakatimeProjectDurationQuery) Select(fields ...string) *WakatimeProjectDurationSelect {
	wpdq.ctx.Fields = append(wpdq.ctx.Fields, fields...)
	sbuild := &WakatimeProjectDurationSelect{WakatimeProjectDurationQuery: wpdq}
	sbuild.label = wakatimeprojectduration.Label
	sbuild.flds, sbuild.scan = &wpdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WakatimeProjectDurationSelect configured with the given aggregations.
func (wpdq *WakatimeProjectDurationQuery) Aggregate(fns ...AggregateFunc) *WakatimeProjectDurationSelect {
	return wpdq.Select().Aggregate(fns...)
}

func (wpdq *WakatimeProjectDurationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wpdq.inters {
		if inter == nil {
			return fmt.Errorf("models: uninitialized interceptor (forgotten import models/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wpdq); err != nil {
				return err
			}
		}
	}
	for _, f := range wpdq.ctx.Fields {
		if !wakatimeprojectduration.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if wpdq.path != nil {
		prev, err := wpdq.path(ctx)
		if err != nil {
			return err
		}
		wpdq.sql = prev
	}
	return nil
}

func (wpdq *WakatimeProjectDurationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WakatimeProjectDuration, error) {
	var (
		nodes = []*WakatimeProjectDuration{}
		_spec = wpdq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WakatimeProjectDuration).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WakatimeProjectDuration{config: wpdq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(wpdq.modifiers) > 0 {
		_spec.Modifiers = wpdq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wpdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range wpdq.loadTotal {
		if err := wpdq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wpdq *WakatimeProjectDurationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wpdq.querySpec()
	if len(wpdq.modifiers) > 0 {
		_spec.Modifiers = wpdq.modifiers
	}
	_spec.Node.Columns = wpdq.ctx.Fields
	if len(wpdq.ctx.Fields) > 0 {
		_spec.Unique = wpdq.ctx.Unique != nil && *wpdq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wpdq.driver, _spec)
}

func (wpdq *WakatimeProjectDurationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(wakatimeprojectduration.Table, wakatimeprojectduration.Columns, sqlgraph.NewFieldSpec(wakatimeprojectduration.FieldID, field.TypeInt64))
	_spec.From = wpdq.sql
	if unique := wpdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wpdq.path != nil {
		_spec.Unique = true
	}
	if fields := wpdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wakatimeprojectduration.FieldID)
		for i := range fields {
			if fields[i] != wakatimeprojectduration.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wpdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wpdq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wpdq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wpdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wpdq *WakatimeProjectDurationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wpdq.driver.Dialect())
	t1 := builder.Table(wakatimeprojectduration.Table)
	columns := wpdq.ctx.Fields
	if len(columns) == 0 {
		columns = wakatimeprojectduration.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wpdq.sql != nil {
		selector = wpdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wpdq.ctx.Unique != nil && *wpdq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wpdq.predicates {
		p(selector)
	}
	for _, p := range wpdq.order {
		p(selector)
	}
	if offset := wpdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wpdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WakatimeProjectDurationGroupBy is the group-by builder for WakatimeProjectDuration entities.
type WakatimeProjectDurationGroupBy struct {
	selector
	build *WakatimeProjectDurationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wpdgb *WakatimeProjectDurationGroupBy) Aggregate(fns ...AggregateFunc) *WakatimeProjectDurationGroupBy {
	wpdgb.fns = append(wpdgb.fns, fns...)
	return wpdgb
}

// Scan applies the selector query and scans the result into the given value.
func (wpdgb *WakatimeProjectDurationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wpdgb.build.ctx, "GroupBy")
	if err := wpdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WakatimeProjectDurationQuery, *WakatimeProjectDurationGroupBy](ctx, wpdgb.build, wpdgb, wpdgb.build.inters, v)
}

func (wpdgb *WakatimeProjectDurationGroupBy) sqlScan(ctx context.Context, root *WakatimeProjectDurationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wpdgb.fns))
	for _, fn := range wpdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wpdgb.flds)+len(wpdgb.fns))
		for _, f := range *wpdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wpdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wpdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WakatimeProjectDurationSelect is the builder for selecting fields of WakatimeProjectDuration entities.
type WakatimeProjectDurationSelect struct {
	*WakatimeProjectDurationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wpds *WakatimeProjectDurationSelect) Aggregate(fns ...AggregateFunc) *WakatimeProjectDurationSelect {
	wpds.fns = append(wpds.fns, fns...)
	return wpds
}

// Scan applies the selector query and scans the result into the given value.
func (wpds *WakatimeProjectDurationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wpds.ctx, "Select")
	if err := wpds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WakatimeProjectDurationQuery, *WakatimeProjectDurationSelect](ctx, wpds.WakatimeProjectDurationQuery, wpds, wpds.inters, v)
}

func (wpds *WakatimeProjectDurationSelect) sqlScan(ctx context.Context, root *WakatimeProjectDurationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wpds.fns))
	for _, fn := range wpds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wpds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wpds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
