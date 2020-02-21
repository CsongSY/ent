// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/examples/o2orecur/ent/node"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeNode = "Node"
)

// NodeMutation represents an operation that mutate the Nodes
// nodes in the graph.
type NodeMutation struct {
	config
	op            Op
	typ           string
	id            *int
	value         *int
	addvalue      *int
	clearedFields map[string]bool
	prev          map[int]struct{}
	clearedprev   bool
	next          map[int]struct{}
	clearednext   bool
}

var _ ent.Mutation = (*NodeMutation)(nil)

// newNodeMutation creates new mutation for $n.Name.
func newNodeMutation(c config, op Op) *NodeMutation {
	return &NodeMutation{
		op:            op,
		typ:           TypeNode,
		clearedFields: make(map[string]bool),
	}
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *NodeMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetValue sets the value field.
func (m *NodeMutation) SetValue(i int) {
	m.value = &i
	m.addvalue = nil
}

// Value returns the value value in the mutation.
func (m *NodeMutation) Value() (r int, exists bool) {
	v := m.value
	if v == nil {
		return
	}
	return *v, true
}

// AddValue adds i to value.
func (m *NodeMutation) AddValue(i int) {
	if m.addvalue != nil {
		*m.addvalue += i
	} else {
		m.addvalue = &i
	}
}

// AddedValue returns the value that was added to the value field in this mutation.
func (m *NodeMutation) AddedValue() (r int, exists bool) {
	v := m.addvalue
	if v == nil {
		return
	}
	return *v, true
}

// ResetValue reset all changes of the value field.
func (m *NodeMutation) ResetValue() {
	m.value = nil
	m.addvalue = nil
}

// SetPrevID sets the prev edge to Node by id.
func (m *NodeMutation) SetPrevID(id int) {
	if m.prev == nil {
		m.prev = make(map[int]struct{})
	}
	m.prev[id] = struct{}{}
}

// ClearPrev clears the prev edge to Node.
func (m *NodeMutation) ClearPrev() {
	m.clearedprev = true
}

// PrevCleared returns if the edge prev was cleared.
func (m *NodeMutation) PrevCleared() bool {
	return m.clearedprev
}

// PrevIDs returns the prev ids in the mutation.
func (m *NodeMutation) PrevIDs() (ids []int) {
	for id := range m.prev {
		ids = append(ids, id)
	}
	return
}

// ResetPrev reset all changes of the prev edge.
func (m *NodeMutation) ResetPrev() {
	m.prev = nil
	m.clearedprev = false
}

// SetNextID sets the next edge to Node by id.
func (m *NodeMutation) SetNextID(id int) {
	if m.next == nil {
		m.next = make(map[int]struct{})
	}
	m.next[id] = struct{}{}
}

// ClearNext clears the next edge to Node.
func (m *NodeMutation) ClearNext() {
	m.clearednext = true
}

// NextCleared returns if the edge next was cleared.
func (m *NodeMutation) NextCleared() bool {
	return m.clearednext
}

// NextIDs returns the next ids in the mutation.
func (m *NodeMutation) NextIDs() (ids []int) {
	for id := range m.next {
		ids = append(ids, id)
	}
	return
}

// ResetNext reset all changes of the next edge.
func (m *NodeMutation) ResetNext() {
	m.next = nil
	m.clearednext = false
}

// Op returns the operation name.
func (m *NodeMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Node).
func (m *NodeMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *NodeMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.value != nil {
		fields = append(fields, node.FieldValue)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *NodeMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case node.FieldValue:
		return m.Value()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *NodeMutation) SetField(name string, value ent.Value) error {
	switch name {
	case node.FieldValue:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetValue(v)
		return nil
	}
	return fmt.Errorf("unknown Node field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *NodeMutation) AddedFields() []string {
	var fields []string
	if m.addvalue != nil {
		fields = append(fields, node.FieldValue)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *NodeMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case node.FieldValue:
		return m.AddedValue()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *NodeMutation) AddField(name string, value ent.Value) error {
	switch name {
	case node.FieldValue:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddValue(v)
		return nil
	}
	return fmt.Errorf("unknown Node numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *NodeMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *NodeMutation) FieldCleared(name string) bool {
	return m.clearedFields[name]
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *NodeMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Node nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *NodeMutation) ResetField(name string) error {
	switch name {
	case node.FieldValue:
		m.ResetValue()
		return nil
	}
	return fmt.Errorf("unknown Node field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *NodeMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.prev != nil {
		edges = append(edges, node.EdgePrev)
	}
	if m.next != nil {
		edges = append(edges, node.EdgeNext)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *NodeMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case node.EdgePrev:
		ids := make([]int, 0, len(m.prev))
		for id := range m.prev {
			ids = append(ids, id)
		}
	case node.EdgeNext:
		ids := make([]int, 0, len(m.next))
		for id := range m.next {
			ids = append(ids, id)
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *NodeMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *NodeMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *NodeMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedprev {
		edges = append(edges, node.EdgePrev)
	}
	if m.clearednext {
		edges = append(edges, node.EdgeNext)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *NodeMutation) EdgeCleared(name string) bool {
	switch name {
	case node.EdgePrev:
		return m.clearedprev
	case node.EdgeNext:
		return m.clearednext
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *NodeMutation) ClearEdge(name string) error {
	switch name {
	case node.EdgePrev:
		m.ClearPrev()
		return nil
	case node.EdgeNext:
		m.ClearNext()
		return nil
	}
	return fmt.Errorf("unknown Node unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *NodeMutation) ResetEdge(name string) error {
	switch name {
	case node.EdgePrev:
		m.ResetPrev()
		return nil
	case node.EdgeNext:
		m.ResetNext()
		return nil
	}
	return fmt.Errorf("unknown Node edge %s", name)
}
