// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/wakatimecategory"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WakatimeCategoryUpdate is the builder for updating WakatimeCategory entities.
type WakatimeCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *WakatimeCategoryMutation
}

// Where appends a list predicates to the WakatimeCategoryUpdate builder.
func (wcu *WakatimeCategoryUpdate) Where(ps ...predicate.WakatimeCategory) *WakatimeCategoryUpdate {
	wcu.mutation.Where(ps...)
	return wcu
}

// SetUpdateTime sets the "update_time" field.
func (wcu *WakatimeCategoryUpdate) SetUpdateTime(t time.Time) *WakatimeCategoryUpdate {
	wcu.mutation.SetUpdateTime(t)
	return wcu
}

// SetWakatimeID sets the "wakatime_id" field.
func (wcu *WakatimeCategoryUpdate) SetWakatimeID(u uuid.UUID) *WakatimeCategoryUpdate {
	wcu.mutation.SetWakatimeID(u)
	return wcu
}

// SetNillableWakatimeID sets the "wakatime_id" field if the given value is not nil.
func (wcu *WakatimeCategoryUpdate) SetNillableWakatimeID(u *uuid.UUID) *WakatimeCategoryUpdate {
	if u != nil {
		wcu.SetWakatimeID(*u)
	}
	return wcu
}

// SetUserID sets the "user_id" field.
func (wcu *WakatimeCategoryUpdate) SetUserID(u uuid.UUID) *WakatimeCategoryUpdate {
	wcu.mutation.SetUserID(u)
	return wcu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wcu *WakatimeCategoryUpdate) SetNillableUserID(u *uuid.UUID) *WakatimeCategoryUpdate {
	if u != nil {
		wcu.SetUserID(*u)
	}
	return wcu
}

// SetName sets the "name" field.
func (wcu *WakatimeCategoryUpdate) SetName(s string) *WakatimeCategoryUpdate {
	wcu.mutation.SetName(s)
	return wcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (wcu *WakatimeCategoryUpdate) SetNillableName(s *string) *WakatimeCategoryUpdate {
	if s != nil {
		wcu.SetName(*s)
	}
	return wcu
}

// SetTotalSeconds sets the "total_seconds" field.
func (wcu *WakatimeCategoryUpdate) SetTotalSeconds(i int64) *WakatimeCategoryUpdate {
	wcu.mutation.ResetTotalSeconds()
	wcu.mutation.SetTotalSeconds(i)
	return wcu
}

// SetNillableTotalSeconds sets the "total_seconds" field if the given value is not nil.
func (wcu *WakatimeCategoryUpdate) SetNillableTotalSeconds(i *int64) *WakatimeCategoryUpdate {
	if i != nil {
		wcu.SetTotalSeconds(*i)
	}
	return wcu
}

// AddTotalSeconds adds i to the "total_seconds" field.
func (wcu *WakatimeCategoryUpdate) AddTotalSeconds(i int64) *WakatimeCategoryUpdate {
	wcu.mutation.AddTotalSeconds(i)
	return wcu
}

// Mutation returns the WakatimeCategoryMutation object of the builder.
func (wcu *WakatimeCategoryUpdate) Mutation() *WakatimeCategoryMutation {
	return wcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wcu *WakatimeCategoryUpdate) Save(ctx context.Context) (int, error) {
	wcu.defaults()
	return withHooks(ctx, wcu.sqlSave, wcu.mutation, wcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wcu *WakatimeCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := wcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wcu *WakatimeCategoryUpdate) Exec(ctx context.Context) error {
	_, err := wcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcu *WakatimeCategoryUpdate) ExecX(ctx context.Context) {
	if err := wcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wcu *WakatimeCategoryUpdate) defaults() {
	if _, ok := wcu.mutation.UpdateTime(); !ok {
		v := wakatimecategory.UpdateDefaultUpdateTime()
		wcu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wcu *WakatimeCategoryUpdate) check() error {
	if v, ok := wcu.mutation.Name(); ok {
		if err := wakatimecategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`models: validator failed for field "WakatimeCategory.name": %w`, err)}
		}
	}
	return nil
}

