// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/predicate"
	"ent/models/source"
	"ent/types/auth"
	"ent/utils/timeutil"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SourceUpdate is the builder for updating Source entities.
type SourceUpdate struct {
	config
	hooks    []Hook
	mutation *SourceMutation
}

// Where appends a list predicates to the SourceUpdate builder.
func (su *SourceUpdate) Where(ps ...predicate.Source) *SourceUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetType sets the "type" field.
func (su *SourceUpdate) SetType(a auth.Type) *SourceUpdate {
	su.mutation.ResetType()
	su.mutation.SetType(a)
	return su
}

// SetNillableType sets the "type" field if the given value is not nil.
func (su *SourceUpdate) SetNillableType(a *auth.Type) *SourceUpdate {
	if a != nil {
		su.SetType(*a)
	}
	return su
}

// AddType adds a to the "type" field.
func (su *SourceUpdate) AddType(a auth.Type) *SourceUpdate {
	su.mutation.AddType(a)
	return su
}

// SetName sets the "name" field.
func (su *SourceUpdate) SetName(s string) *SourceUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *SourceUpdate) SetNillableName(s *string) *SourceUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// SetIsActive sets the "is_active" field.
func (su *SourceUpdate) SetIsActive(b bool) *SourceUpdate {
	su.mutation.SetIsActive(b)
	return su
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (su *SourceUpdate) SetNillableIsActive(b *bool) *SourceUpdate {
	if b != nil {
		su.SetIsActive(*b)
	}
	return su
}

// SetIsSyncEnabled sets the "is_sync_enabled" field.
func (su *SourceUpdate) SetIsSyncEnabled(b bool) *SourceUpdate {
	su.mutation.SetIsSyncEnabled(b)
	return su
}

// SetNillableIsSyncEnabled sets the "is_sync_enabled" field if the given value is not nil.
func (su *SourceUpdate) SetNillableIsSyncEnabled(b *bool) *SourceUpdate {
	if b != nil {
		su.SetIsSyncEnabled(*b)
	}
	return su
}

// SetCfg sets the "cfg" field.
func (su *SourceUpdate) SetCfg(a auth.Config[interface{}]) *SourceUpdate {
	su.mutation.SetCfg(a)
	return su
}

// SetNillableCfg sets the "cfg" field if the given value is not nil.
func (su *SourceUpdate) SetNillableCfg(a *auth.Config[interface{}]) *SourceUpdate {
	if a != nil {
		su.SetCfg(*a)
	}
	return su
}

// SetUpdateTime sets the "update_time" field.
func (su *SourceUpdate) SetUpdateTime(ts timeutil.TimeStamp) *SourceUpdate {
	su.mutation.ResetUpdateTime()
	su.mutation.SetUpdateTime(ts)
	return su
}

// AddUpdateTime adds ts to the "update_time" field.
func (su *SourceUpdate) AddUpdateTime(ts timeutil.TimeStamp) *SourceUpdate {
	su.mutation.AddUpdateTime(ts)
	return su
}

// Mutation returns the SourceMutation object of the builder.
func (su *SourceUpdate) Mutation() *SourceMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SourceUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SourceUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SourceUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SourceUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SourceUpdate) defaults() {
	if _, ok := su.mutation.UpdateTime(); !ok {
		v := source.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
	}
}

