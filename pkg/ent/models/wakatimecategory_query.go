// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/wakatimecategory"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WakatimeCategoryQuery is the builder for querying WakatimeCategory entities.
type WakatimeCategoryQuery struct {
	config
	ctx        *QueryContext
	order      []wakatimecategory.OrderOption
	inters     []Interceptor
	predicates []predicate.WakatimeCategory
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*WakatimeCategory) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WakatimeCategoryQuery builder.
func (wcq *WakatimeCategoryQuery) Where(ps ...predicate.WakatimeCategory) *WakatimeCategoryQuery {
	wcq.predicates = append(wcq.predicates, ps...)
	return wcq
}

// Limit the number of records to be returned by this query.
func (wcq *WakatimeCategoryQuery) Limit(limit int) *WakatimeCategoryQuery {
	wcq.ctx.Limit = &limit
	return wcq
}

// Offset to start from.
func (wcq *WakatimeCategoryQuery) Offset(offset int) *WakatimeCategoryQuery {
	wcq.ctx.Offset = &offset
	return wcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wcq *WakatimeCategoryQuery) Unique(unique bool) *WakatimeCategoryQuery {
	wcq.ctx.Unique = &unique
	return wcq
}

// Order specifies how the records should be ordered.
func (wcq *WakatimeCategoryQuery) Order(o ...wakatimecategory.OrderOption) *WakatimeCategoryQuery {
	wcq.order = append(wcq.order, o...)
	return wcq
}

