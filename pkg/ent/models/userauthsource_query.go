// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/userauthsource"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserAuthSourceQuery is the builder for querying UserAuthSource entities.
type UserAuthSourceQuery struct {
	config
	ctx        *QueryContext
	order      []userauthsource.OrderOption
	inters     []Interceptor
	predicates []predicate.UserAuthSource
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserAuthSourceQuery builder.
func (uasq *UserAuthSourceQuery) Where(ps ...predicate.UserAuthSource) *UserAuthSourceQuery {
	uasq.predicates = append(uasq.predicates, ps...)
	return uasq
}

// Limit the number of records to be returned by this query.
func (uasq *UserAuthSourceQuery) Limit(limit int) *UserAuthSourceQuery {
	uasq.ctx.Limit = &limit
	return uasq
}

// Offset to start from.
func (uasq *UserAuthSourceQuery) Offset(offset int) *UserAuthSourceQuery {
	uasq.ctx.Offset = &offset
	return uasq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uasq *UserAuthSourceQuery) Unique(unique bool) *UserAuthSourceQuery {
	uasq.ctx.Unique = &unique
	return uasq
}

// Order specifies how the records should be ordered.
func (uasq *UserAuthSourceQuery) Order(o ...userauthsource.OrderOption) *UserAuthSourceQuery {
	uasq.order = append(uasq.order, o...)
	return uasq
}

// First returns the first UserAuthSource entity from the query.
// Returns a *NotFoundError when no UserAuthSource was found.
func (uasq *UserAuthSourceQuery) First(ctx context.Context) (*UserAuthSource, error) {
	nodes, err := uasq.Limit(1).All(setContextOp(ctx, uasq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userauthsource.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uasq *UserAuthSourceQuery) FirstX(ctx context.Context) *UserAuthSource {
	node, err := uasq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserAuthSource ID from the query.
// Returns a *NotFoundError when no UserAuthSource ID was found.
func (uasq *UserAuthSourceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uasq.Limit(1).IDs(setContextOp(ctx, uasq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userauthsource.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uasq *UserAuthSourceQuery) FirstIDX(ctx context.Context) int {
	id, err := uasq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserAuthSource entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserAuthSource entity is found.
// Returns a *NotFoundError when no UserAuthSource entities are found.
func (uasq *UserAuthSourceQuery) Only(ctx context.Context) (*UserAuthSource, error) {
	nodes, err := uasq.Limit(2).All(setContextOp(ctx, uasq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userauthsource.Label}
	default:
		return nil, &NotSingularError{userauthsource.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uasq *UserAuthSourceQuery) OnlyX(ctx context.Context) *UserAuthSource {
	node, err := uasq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserAuthSource ID in the query.
// Returns a *NotSingularError when more than one UserAuthSource ID is found.
// Returns a *NotFoundError when no entities are found.
func (uasq *UserAuthSourceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uasq.Limit(2).IDs(setContextOp(ctx, uasq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userauthsource.Label}
	default:
		err = &NotSingularError{userauthsource.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uasq *UserAuthSourceQuery) OnlyIDX(ctx context.Context) int {
	id, err := uasq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserAuthSources.
func (uasq *UserAuthSourceQuery) All(ctx context.Context) ([]*UserAuthSource, error) {
	ctx = setContextOp(ctx, uasq.ctx, "All")
	if err := uasq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserAuthSource, *UserAuthSourceQuery]()
	return withInterceptors[[]*UserAuthSource](ctx, uasq, qr, uasq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uasq *UserAuthSourceQuery) AllX(ctx context.Context) []*UserAuthSource {
	nodes, err := uasq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserAuthSource IDs.
func (uasq *UserAuthSourceQuery) IDs(ctx context.Context) (ids []int, err error) {
	if uasq.ctx.Unique == nil && uasq.path != nil {
		uasq.Unique(true)
	}
	ctx = setContextOp(ctx, uasq.ctx, "IDs")
	if err = uasq.Select(userauthsource.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uasq *UserAuthSourceQuery) IDsX(ctx context.Context) []int {
	ids, err := uasq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uasq *UserAuthSourceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uasq.ctx, "Count")
	if err := uasq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uasq, querierCount[*UserAuthSourceQuery](), uasq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uasq *UserAuthSourceQuery) CountX(ctx context.Context) int {
	count, err := uasq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uasq *UserAuthSourceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uasq.ctx, "Exist")
	switch _, err := uasq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("models: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uasq *UserAuthSourceQuery) ExistX(ctx context.Context) bool {
	exist, err := uasq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserAuthSourceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uasq *UserAuthSourceQuery) Clone() *UserAuthSourceQuery {
	if uasq == nil {
		return nil
	}
	return &UserAuthSourceQuery{
		config:     uasq.config,
		ctx:        uasq.ctx.Clone(),
		order:      append([]userauthsource.OrderOption{}, uasq.order...),
		inters:     append([]Interceptor{}, uasq.inters...),
		predicates: append([]predicate.UserAuthSource{}, uasq.predicates...),
		// clone intermediate query.
		sql:  uasq.sql.Clone(),
		path: uasq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserAuthSource.Query().
//		GroupBy(userauthsource.FieldUserID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
func (uasq *UserAuthSourceQuery) GroupBy(field string, fields ...string) *UserAuthSourceGroupBy {
	uasq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserAuthSourceGroupBy{build: uasq}
	grbuild.flds = &uasq.ctx.Fields
	grbuild.label = userauthsource.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//	}
//
//	client.UserAuthSource.Query().
//		Select(userauthsource.FieldUserID).
//		Scan(ctx, &v)
func (uasq *UserAuthSourceQuery) Select(fields ...string) *UserAuthSourceSelect {
	uasq.ctx.Fields = append(uasq.ctx.Fields, fields...)
	sbuild := &UserAuthSourceSelect{UserAuthSourceQuery: uasq}
	sbuild.label = userauthsource.Label
	sbuild.flds, sbuild.scan = &uasq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserAuthSourceSelect configured with the given aggregations.
func (uasq *UserAuthSourceQuery) Aggregate(fns ...AggregateFunc) *UserAuthSourceSelect {
	return uasq.Select().Aggregate(fns...)
}

func (uasq *UserAuthSourceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uasq.inters {
		if inter == nil {
			return fmt.Errorf("models: uninitialized interceptor (forgotten import models/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uasq); err != nil {
				return err
			}
		}
	}
	for _, f := range uasq.ctx.Fields {
		if !userauthsource.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
		}
	}
	if uasq.path != nil {
		prev, err := uasq.path(ctx)
		if err != nil {
			return err
		}
		uasq.sql = prev
	}
	return nil
}

func (uasq *UserAuthSourceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserAuthSource, error) {
	var (
		nodes = []*UserAuthSource{}
		_spec = uasq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserAuthSource).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserAuthSource{config: uasq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uasq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (uasq *UserAuthSourceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uasq.querySpec()
	_spec.Node.Columns = uasq.ctx.Fields
	if len(uasq.ctx.Fields) > 0 {
		_spec.Unique = uasq.ctx.Unique != nil && *uasq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uasq.driver, _spec)
}

func (uasq *UserAuthSourceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userauthsource.Table, userauthsource.Columns, sqlgraph.NewFieldSpec(userauthsource.FieldID, field.TypeInt))
	_spec.From = uasq.sql
	if unique := uasq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uasq.path != nil {
		_spec.Unique = true
	}
	if fields := uasq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userauthsource.FieldID)
		for i := range fields {
			if fields[i] != userauthsource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uasq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uasq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uasq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uasq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uasq *UserAuthSourceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uasq.driver.Dialect())
	t1 := builder.Table(userauthsource.Table)
	columns := uasq.ctx.Fields
	if len(columns) == 0 {
		columns = userauthsource.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uasq.sql != nil {
		selector = uasq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uasq.ctx.Unique != nil && *uasq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uasq.predicates {
		p(selector)
	}
	for _, p := range uasq.order {
		p(selector)
	}
	if offset := uasq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uasq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserAuthSourceGroupBy is the group-by builder for UserAuthSource entities.
type UserAuthSourceGroupBy struct {
	selector
	build *UserAuthSourceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uasgb *UserAuthSourceGroupBy) Aggregate(fns ...AggregateFunc) *UserAuthSourceGroupBy {
	uasgb.fns = append(uasgb.fns, fns...)
	return uasgb
}

// Scan applies the selector query and scans the result into the given value.
func (uasgb *UserAuthSourceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uasgb.build.ctx, "GroupBy")
	if err := uasgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserAuthSourceQuery, *UserAuthSourceGroupBy](ctx, uasgb.build, uasgb, uasgb.build.inters, v)
}

func (uasgb *UserAuthSourceGroupBy) sqlScan(ctx context.Context, root *UserAuthSourceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(uasgb.fns))
	for _, fn := range uasgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*uasgb.flds)+len(uasgb.fns))
		for _, f := range *uasgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*uasgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uasgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserAuthSourceSelect is the builder for selecting fields of UserAuthSource entities.
type UserAuthSourceSelect struct {
	*UserAuthSourceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uass *UserAuthSourceSelect) Aggregate(fns ...AggregateFunc) *UserAuthSourceSelect {
	uass.fns = append(uass.fns, fns...)
	return uass
}

// Scan applies the selector query and scans the result into the given value.
func (uass *UserAuthSourceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uass.ctx, "Select")
	if err := uass.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserAuthSourceQuery, *UserAuthSourceSelect](ctx, uass.UserAuthSourceQuery, uass, uass.inters, v)
}

func (uass *UserAuthSourceSelect) sqlScan(ctx context.Context, root *UserAuthSourceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uass.fns))
	for _, fn := range uass.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uass.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uass.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
