// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/wakatimelanguage"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WakatimeLanguageQuery is the builder for querying WakatimeLanguage entities.
type WakatimeLanguageQuery struct {
	config
	ctx        *QueryContext
	order      []wakatimelanguage.OrderOption
	inters     []Interceptor
	predicates []predicate.WakatimeLanguage
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*WakatimeLanguage) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WakatimeLanguageQuery builder.
func (wlq *WakatimeLanguageQuery) Where(ps ...predicate.WakatimeLanguage) *WakatimeLanguageQuery {
	wlq.predicates = append(wlq.predicates, ps...)
	return wlq
}

// Limit the number of records to be returned by this query.
func (wlq *WakatimeLanguageQuery) Limit(limit int) *WakatimeLanguageQuery {
	wlq.ctx.Limit = &limit
	return wlq
}

// Offset to start from.
func (wlq *WakatimeLanguageQuery) Offset(offset int) *WakatimeLanguageQuery {
	wlq.ctx.Offset = &offset
	return wlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wlq *WakatimeLanguageQuery) Unique(unique bool) *WakatimeLanguageQuery {
	wlq.ctx.Unique = &unique
	return wlq
}

// Order specifies how the records should be ordered.
func (wlq *WakatimeLanguageQuery) Order(o ...wakatimelanguage.OrderOption) *WakatimeLanguageQuery {
	wlq.order = append(wlq.order, o...)
	return wlq
}