// First returns the first WakatimeCategory entity from the query.
// Returns a *NotFoundError when no WakatimeCategory was found.
func (wcq *WakatimeCategoryQuery) First(ctx context.Context) (*WakatimeCategory, error) {
	nodes, err := wcq.Limit(1).All(setContextOp(ctx, wcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{wakatimecategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wcq *WakatimeCategoryQuery) FirstX(ctx context.Context) *WakatimeCategory {
	node, err := wcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WakatimeCategory ID from the query.
// Returns a *NotFoundError when no WakatimeCategory ID was found.
func (wcq *WakatimeCategoryQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = wcq.Limit(1).IDs(setContextOp(ctx, wcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{wakatimecategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wcq *WakatimeCategoryQuery) FirstIDX(ctx context.Context) int64 {
	id, err := wcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WakatimeCategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WakatimeCategory entity is found.
// Returns a *NotFoundError when no WakatimeCategory entities are found.
func (wcq *WakatimeCategoryQuery) Only(ctx context.Context) (*WakatimeCategory, error) {
	nodes, err := wcq.Limit(2).All(setContextOp(ctx, wcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{wakatimecategory.Label}
	default:
		return nil, &NotSingularError{wakatimecategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wcq *WakatimeCategoryQuery) OnlyX(ctx context.Context) *WakatimeCategory {
	node, err := wcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WakatimeCategory ID in the query.
// Returns a *NotSingularError when more than one WakatimeCategory ID is found.
// Returns a *NotFoundError when no entities are found.
func (wcq *WakatimeCategoryQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = wcq.Limit(2).IDs(setContextOp(ctx, wcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{wakatimecategory.Label}
	default:
		err = &NotSingularError{wakatimecategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wcq *WakatimeCategoryQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := wcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WakatimeCategories.
func (wcq *WakatimeCategoryQuery) All(ctx context.Context) ([]*WakatimeCategory, error) {
	ctx = setContextOp(ctx, wcq.ctx, "All")
	if err := wcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WakatimeCategory, *WakatimeCategoryQuery]()
	return withInterceptors[[]*WakatimeCategory](ctx, wcq, qr, wcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wcq *WakatimeCategoryQuery) AllX(ctx context.Context) []*WakatimeCategory {
	nodes, err := wcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WakatimeCategory IDs.
func (wcq *WakatimeCategoryQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if wcq.ctx.Unique == nil && wcq.path != nil {
		wcq.Unique(true)
	}
	ctx = setContextOp(ctx, wcq.ctx, "IDs")
	if err = wcq.Select(wakatimecategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wcq *WakatimeCategoryQuery) IDsX(ctx context.Context) []int64 {
	ids, err := wcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wcq *WakatimeCategoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wcq.ctx, "Count")
	if err := wcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wcq, querierCount[*WakatimeCategoryQuery](), wcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wcq *WakatimeCategoryQuery) CountX(ctx context.Context) int {
	count, err := wcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wcq *WakatimeCategoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wcq.ctx, "Exist")
	switch _, err := wcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("models: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wcq *WakatimeCategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := wcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WakatimeCategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wcq *WakatimeCategoryQuery) Clone() *WakatimeCategoryQuery {
	if wcq == nil {
		return nil
	}
	return &WakatimeCategoryQuery{
		config:     wcq.config,
		ctx:        wcq.ctx.Clone(),
		order:      append([]wakatimecategory.OrderOption{}, wcq.order...),
		inters:     append([]Interceptor{}, wcq.inters...),
		predicates: append([]predicate.WakatimeCategory{}, wcq.predicates...),
		// clone intermediate query.
		sql:  wcq.sql.Clone(),
		path: wcq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		WakatimeID int64 `json:"wakatime_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WakatimeCategory.Query().
//		GroupBy(wakatimecategory.FieldWakatimeID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
func (wcq *WakatimeCategoryQuery) GroupBy(field string, fields ...string) *WakatimeCategoryGroupBy {
	wcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WakatimeCategoryGroupBy{build: wcq}
	grbuild.flds = &wcq.ctx.Fields
	grbuild.label = wakatimecategory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		WakatimeID int64 `json:"wakatime_id,omitempty"`
//	}
//
//	client.WakatimeCategory.Query().
//		Select(wakatimecategory.FieldWakatimeID).
//		Scan(ctx, &v)
func (wcq *WakatimeCategoryQuery) Select(fields ...string) *WakatimeCategorySelect {
	wcq.ctx.Fields = append(wcq.ctx.Fields, fields...)
	sbuild := &WakatimeCategorySelect{WakatimeCategoryQuery: wcq}
	sbuild.label = wakatimecategory.Label
	sbuild.flds, sbuild.scan = &wcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WakatimeCategorySelect configured with the given aggregations.
func (wcq *WakatimeCategoryQuery) Aggregate(fns ...AggregateFunc) *WakatimeCategorySelect {
	return wcq.Select().Aggregate(fns...)
}

func (wcq *WakatimeCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wcq.inters {
		if inter == nil {
			return fmt.Errorf("models: uninitialized interceptor (forgotten import models/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wcq); err != nil {
				return err
			}
		}
	}
	for _, f := range wcq.ctx.Fields {
		if !wakatimecategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if wcq.path != nil {
		prev, err := wcq.path(ctx)
		if err != nil {
			return err
		}
		wcq.sql = prev
	}
	return nil
}

func (wcq *WakatimeCategoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WakatimeCategory, error) {
	var (
		nodes = []*WakatimeCategory{}
		_spec = wcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WakatimeCategory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WakatimeCategory{config: wcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(wcq.modifiers) > 0 {
		_spec.Modifiers = wcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range wcq.loadTotal {
		if err := wcq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wcq *WakatimeCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wcq.querySpec()
	if len(wcq.modifiers) > 0 {
		_spec.Modifiers = wcq.modifiers
	}
	_spec.Node.Columns = wcq.ctx.Fields
	if len(wcq.ctx.Fields) > 0 {
		_spec.Unique = wcq.ctx.Unique != nil && *wcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wcq.driver, _spec)
}

func (wcq *WakatimeCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(wakatimecategory.Table, wakatimecategory.Columns, sqlgraph.NewFieldSpec(wakatimecategory.FieldID, field.TypeInt64))
	_spec.From = wcq.sql
	if unique := wcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wcq.path != nil {
		_spec.Unique = true
	}
	if fields := wcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wakatimecategory.FieldID)
		for i := range fields {
			if fields[i] != wakatimecategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wcq *WakatimeCategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wcq.driver.Dialect())
	t1 := builder.Table(wakatimecategory.Table)
	columns := wcq.ctx.Fields
	if len(columns) == 0 {
		columns = wakatimecategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wcq.sql != nil {
		selector = wcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wcq.ctx.Unique != nil && *wcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wcq.predicates {
		p(selector)
	}
	for _, p := range wcq.order {
		p(selector)
	}
	if offset := wcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WakatimeCategoryGroupBy is the group-by builder for WakatimeCategory entities.
type WakatimeCategoryGroupBy struct {
	selector
	build *WakatimeCategoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wcgb *WakatimeCategoryGroupBy) Aggregate(fns ...AggregateFunc) *WakatimeCategoryGroupBy {
	wcgb.fns = append(wcgb.fns, fns...)
	return wcgb
}

// Scan applies the selector query and scans the result into the given value.
func (wcgb *WakatimeCategoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wcgb.build.ctx, "GroupBy")
	if err := wcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WakatimeCategoryQuery, *WakatimeCategoryGroupBy](ctx, wcgb.build, wcgb, wcgb.build.inters, v)
}

func (wcgb *WakatimeCategoryGroupBy) sqlScan(ctx context.Context, root *WakatimeCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wcgb.fns))
	for _, fn := range wcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wcgb.flds)+len(wcgb.fns))
		for _, f := range *wcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WakatimeCategorySelect is the builder for selecting fields of WakatimeCategory entities.
type WakatimeCategorySelect struct {
	*WakatimeCategoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wcs *WakatimeCategorySelect) Aggregate(fns ...AggregateFunc) *WakatimeCategorySelect {
	wcs.fns = append(wcs.fns, fns...)
	return wcs
}

// Scan applies the selector query and scans the result into the given value.
func (wcs *WakatimeCategorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wcs.ctx, "Select")
	if err := wcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WakatimeCategoryQuery, *WakatimeCategorySelect](ctx, wcs.WakatimeCategoryQuery, wcs, wcs.inters, v)
}

func (wcs *WakatimeCategorySelect) sqlScan(ctx context.Context, root *WakatimeCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wcs.fns))
	for _, fn := range wcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
