// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/webauthncredential"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WebAuthnCredentialQuery is the builder for querying WebAuthnCredential entities.
type WebAuthnCredentialQuery struct {
	config
	ctx        *QueryContext
	order      []webauthncredential.OrderOption
	inters     []Interceptor
	predicates []predicate.WebAuthnCredential
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*WebAuthnCredential) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WebAuthnCredentialQuery builder.
func (wacq *WebAuthnCredentialQuery) Where(ps ...predicate.WebAuthnCredential) *WebAuthnCredentialQuery {
	wacq.predicates = append(wacq.predicates, ps...)
	return wacq
}

// Limit the number of records to be returned by this query.
func (wacq *WebAuthnCredentialQuery) Limit(limit int) *WebAuthnCredentialQuery {
	wacq.ctx.Limit = &limit
	return wacq
}

// Offset to start from.
func (wacq *WebAuthnCredentialQuery) Offset(offset int) *WebAuthnCredentialQuery {
	wacq.ctx.Offset = &offset
	return wacq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wacq *WebAuthnCredentialQuery) Unique(unique bool) *WebAuthnCredentialQuery {
	wacq.ctx.Unique = &unique
	return wacq
}

// Order specifies how the records should be ordered.
func (wacq *WebAuthnCredentialQuery) Order(o ...webauthncredential.OrderOption) *WebAuthnCredentialQuery {
	wacq.order = append(wacq.order, o...)
	return wacq
}