// First returns the first WakatimeLanguage entity from the query.
// Returns a *NotFoundError when no WakatimeLanguage was found.
func (wlq *WakatimeLanguageQuery) First(ctx context.Context) (*WakatimeLanguage, error) {
	nodes, err := wlq.Limit(1).All(setContextOp(ctx, wlq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{wakatimelanguage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wlq *WakatimeLanguageQuery) FirstX(ctx context.Context) *WakatimeLanguage {
	node, err := wlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WakatimeLanguage ID from the query.
// Returns a *NotFoundError when no WakatimeLanguage ID was found.
func (wlq *WakatimeLanguageQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = wlq.Limit(1).IDs(setContextOp(ctx, wlq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{wakatimelanguage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wlq *WakatimeLanguageQuery) FirstIDX(ctx context.Context) int64 {
	id, err := wlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WakatimeLanguage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WakatimeLanguage entity is found.
// Returns a *NotFoundError when no WakatimeLanguage entities are found.
func (wlq *WakatimeLanguageQuery) Only(ctx context.Context) (*WakatimeLanguage, error) {
	nodes, err := wlq.Limit(2).All(setContextOp(ctx, wlq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{wakatimelanguage.Label}
	default:
		return nil, &NotSingularError{wakatimelanguage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wlq *WakatimeLanguageQuery) OnlyX(ctx context.Context) *WakatimeLanguage {
	node, err := wlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WakatimeLanguage ID in the query.
// Returns a *NotSingularError when more than one WakatimeLanguage ID is found.
// Returns a *NotFoundError when no entities are found.
func (wlq *WakatimeLanguageQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = wlq.Limit(2).IDs(setContextOp(ctx, wlq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{wakatimelanguage.Label}
	default:
		err = &NotSingularError{wakatimelanguage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wlq *WakatimeLanguageQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := wlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WakatimeLanguages.
func (wlq *WakatimeLanguageQuery) All(ctx context.Context) ([]*WakatimeLanguage, error) {
	ctx = setContextOp(ctx, wlq.ctx, "All")
	if err := wlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WakatimeLanguage, *WakatimeLanguageQuery]()
	return withInterceptors[[]*WakatimeLanguage](ctx, wlq, qr, wlq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wlq *WakatimeLanguageQuery) AllX(ctx context.Context) []*WakatimeLanguage {
	nodes, err := wlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WakatimeLanguage IDs.
func (wlq *WakatimeLanguageQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if wlq.ctx.Unique == nil && wlq.path != nil {
		wlq.Unique(true)
	}
	ctx = setContextOp(ctx, wlq.ctx, "IDs")
	if err = wlq.Select(wakatimelanguage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wlq *WakatimeLanguageQuery) IDsX(ctx context.Context) []int64 {
	ids, err := wlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wlq *WakatimeLanguageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wlq.ctx, "Count")
	if err := wlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wlq, querierCount[*WakatimeLanguageQuery](), wlq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wlq *WakatimeLanguageQuery) CountX(ctx context.Context) int {
	count, err := wlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wlq *WakatimeLanguageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wlq.ctx, "Exist")
	switch _, err := wlq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("models: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wlq *WakatimeLanguageQuery) ExistX(ctx context.Context) bool {
	exist, err := wlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WakatimeLanguageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wlq *WakatimeLanguageQuery) Clone() *WakatimeLanguageQuery {
	if wlq == nil {
		return nil
	}
	return &WakatimeLanguageQuery{
		config:     wlq.config,
		ctx:        wlq.ctx.Clone(),
		order:      append([]wakatimelanguage.OrderOption{}, wlq.order...),
		inters:     append([]Interceptor{}, wlq.inters...),
		predicates: append([]predicate.WakatimeLanguage{}, wlq.predicates...),
		// clone intermediate query.
		sql:  wlq.sql.Clone(),
		path: wlq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (wlq *WakatimeLanguageQuery) GroupBy(field string, fields ...string) *WakatimeLanguageGroupBy {
	wlq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WakatimeLanguageGroupBy{build: wlq}
	grbuild.flds = &wlq.ctx.Fields
	grbuild.label = wakatimelanguage.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (wlq *WakatimeLanguageQuery) Select(fields ...string) *WakatimeLanguageSelect {
	wlq.ctx.Fields = append(wlq.ctx.Fields, fields...)
	sbuild := &WakatimeLanguageSelect{WakatimeLanguageQuery: wlq}
	sbuild.label = wakatimelanguage.Label
	sbuild.flds, sbuild.scan = &wlq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WakatimeLanguageSelect configured with the given aggregations.
func (wlq *WakatimeLanguageQuery) Aggregate(fns ...AggregateFunc) *WakatimeLanguageSelect {
	return wlq.Select().Aggregate(fns...)
}

func (wlq *WakatimeLanguageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wlq.inters {
		if inter == nil {
			return fmt.Errorf("models: uninitialized interceptor (forgotten import models/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wlq); err != nil {
				return err
			}
		}
	}
	for _, f := range wlq.ctx.Fields {
		if !wakatimelanguage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if wlq.path != nil {
		prev, err := wlq.path(ctx)
		if err != nil {
			return err
		}
		wlq.sql = prev
	}
	return nil
}

func (wlq *WakatimeLanguageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WakatimeLanguage, error) {
	var (
		nodes = []*WakatimeLanguage{}
		_spec = wlq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WakatimeLanguage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WakatimeLanguage{config: wlq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(wlq.modifiers) > 0 {
		_spec.Modifiers = wlq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range wlq.loadTotal {
		if err := wlq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wlq *WakatimeLanguageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wlq.querySpec()
	if len(wlq.modifiers) > 0 {
		_spec.Modifiers = wlq.modifiers
	}
	_spec.Node.Columns = wlq.ctx.Fields
	if len(wlq.ctx.Fields) > 0 {
		_spec.Unique = wlq.ctx.Unique != nil && *wlq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wlq.driver, _spec)
}

func (wlq *WakatimeLanguageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(wakatimelanguage.Table, wakatimelanguage.Columns, sqlgraph.NewFieldSpec(wakatimelanguage.FieldID, field.TypeInt64))
	_spec.From = wlq.sql
	if unique := wlq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wlq.path != nil {
		_spec.Unique = true
	}
	if fields := wlq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wakatimelanguage.FieldID)
		for i := range fields {
			if fields[i] != wakatimelanguage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wlq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wlq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wlq *WakatimeLanguageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wlq.driver.Dialect())
	t1 := builder.Table(wakatimelanguage.Table)
	columns := wlq.ctx.Fields
	if len(columns) == 0 {
		columns = wakatimelanguage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wlq.sql != nil {
		selector = wlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wlq.ctx.Unique != nil && *wlq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wlq.predicates {
		p(selector)
	}
	for _, p := range wlq.order {
		p(selector)
	}
	if offset := wlq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wlq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WakatimeLanguageGroupBy is the group-by builder for WakatimeLanguage entities.
type WakatimeLanguageGroupBy struct {
	selector
	build *WakatimeLanguageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wlgb *WakatimeLanguageGroupBy) Aggregate(fns ...AggregateFunc) *WakatimeLanguageGroupBy {
	wlgb.fns = append(wlgb.fns, fns...)
	return wlgb
}

// Scan applies the selector query and scans the result into the given value.
func (wlgb *WakatimeLanguageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wlgb.build.ctx, "GroupBy")
	if err := wlgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WakatimeLanguageQuery, *WakatimeLanguageGroupBy](ctx, wlgb.build, wlgb, wlgb.build.inters, v)
}

func (wlgb *WakatimeLanguageGroupBy) sqlScan(ctx context.Context, root *WakatimeLanguageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wlgb.fns))
	for _, fn := range wlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wlgb.flds)+len(wlgb.fns))
		for _, f := range *wlgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wlgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wlgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WakatimeLanguageSelect is the builder for selecting fields of WakatimeLanguage entities.
type WakatimeLanguageSelect struct {
	*WakatimeLanguageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wls *WakatimeLanguageSelect) Aggregate(fns ...AggregateFunc) *WakatimeLanguageSelect {
	wls.fns = append(wls.fns, fns...)
	return wls
}

// Scan applies the selector query and scans the result into the given value.
func (wls *WakatimeLanguageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wls.ctx, "Select")
	if err := wls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WakatimeLanguageQuery, *WakatimeLanguageSelect](ctx, wls.WakatimeLanguageQuery, wls, wls.inters, v)
}

func (wls *WakatimeLanguageSelect) sqlScan(ctx context.Context, root *WakatimeLanguageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wls.fns))
	for _, fn := range wls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
