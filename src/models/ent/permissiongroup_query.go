// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"service-api/src/models/ent/permissiongroup"
	"service-api/src/models/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PermissionGroupQuery is the builder for querying PermissionGroup entities.
type PermissionGroupQuery struct {
	config
	ctx        *QueryContext
	order      []permissiongroup.OrderOption
	inters     []Interceptor
	predicates []predicate.PermissionGroup
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PermissionGroupQuery builder.
func (pgq *PermissionGroupQuery) Where(ps ...predicate.PermissionGroup) *PermissionGroupQuery {
	pgq.predicates = append(pgq.predicates, ps...)
	return pgq
}

// Limit the number of records to be returned by this query.
func (pgq *PermissionGroupQuery) Limit(limit int) *PermissionGroupQuery {
	pgq.ctx.Limit = &limit
	return pgq
}

// Offset to start from.
func (pgq *PermissionGroupQuery) Offset(offset int) *PermissionGroupQuery {
	pgq.ctx.Offset = &offset
	return pgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pgq *PermissionGroupQuery) Unique(unique bool) *PermissionGroupQuery {
	pgq.ctx.Unique = &unique
	return pgq
}

// Order specifies how the records should be ordered.
func (pgq *PermissionGroupQuery) Order(o ...permissiongroup.OrderOption) *PermissionGroupQuery {
	pgq.order = append(pgq.order, o...)
	return pgq
}

// First returns the first PermissionGroup entity from the query.
// Returns a *NotFoundError when no PermissionGroup was found.
func (pgq *PermissionGroupQuery) First(ctx context.Context) (*PermissionGroup, error) {
	nodes, err := pgq.Limit(1).All(setContextOp(ctx, pgq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{permissiongroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pgq *PermissionGroupQuery) FirstX(ctx context.Context) *PermissionGroup {
	node, err := pgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PermissionGroup ID from the query.
// Returns a *NotFoundError when no PermissionGroup ID was found.
func (pgq *PermissionGroupQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pgq.Limit(1).IDs(setContextOp(ctx, pgq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{permissiongroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pgq *PermissionGroupQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PermissionGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PermissionGroup entity is found.
// Returns a *NotFoundError when no PermissionGroup entities are found.
func (pgq *PermissionGroupQuery) Only(ctx context.Context) (*PermissionGroup, error) {
	nodes, err := pgq.Limit(2).All(setContextOp(ctx, pgq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{permissiongroup.Label}
	default:
		return nil, &NotSingularError{permissiongroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pgq *PermissionGroupQuery) OnlyX(ctx context.Context) *PermissionGroup {
	node, err := pgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PermissionGroup ID in the query.
// Returns a *NotSingularError when more than one PermissionGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (pgq *PermissionGroupQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pgq.Limit(2).IDs(setContextOp(ctx, pgq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{permissiongroup.Label}
	default:
		err = &NotSingularError{permissiongroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pgq *PermissionGroupQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PermissionGroups.
func (pgq *PermissionGroupQuery) All(ctx context.Context) ([]*PermissionGroup, error) {
	ctx = setContextOp(ctx, pgq.ctx, "All")
	if err := pgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PermissionGroup, *PermissionGroupQuery]()
	return withInterceptors[[]*PermissionGroup](ctx, pgq, qr, pgq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pgq *PermissionGroupQuery) AllX(ctx context.Context) []*PermissionGroup {
	nodes, err := pgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PermissionGroup IDs.
func (pgq *PermissionGroupQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if pgq.ctx.Unique == nil && pgq.path != nil {
		pgq.Unique(true)
	}
	ctx = setContextOp(ctx, pgq.ctx, "IDs")
	if err = pgq.Select(permissiongroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pgq *PermissionGroupQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pgq *PermissionGroupQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pgq.ctx, "Count")
	if err := pgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pgq, querierCount[*PermissionGroupQuery](), pgq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pgq *PermissionGroupQuery) CountX(ctx context.Context) int {
	count, err := pgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pgq *PermissionGroupQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pgq.ctx, "Exist")
	switch _, err := pgq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pgq *PermissionGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := pgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PermissionGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pgq *PermissionGroupQuery) Clone() *PermissionGroupQuery {
	if pgq == nil {
		return nil
	}
	return &PermissionGroupQuery{
		config:     pgq.config,
		ctx:        pgq.ctx.Clone(),
		order:      append([]permissiongroup.OrderOption{}, pgq.order...),
		inters:     append([]Interceptor{}, pgq.inters...),
		predicates: append([]predicate.PermissionGroup{}, pgq.predicates...),
		// clone intermediate query.
		sql:  pgq.sql.Clone(),
		path: pgq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PermissionGroup.Query().
//		GroupBy(permissiongroup.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pgq *PermissionGroupQuery) GroupBy(field string, fields ...string) *PermissionGroupGroupBy {
	pgq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PermissionGroupGroupBy{build: pgq}
	grbuild.flds = &pgq.ctx.Fields
	grbuild.label = permissiongroup.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.PermissionGroup.Query().
//		Select(permissiongroup.FieldCreateTime).
//		Scan(ctx, &v)
func (pgq *PermissionGroupQuery) Select(fields ...string) *PermissionGroupSelect {
	pgq.ctx.Fields = append(pgq.ctx.Fields, fields...)
	sbuild := &PermissionGroupSelect{PermissionGroupQuery: pgq}
	sbuild.label = permissiongroup.Label
	sbuild.flds, sbuild.scan = &pgq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PermissionGroupSelect configured with the given aggregations.
func (pgq *PermissionGroupQuery) Aggregate(fns ...AggregateFunc) *PermissionGroupSelect {
	return pgq.Select().Aggregate(fns...)
}

func (pgq *PermissionGroupQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pgq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pgq); err != nil {
				return err
			}
		}
	}
	for _, f := range pgq.ctx.Fields {
		if !permissiongroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pgq.path != nil {
		prev, err := pgq.path(ctx)
		if err != nil {
			return err
		}
		pgq.sql = prev
	}
	return nil
}

func (pgq *PermissionGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PermissionGroup, error) {
	var (
		nodes = []*PermissionGroup{}
		_spec = pgq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PermissionGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PermissionGroup{config: pgq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pgq *PermissionGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pgq.querySpec()
	_spec.Node.Columns = pgq.ctx.Fields
	if len(pgq.ctx.Fields) > 0 {
		_spec.Unique = pgq.ctx.Unique != nil && *pgq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pgq.driver, _spec)
}

func (pgq *PermissionGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(permissiongroup.Table, permissiongroup.Columns, sqlgraph.NewFieldSpec(permissiongroup.FieldID, field.TypeUUID))
	_spec.From = pgq.sql
	if unique := pgq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pgq.path != nil {
		_spec.Unique = true
	}
	if fields := pgq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permissiongroup.FieldID)
		for i := range fields {
			if fields[i] != permissiongroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pgq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pgq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pgq *PermissionGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pgq.driver.Dialect())
	t1 := builder.Table(permissiongroup.Table)
	columns := pgq.ctx.Fields
	if len(columns) == 0 {
		columns = permissiongroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pgq.sql != nil {
		selector = pgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pgq.ctx.Unique != nil && *pgq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pgq.predicates {
		p(selector)
	}
	for _, p := range pgq.order {
		p(selector)
	}
	if offset := pgq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pgq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PermissionGroupGroupBy is the group-by builder for PermissionGroup entities.
type PermissionGroupGroupBy struct {
	selector
	build *PermissionGroupQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pggb *PermissionGroupGroupBy) Aggregate(fns ...AggregateFunc) *PermissionGroupGroupBy {
	pggb.fns = append(pggb.fns, fns...)
	return pggb
}

// Scan applies the selector query and scans the result into the given value.
func (pggb *PermissionGroupGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pggb.build.ctx, "GroupBy")
	if err := pggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PermissionGroupQuery, *PermissionGroupGroupBy](ctx, pggb.build, pggb, pggb.build.inters, v)
}

func (pggb *PermissionGroupGroupBy) sqlScan(ctx context.Context, root *PermissionGroupQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pggb.fns))
	for _, fn := range pggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pggb.flds)+len(pggb.fns))
		for _, f := range *pggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PermissionGroupSelect is the builder for selecting fields of PermissionGroup entities.
type PermissionGroupSelect struct {
	*PermissionGroupQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pgs *PermissionGroupSelect) Aggregate(fns ...AggregateFunc) *PermissionGroupSelect {
	pgs.fns = append(pgs.fns, fns...)
	return pgs
}

// Scan applies the selector query and scans the result into the given value.
func (pgs *PermissionGroupSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgs.ctx, "Select")
	if err := pgs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PermissionGroupQuery, *PermissionGroupSelect](ctx, pgs.PermissionGroupQuery, pgs, pgs.inters, v)
}

func (pgs *PermissionGroupSelect) sqlScan(ctx context.Context, root *PermissionGroupQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pgs.fns))
	for _, fn := range pgs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pgs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
