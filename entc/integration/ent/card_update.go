// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/card"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/ent/spec"
	"github.com/facebookincubator/ent/entc/integration/ent/user"
	"github.com/facebookincubator/ent/schema/field"
)

// CardUpdate is the builder for updating Card entities.
type CardUpdate struct {
	config
	hooks      []Hook
	mutation   *CardMutation
	predicates []predicate.Card
}

// Where adds a new predicate for the builder.
func (cu *CardUpdate) Where(ps ...predicate.Card) *CardUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetName sets the name field.
func (cu *CardUpdate) SetName(s string) *CardUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the name field if the given value is not nil.
func (cu *CardUpdate) SetNillableName(s *string) *CardUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// ClearName clears the value of name.
func (cu *CardUpdate) ClearName() *CardUpdate {
	cu.mutation.ClearName()
	return cu
}

// SetOwnerID sets the owner edge to User by id.
func (cu *CardUpdate) SetOwnerID(id string) *CardUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cu *CardUpdate) SetNillableOwnerID(id *string) *CardUpdate {
	if id != nil {
		cu = cu.SetOwnerID(*id)
	}
	return cu
}

// SetOwner sets the owner edge to User.
func (cu *CardUpdate) SetOwner(u *User) *CardUpdate {
	return cu.SetOwnerID(u.ID)
}

// AddSpecIDs adds the spec edge to Spec by ids.
func (cu *CardUpdate) AddSpecIDs(ids ...string) *CardUpdate {
	cu.mutation.AddSpecIDs(ids...)
	return cu
}

// AddSpec adds the spec edges to Spec.
func (cu *CardUpdate) AddSpec(s ...*Spec) *CardUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddSpecIDs(ids...)
}

// ClearOwner clears the owner edge to User.
func (cu *CardUpdate) ClearOwner() *CardUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// RemoveSpecIDs removes the spec edge to Spec by ids.
func (cu *CardUpdate) RemoveSpecIDs(ids ...string) *CardUpdate {
	cu.mutation.RemoveSpecIDs(ids...)
	return cu
}

// RemoveSpec removes spec edges to Spec.
func (cu *CardUpdate) RemoveSpec(s ...*Spec) *CardUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveSpecIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CardUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := cu.mutation.UpdateTime(); !ok {
		v := card.UpdateDefaultUpdateTime()
		cu.mutation.SetUpdateTime(v)
	}
	if v, ok := cu.mutation.Name(); ok {
		if err := card.NameValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if len(cu.mutation.OwnerIDs()) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			return affected, err
		})
		for _, hook := range cu.hooks {
			mut = hook(mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CardUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CardUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CardUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   card.Table,
			Columns: card.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: card.FieldID,
			},
		},
	}
	if ps := cu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: card.FieldUpdateTime,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: card.FieldName,
		})
	}
	if cu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: card.FieldName,
		})
	}
	if cu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return 0, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := cu.mutation.RemovedSpecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.SpecTable,
			Columns: card.SpecPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: spec.FieldID,
				},
			},
		}
		for _, k := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return 0, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.SpecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.SpecTable,
			Columns: card.SpecPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: spec.FieldID,
				},
			},
		}
		for _, k := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return 0, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CardUpdateOne is the builder for updating a single Card entity.
type CardUpdateOne struct {
	config
	hooks    []Hook
	mutation *CardMutation
}

// SetName sets the name field.
func (cuo *CardUpdateOne) SetName(s string) *CardUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the name field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableName(s *string) *CardUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// ClearName clears the value of name.
func (cuo *CardUpdateOne) ClearName() *CardUpdateOne {
	cuo.mutation.ClearName()
	return cuo
}

// SetOwnerID sets the owner edge to User by id.
func (cuo *CardUpdateOne) SetOwnerID(id string) *CardUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableOwnerID(id *string) *CardUpdateOne {
	if id != nil {
		cuo = cuo.SetOwnerID(*id)
	}
	return cuo
}

// SetOwner sets the owner edge to User.
func (cuo *CardUpdateOne) SetOwner(u *User) *CardUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// AddSpecIDs adds the spec edge to Spec by ids.
func (cuo *CardUpdateOne) AddSpecIDs(ids ...string) *CardUpdateOne {
	cuo.mutation.AddSpecIDs(ids...)
	return cuo
}

// AddSpec adds the spec edges to Spec.
func (cuo *CardUpdateOne) AddSpec(s ...*Spec) *CardUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddSpecIDs(ids...)
}

// ClearOwner clears the owner edge to User.
func (cuo *CardUpdateOne) ClearOwner() *CardUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// RemoveSpecIDs removes the spec edge to Spec by ids.
func (cuo *CardUpdateOne) RemoveSpecIDs(ids ...string) *CardUpdateOne {
	cuo.mutation.RemoveSpecIDs(ids...)
	return cuo
}

// RemoveSpec removes spec edges to Spec.
func (cuo *CardUpdateOne) RemoveSpec(s ...*Spec) *CardUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveSpecIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (cuo *CardUpdateOne) Save(ctx context.Context) (*Card, error) {
	if _, ok := cuo.mutation.UpdateTime(); !ok {
		v := card.UpdateDefaultUpdateTime()
		cuo.mutation.SetUpdateTime(v)
	}
	if v, ok := cuo.mutation.Name(); ok {
		if err := card.NameValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if len(cuo.mutation.OwnerIDs()) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	var (
		err  error
		node *Card
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			return node, err
		})
		for _, hook := range cuo.hooks {
			mut = hook(mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CardUpdateOne) SaveX(ctx context.Context) *Card {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CardUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CardUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CardUpdateOne) sqlSave(ctx context.Context) (c *Card, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   card.Table,
			Columns: card.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: card.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Card.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: card.FieldUpdateTime,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: card.FieldName,
		})
	}
	if cuo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: card.FieldName,
		})
	}
	if cuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := cuo.mutation.RemovedSpecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.SpecTable,
			Columns: card.SpecPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: spec.FieldID,
				},
			},
		}
		for _, k := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.SpecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.SpecTable,
			Columns: card.SpecPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: spec.FieldID,
				},
			},
		}
		for _, k := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	c = &Card{config: cuo.config}
	_spec.Assign = c.assignValues
	_spec.ScanValues = c.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return c, nil
}
