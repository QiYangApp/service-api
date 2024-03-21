// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"ent/models/wakatimelanguage"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WakatimeLanguageCreate is the builder for creating a WakatimeLanguage entity.
type WakatimeLanguageCreate struct {
	config
	mutation *WakatimeLanguageMutation
	hooks    []Hook
}

// Mutation returns the WakatimeLanguageMutation object of the builder.
func (wlc *WakatimeLanguageCreate) Mutation() *WakatimeLanguageMutation {
	return wlc.mutation
}

// Save creates the WakatimeLanguage in the database.
func (wlc *WakatimeLanguageCreate) Save(ctx context.Context) (*WakatimeLanguage, error) {
	return withHooks(ctx, wlc.sqlSave, wlc.mutation, wlc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wlc *WakatimeLanguageCreate) SaveX(ctx context.Context) *WakatimeLanguage {
	v, err := wlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wlc *WakatimeLanguageCreate) Exec(ctx context.Context) error {
	_, err := wlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wlc *WakatimeLanguageCreate) ExecX(ctx context.Context) {
	if err := wlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wlc *WakatimeLanguageCreate) check() error {
	return nil
}

func (wlc *WakatimeLanguageCreate) sqlSave(ctx context.Context) (*WakatimeLanguage, error) {
	if err := wlc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wlc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	wlc.mutation.id = &_node.ID
	wlc.mutation.done = true
	return _node, nil
}

func (wlc *WakatimeLanguageCreate) createSpec() (*WakatimeLanguage, *sqlgraph.CreateSpec) {
	var (
		_node = &WakatimeLanguage{config: wlc.config}
		_spec = sqlgraph.NewCreateSpec(wakatimelanguage.Table, sqlgraph.NewFieldSpec(wakatimelanguage.FieldID, field.TypeInt64))
	)
	return _node, _spec
}

// WakatimeLanguageCreateBulk is the builder for creating many WakatimeLanguage entities in bulk.
type WakatimeLanguageCreateBulk struct {
	config
	err      error
	builders []*WakatimeLanguageCreate
}

// Save creates the WakatimeLanguage entities in the database.
func (wlcb *WakatimeLanguageCreateBulk) Save(ctx context.Context) ([]*WakatimeLanguage, error) {
	if wlcb.err != nil {
		return nil, wlcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wlcb.builders))
	nodes := make([]*WakatimeLanguage, len(wlcb.builders))
	mutators := make([]Mutator, len(wlcb.builders))
	for i := range wlcb.builders {
		func(i int, root context.Context) {
			builder := wlcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WakatimeLanguageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wlcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wlcb *WakatimeLanguageCreateBulk) SaveX(ctx context.Context) []*WakatimeLanguage {
	v, err := wlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wlcb *WakatimeLanguageCreateBulk) Exec(ctx context.Context) error {
	_, err := wlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wlcb *WakatimeLanguageCreateBulk) ExecX(ctx context.Context) {
	if err := wlcb.Exec(ctx); err != nil {
		panic(err)
	}
}