func (su *SourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(source.Table, source.Columns, sqlgraph.NewFieldSpec(source.FieldID, field.TypeInt64))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.GetType(); ok {
		_spec.SetField(source.FieldType, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedType(); ok {
		_spec.AddField(source.FieldType, field.TypeInt, value)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(source.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.IsActive(); ok {
		_spec.SetField(source.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := su.mutation.IsSyncEnabled(); ok {
		_spec.SetField(source.FieldIsSyncEnabled, field.TypeBool, value)
	}
	if value, ok := su.mutation.Cfg(); ok {
		_spec.SetField(source.FieldCfg, field.TypeJSON, value)
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.SetField(source.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedUpdateTime(); ok {
		_spec.AddField(source.FieldUpdateTime, field.TypeInt64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{source.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SourceUpdateOne is the builder for updating a single Source entity.
type SourceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SourceMutation
}

// SetType sets the "type" field.
func (suo *SourceUpdateOne) SetType(a auth.Type) *SourceUpdateOne {
	suo.mutation.ResetType()
	suo.mutation.SetType(a)
	return suo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (suo *SourceUpdateOne) SetNillableType(a *auth.Type) *SourceUpdateOne {
	if a != nil {
		suo.SetType(*a)
	}
	return suo
}

// AddType adds a to the "type" field.
func (suo *SourceUpdateOne) AddType(a auth.Type) *SourceUpdateOne {
	suo.mutation.AddType(a)
	return suo
}

// SetName sets the "name" field.
func (suo *SourceUpdateOne) SetName(s string) *SourceUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *SourceUpdateOne) SetNillableName(s *string) *SourceUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// SetIsActive sets the "is_active" field.
func (suo *SourceUpdateOne) SetIsActive(b bool) *SourceUpdateOne {
	suo.mutation.SetIsActive(b)
	return suo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (suo *SourceUpdateOne) SetNillableIsActive(b *bool) *SourceUpdateOne {
	if b != nil {
		suo.SetIsActive(*b)
	}
	return suo
}

// SetIsSyncEnabled sets the "is_sync_enabled" field.
func (suo *SourceUpdateOne) SetIsSyncEnabled(b bool) *SourceUpdateOne {
	suo.mutation.SetIsSyncEnabled(b)
	return suo
}

// SetNillableIsSyncEnabled sets the "is_sync_enabled" field if the given value is not nil.
func (suo *SourceUpdateOne) SetNillableIsSyncEnabled(b *bool) *SourceUpdateOne {
	if b != nil {
		suo.SetIsSyncEnabled(*b)
	}
	return suo
}

// SetCfg sets the "cfg" field.
func (suo *SourceUpdateOne) SetCfg(a auth.Config[interface{}]) *SourceUpdateOne {
	suo.mutation.SetCfg(a)
	return suo
}

// SetNillableCfg sets the "cfg" field if the given value is not nil.
func (suo *SourceUpdateOne) SetNillableCfg(a *auth.Config[interface{}]) *SourceUpdateOne {
	if a != nil {
		suo.SetCfg(*a)
	}
	return suo
}

// SetUpdateTime sets the "update_time" field.
func (suo *SourceUpdateOne) SetUpdateTime(ts timeutil.TimeStamp) *SourceUpdateOne {
	suo.mutation.ResetUpdateTime()
	suo.mutation.SetUpdateTime(ts)
	return suo
}

// AddUpdateTime adds ts to the "update_time" field.
func (suo *SourceUpdateOne) AddUpdateTime(ts timeutil.TimeStamp) *SourceUpdateOne {
	suo.mutation.AddUpdateTime(ts)
	return suo
}

// Mutation returns the SourceMutation object of the builder.
func (suo *SourceUpdateOne) Mutation() *SourceMutation {
	return suo.mutation
}

// Where appends a list predicates to the SourceUpdate builder.
func (suo *SourceUpdateOne) Where(ps ...predicate.Source) *SourceUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SourceUpdateOne) Select(field string, fields ...string) *SourceUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Source entity.
func (suo *SourceUpdateOne) Save(ctx context.Context) (*Source, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SourceUpdateOne) SaveX(ctx context.Context) *Source {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SourceUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SourceUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SourceUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdateTime(); !ok {
		v := source.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
	}
}

func (suo *SourceUpdateOne) sqlSave(ctx context.Context) (_node *Source, err error) {
	_spec := sqlgraph.NewUpdateSpec(source.Table, source.Columns, sqlgraph.NewFieldSpec(source.FieldID, field.TypeInt64))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "Source.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, source.FieldID)
		for _, f := range fields {
			if !source.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != source.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.GetType(); ok {
		_spec.SetField(source.FieldType, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedType(); ok {
		_spec.AddField(source.FieldType, field.TypeInt, value)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(source.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.IsActive(); ok {
		_spec.SetField(source.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := suo.mutation.IsSyncEnabled(); ok {
		_spec.SetField(source.FieldIsSyncEnabled, field.TypeBool, value)
	}
	if value, ok := suo.mutation.Cfg(); ok {
		_spec.SetField(source.FieldCfg, field.TypeJSON, value)
	}
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.SetField(source.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedUpdateTime(); ok {
		_spec.AddField(source.FieldUpdateTime, field.TypeInt64, value)
	}
	_node = &Source{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{source.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}