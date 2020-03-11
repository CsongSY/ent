// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/file"
	"github.com/facebookincubator/ent/entc/integration/ent/filetype"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
	"github.com/facebookincubator/ent/schema/field"
)

// FileTypeUpdate is the builder for updating FileType entities.
type FileTypeUpdate struct {
	config
	hooks      []Hook
	mutation   *FileTypeMutation
	predicates []predicate.FileType
}

// Where adds a new predicate for the builder.
func (ftu *FileTypeUpdate) Where(ps ...predicate.FileType) *FileTypeUpdate {
	ftu.predicates = append(ftu.predicates, ps...)
	return ftu
}

// SetName sets the name field.
func (ftu *FileTypeUpdate) SetName(s string) *FileTypeUpdate {
	ftu.mutation.SetName(s)
	return ftu
}

// AddFileIDs adds the files edge to File by ids.
func (ftu *FileTypeUpdate) AddFileIDs(ids ...string) *FileTypeUpdate {
	ftu.mutation.AddFileIDs(ids...)
	return ftu
}

// AddFiles adds the files edges to File.
func (ftu *FileTypeUpdate) AddFiles(f ...*File) *FileTypeUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.AddFileIDs(ids...)
}

// RemoveFileIDs removes the files edge to File by ids.
func (ftu *FileTypeUpdate) RemoveFileIDs(ids ...string) *FileTypeUpdate {
	ftu.mutation.RemoveFileIDs(ids...)
	return ftu
}

// RemoveFiles removes files edges to File.
func (ftu *FileTypeUpdate) RemoveFiles(f ...*File) *FileTypeUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.RemoveFileIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ftu *FileTypeUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(ftu.hooks) == 0 {
		affected, err = ftu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ftu.mutation = mutation
			affected, err = ftu.sqlSave(ctx)
			return affected, err
		})
		for _, hook := range ftu.hooks {
			mut = hook(mut)
		}
		if _, err := mut.Mutate(ctx, ftu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ftu *FileTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ftu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ftu *FileTypeUpdate) Exec(ctx context.Context) error {
	_, err := ftu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftu *FileTypeUpdate) ExecX(ctx context.Context) {
	if err := ftu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ftu *FileTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   filetype.Table,
			Columns: filetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: filetype.FieldID,
			},
		},
	}
	if ps := ftu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filetype.FieldName,
		})
	}
	if nodes := ftu.mutation.RemovedFilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: file.FieldID,
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
	if nodes := ftu.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: file.FieldID,
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
	if n, err = sqlgraph.UpdateNodes(ctx, ftu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filetype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FileTypeUpdateOne is the builder for updating a single FileType entity.
type FileTypeUpdateOne struct {
	config
	hooks    []Hook
	mutation *FileTypeMutation
}

// SetName sets the name field.
func (ftuo *FileTypeUpdateOne) SetName(s string) *FileTypeUpdateOne {
	ftuo.mutation.SetName(s)
	return ftuo
}

// AddFileIDs adds the files edge to File by ids.
func (ftuo *FileTypeUpdateOne) AddFileIDs(ids ...string) *FileTypeUpdateOne {
	ftuo.mutation.AddFileIDs(ids...)
	return ftuo
}

// AddFiles adds the files edges to File.
func (ftuo *FileTypeUpdateOne) AddFiles(f ...*File) *FileTypeUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.AddFileIDs(ids...)
}

// RemoveFileIDs removes the files edge to File by ids.
func (ftuo *FileTypeUpdateOne) RemoveFileIDs(ids ...string) *FileTypeUpdateOne {
	ftuo.mutation.RemoveFileIDs(ids...)
	return ftuo
}

// RemoveFiles removes files edges to File.
func (ftuo *FileTypeUpdateOne) RemoveFiles(f ...*File) *FileTypeUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.RemoveFileIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (ftuo *FileTypeUpdateOne) Save(ctx context.Context) (*FileType, error) {

	var (
		err  error
		node *FileType
	)
	if len(ftuo.hooks) == 0 {
		node, err = ftuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ftuo.mutation = mutation
			node, err = ftuo.sqlSave(ctx)
			return node, err
		})
		for _, hook := range ftuo.hooks {
			mut = hook(mut)
		}
		if _, err := mut.Mutate(ctx, ftuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ftuo *FileTypeUpdateOne) SaveX(ctx context.Context) *FileType {
	ft, err := ftuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ft
}

// Exec executes the query on the entity.
func (ftuo *FileTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ftuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftuo *FileTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ftuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ftuo *FileTypeUpdateOne) sqlSave(ctx context.Context) (ft *FileType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   filetype.Table,
			Columns: filetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: filetype.FieldID,
			},
		},
	}
	id, ok := ftuo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing FileType.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := ftuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filetype.FieldName,
		})
	}
	if nodes := ftuo.mutation.RemovedFilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: file.FieldID,
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
	if nodes := ftuo.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: file.FieldID,
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
	ft = &FileType{config: ftuo.config}
	_spec.Assign = ft.assignValues
	_spec.ScanValues = ft.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ftuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filetype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return ft, nil
}