// First returns the first WebAuthnCredential entity from the query.
// Returns a *NotFoundError when no WebAuthnCredential was found.
func (wacq *WebAuthnCredentialQuery) First(ctx context.Context) (*WebAuthnCredential, error) {
	nodes, err := wacq.Limit(1).All(setContextOp(ctx, wacq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{webauthncredential.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wacq *WebAuthnCredentialQuery) FirstX(ctx context.Context) *WebAuthnCredential {
	node, err := wacq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WebAuthnCredential ID from the query.
// Returns a *NotFoundError when no WebAuthnCredential ID was found.
func (wacq *WebAuthnCredentialQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = wacq.Limit(1).IDs(setContextOp(ctx, wacq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{webauthncredential.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wacq *WebAuthnCredentialQuery) FirstIDX(ctx context.Context) int64 {
	id, err := wacq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WebAuthnCredential entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WebAuthnCredential entity is found.
// Returns a *NotFoundError when no WebAuthnCredential entities are found.
func (wacq *WebAuthnCredentialQuery) Only(ctx context.Context) (*WebAuthnCredential, error) {
	nodes, err := wacq.Limit(2).All(setContextOp(ctx, wacq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{webauthncredential.Label}
	default:
		return nil, &NotSingularError{webauthncredential.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wacq *WebAuthnCredentialQuery) OnlyX(ctx context.Context) *WebAuthnCredential {
	node, err := wacq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WebAuthnCredential ID in the query.
// Returns a *NotSingularError when more than one WebAuthnCredential ID is found.
// Returns a *NotFoundError when no entities are found.
func (wacq *WebAuthnCredentialQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = wacq.Limit(2).IDs(setContextOp(ctx, wacq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{webauthncredential.Label}
	default:
		err = &NotSingularError{webauthncredential.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wacq *WebAuthnCredentialQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := wacq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WebAuthnCredentials.
func (wacq *WebAuthnCredentialQuery) All(ctx context.Context) ([]*WebAuthnCredential, error) {
	ctx = setContextOp(ctx, wacq.ctx, "All")
	if err := wacq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WebAuthnCredential, *WebAuthnCredentialQuery]()
	return withInterceptors[[]*WebAuthnCredential](ctx, wacq, qr, wacq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wacq *WebAuthnCredentialQuery) AllX(ctx context.Context) []*WebAuthnCredential {
	nodes, err := wacq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WebAuthnCredential IDs.
func (wacq *WebAuthnCredentialQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if wacq.ctx.Unique == nil && wacq.path != nil {
		wacq.Unique(true)
	}
	ctx = setContextOp(ctx, wacq.ctx, "IDs")
	if err = wacq.Select(webauthncredential.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wacq *WebAuthnCredentialQuery) IDsX(ctx context.Context) []int64 {
	ids, err := wacq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wacq *WebAuthnCredentialQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wacq.ctx, "Count")
	if err := wacq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wacq, querierCount[*WebAuthnCredentialQuery](), wacq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wacq *WebAuthnCredentialQuery) CountX(ctx context.Context) int {
	count, err := wacq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wacq *WebAuthnCredentialQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wacq.ctx, "Exist")
	switch _, err := wacq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("models: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wacq *WebAuthnCredentialQuery) ExistX(ctx context.Context) bool {
	exist, err := wacq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WebAuthnCredentialQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wacq *WebAuthnCredentialQuery) Clone() *WebAuthnCredentialQuery {
	if wacq == nil {
		return nil
	}
	return &WebAuthnCredentialQuery{
		config:     wacq.config,
		ctx:        wacq.ctx.Clone(),
		order:      append([]webauthncredential.OrderOption{}, wacq.order...),
		inters:     append([]Interceptor{}, wacq.inters...),
		predicates: append([]predicate.WebAuthnCredential{}, wacq.predicates...),
		// clone intermediate query.
		sql:  wacq.sql.Clone(),
		path: wacq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WebAuthnCredential.Query().
//		GroupBy(webauthncredential.FieldName).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
func (wacq *WebAuthnCredentialQuery) GroupBy(field string, fields ...string) *WebAuthnCredentialGroupBy {
	wacq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WebAuthnCredentialGroupBy{build: wacq}
	grbuild.flds = &wacq.ctx.Fields
	grbuild.label = webauthncredential.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.WebAuthnCredential.Query().
//		Select(webauthncredential.FieldName).
//		Scan(ctx, &v)
func (wacq *WebAuthnCredentialQuery) Select(fields ...string) *WebAuthnCredentialSelect {
	wacq.ctx.Fields = append(wacq.ctx.Fields, fields...)
	sbuild := &WebAuthnCredentialSelect{WebAuthnCredentialQuery: wacq}
	sbuild.label = webauthncredential.Label
	sbuild.flds, sbuild.scan = &wacq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WebAuthnCredentialSelect configured with the given aggregations.
func (wacq *WebAuthnCredentialQuery) Aggregate(fns ...AggregateFunc) *WebAuthnCredentialSelect {
	return wacq.Select().Aggregate(fns...)
}

func (wacq *WebAuthnCredentialQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wacq.inters {
		if inter == nil {
			return fmt.Errorf("models: uninitialized interceptor (forgotten import models/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wacq); err != nil {
				return err
			}
		}
	}
	for _, f := range wacq.ctx.Fields {
		if !webauthncredential.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if wacq.path != nil {
		prev, err := wacq.path(ctx)
		if err != nil {
			return err
		}
		wacq.sql = prev
	}
	return nil
}

func (wacq *WebAuthnCredentialQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WebAuthnCredential, error) {
	var (
		nodes = []*WebAuthnCredential{}
		_spec = wacq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WebAuthnCredential).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WebAuthnCredential{config: wacq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(wacq.modifiers) > 0 {
		_spec.Modifiers = wacq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wacq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range wacq.loadTotal {
		if err := wacq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wacq *WebAuthnCredentialQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wacq.querySpec()
	if len(wacq.modifiers) > 0 {
		_spec.Modifiers = wacq.modifiers
	}
	_spec.Node.Columns = wacq.ctx.Fields
	if len(wacq.ctx.Fields) > 0 {
		_spec.Unique = wacq.ctx.Unique != nil && *wacq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wacq.driver, _spec)
}

func (wacq *WebAuthnCredentialQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(webauthncredential.Table, webauthncredential.Columns, sqlgraph.NewFieldSpec(webauthncredential.FieldID, field.TypeInt64))
	_spec.From = wacq.sql
	if unique := wacq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wacq.path != nil {
		_spec.Unique = true
	}
	if fields := wacq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, webauthncredential.FieldID)
		for i := range fields {
			if fields[i] != webauthncredential.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wacq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wacq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wacq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wacq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wacq *WebAuthnCredentialQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wacq.driver.Dialect())
	t1 := builder.Table(webauthncredential.Table)
	columns := wacq.ctx.Fields
	if len(columns) == 0 {
		columns = webauthncredential.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wacq.sql != nil {
		selector = wacq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wacq.ctx.Unique != nil && *wacq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wacq.predicates {
		p(selector)
	}
	for _, p := range wacq.order {
		p(selector)
	}
	if offset := wacq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wacq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WebAuthnCredentialGroupBy is the group-by builder for WebAuthnCredential entities.
type WebAuthnCredentialGroupBy struct {
	selector
	build *WebAuthnCredentialQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wacgb *WebAuthnCredentialGroupBy) Aggregate(fns ...AggregateFunc) *WebAuthnCredentialGroupBy {
	wacgb.fns = append(wacgb.fns, fns...)
	return wacgb
}

// Scan applies the selector query and scans the result into the given value.
func (wacgb *WebAuthnCredentialGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wacgb.build.ctx, "GroupBy")
	if err := wacgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WebAuthnCredentialQuery, *WebAuthnCredentialGroupBy](ctx, wacgb.build, wacgb, wacgb.build.inters, v)
}

func (wacgb *WebAuthnCredentialGroupBy) sqlScan(ctx context.Context, root *WebAuthnCredentialQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wacgb.fns))
	for _, fn := range wacgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wacgb.flds)+len(wacgb.fns))
		for _, f := range *wacgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wacgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wacgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WebAuthnCredentialSelect is the builder for selecting fields of WebAuthnCredential entities.
type WebAuthnCredentialSelect struct {
	*WebAuthnCredentialQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wacs *WebAuthnCredentialSelect) Aggregate(fns ...AggregateFunc) *WebAuthnCredentialSelect {
	wacs.fns = append(wacs.fns, fns...)
	return wacs
}

// Scan applies the selector query and scans the result into the given value.
func (wacs *WebAuthnCredentialSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wacs.ctx, "Select")
	if err := wacs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WebAuthnCredentialQuery, *WebAuthnCredentialSelect](ctx, wacs.WebAuthnCredentialQuery, wacs, wacs.inters, v)
}

func (wacs *WebAuthnCredentialSelect) sqlScan(ctx context.Context, root *WebAuthnCredentialQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wacs.fns))
	for _, fn := range wacs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wacs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wacs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