func (wcu *WakatimeCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(wakatimecategory.Table, wakatimecategory.Columns, sqlgraph.NewFieldSpec(wakatimecategory.FieldID, field.TypeInt64))
	if ps := wcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wcu.mutation.UpdateTime(); ok {
		_spec.SetField(wakatimecategory.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := wcu.mutation.WakatimeID(); ok {
		_spec.SetField(wakatimecategory.FieldWakatimeID, field.TypeUUID, value)
	}
	if value, ok := wcu.mutation.UserID(); ok {
		_spec.SetField(wakatimecategory.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := wcu.mutation.Name(); ok {
		_spec.SetField(wakatimecategory.FieldName, field.TypeString, value)
	}
	if value, ok := wcu.mutation.TotalSeconds(); ok {
		_spec.SetField(wakatimecategory.FieldTotalSeconds, field.TypeInt64, value)
	}
	if value, ok := wcu.mutation.AddedTotalSeconds(); ok {
		_spec.AddField(wakatimecategory.FieldTotalSeconds, field.TypeInt64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wakatimecategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wcu.mutation.done = true
	return n, nil
}

// WakatimeCategoryUpdateOne is the builder for updating a single WakatimeCategory entity.
type WakatimeCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WakatimeCategoryMutation
}

// SetUpdateTime sets the "update_time" field.
func (wcuo *WakatimeCategoryUpdateOne) SetUpdateTime(t time.Time) *WakatimeCategoryUpdateOne {
	wcuo.mutation.SetUpdateTime(t)
	return wcuo
}

// SetWakatimeID sets the "wakatime_id" field.
func (wcuo *WakatimeCategoryUpdateOne) SetWakatimeID(u uuid.UUID) *WakatimeCategoryUpdateOne {
	wcuo.mutation.SetWakatimeID(u)
	return wcuo
}

// SetNillableWakatimeID sets the "wakatime_id" field if the given value is not nil.
func (wcuo *WakatimeCategoryUpdateOne) SetNillableWakatimeID(u *uuid.UUID) *WakatimeCategoryUpdateOne {
	if u != nil {
		wcuo.SetWakatimeID(*u)
	}
	return wcuo
}

// SetUserID sets the "user_id" field.
func (wcuo *WakatimeCategoryUpdateOne) SetUserID(u uuid.UUID) *WakatimeCategoryUpdateOne {
	wcuo.mutation.SetUserID(u)
	return wcuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wcuo *WakatimeCategoryUpdateOne) SetNillableUserID(u *uuid.UUID) *WakatimeCategoryUpdateOne {
	if u != nil {
		wcuo.SetUserID(*u)
	}
	return wcuo
}

// SetName sets the "name" field.
func (wcuo *WakatimeCategoryUpdateOne) SetName(s string) *WakatimeCategoryUpdateOne {
	wcuo.mutation.SetName(s)
	return wcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (wcuo *WakatimeCategoryUpdateOne) SetNillableName(s *string) *WakatimeCategoryUpdateOne {
	if s != nil {
		wcuo.SetName(*s)
	}
	return wcuo
}

// SetTotalSeconds sets the "total_seconds" field.
func (wcuo *WakatimeCategoryUpdateOne) SetTotalSeconds(i int64) *WakatimeCategoryUpdateOne {
	wcuo.mutation.ResetTotalSeconds()
	wcuo.mutation.SetTotalSeconds(i)
	return wcuo
}

// SetNillableTotalSeconds sets the "total_seconds" field if the given value is not nil.
func (wcuo *WakatimeCategoryUpdateOne) SetNillableTotalSeconds(i *int64) *WakatimeCategoryUpdateOne {
	if i != nil {
		wcuo.SetTotalSeconds(*i)
	}
	return wcuo
}

// AddTotalSeconds adds i to the "total_seconds" field.
func (wcuo *WakatimeCategoryUpdateOne) AddTotalSeconds(i int64) *WakatimeCategoryUpdateOne {
	wcuo.mutation.AddTotalSeconds(i)
	return wcuo
}

// Mutation returns the WakatimeCategoryMutation object of the builder.
func (wcuo *WakatimeCategoryUpdateOne) Mutation() *WakatimeCategoryMutation {
	return wcuo.mutation
}

// Where appends a list predicates to the WakatimeCategoryUpdate builder.
func (wcuo *WakatimeCategoryUpdateOne) Where(ps ...predicate.WakatimeCategory) *WakatimeCategoryUpdateOne {
	wcuo.mutation.Where(ps...)
	return wcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wcuo *WakatimeCategoryUpdateOne) Select(field string, fields ...string) *WakatimeCategoryUpdateOne {
	wcuo.fields = append([]string{field}, fields...)
	return wcuo
}

// Save executes the query and returns the updated WakatimeCategory entity.
func (wcuo *WakatimeCategoryUpdateOne) Save(ctx context.Context) (*WakatimeCategory, error) {
	wcuo.defaults()
	return withHooks(ctx, wcuo.sqlSave, wcuo.mutation, wcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wcuo *WakatimeCategoryUpdateOne) SaveX(ctx context.Context) *WakatimeCategory {
	node, err := wcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wcuo *WakatimeCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := wcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcuo *WakatimeCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := wcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wcuo *WakatimeCategoryUpdateOne) defaults() {
	if _, ok := wcuo.mutation.UpdateTime(); !ok {
		v := wakatimecategory.UpdateDefaultUpdateTime()
		wcuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wcuo *WakatimeCategoryUpdateOne) check() error {
	if v, ok := wcuo.mutation.Name(); ok {
		if err := wakatimecategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`models: validator failed for field "WakatimeCategory.name": %w`, err)}
		}
	}
	return nil
}

func (wcuo *WakatimeCategoryUpdateOne) sqlSave(ctx context.Context) (_node *WakatimeCategory, err error) {
	if err := wcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(wakatimecategory.Table, wakatimecategory.Columns, sqlgraph.NewFieldSpec(wakatimecategory.FieldID, field.TypeInt64))
	id, ok := wcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "WakatimeCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wakatimecategory.FieldID)
		for _, f := range fields {
			if !wakatimecategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != wakatimecategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wcuo.mutation.UpdateTime(); ok {
		_spec.SetField(wakatimecategory.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := wcuo.mutation.WakatimeID(); ok {
		_spec.SetField(wakatimecategory.FieldWakatimeID, field.TypeUUID, value)
	}
	if value, ok := wcuo.mutation.UserID(); ok {
		_spec.SetField(wakatimecategory.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := wcuo.mutation.Name(); ok {
		_spec.SetField(wakatimecategory.FieldName, field.TypeString, value)
	}
	if value, ok := wcuo.mutation.TotalSeconds(); ok {
		_spec.SetField(wakatimecategory.FieldTotalSeconds, field.TypeInt64, value)
	}
	if value, ok := wcuo.mutation.AddedTotalSeconds(); ok {
		_spec.AddField(wakatimecategory.FieldTotalSeconds, field.TypeInt64, value)
	}
	_node = &WakatimeCategory{config: wcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wakatimecategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wcuo.mutation.done = true
	return _node, nil
}
