// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/facebookincubator/ent/entc/integration/ent"
	"github.com/facebookincubator/ent/entc/integration/json/ent/user"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	url           **url.URL
	raw           *json.RawMessage
	dirs          *[]http.Dir
	ints          *[]int
	floats        *[]float64
	strings       *[]string
	clearedFields map[string]bool
}

var _ ent.Mutation = (*UserMutation)(nil)

// newUserMutation creates new mutation for $n.Name.
func newUserMutation(c config, op Op) *UserMutation {
	return &UserMutation{
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]bool),
	}
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetURL sets the url field.
func (m *UserMutation) SetURL(u *url.URL) {
	m.url = &u
}

// URL returns the url value in the mutation.
func (m *UserMutation) URL() (r *url.URL, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// ClearURL clears the value of url.
func (m *UserMutation) ClearURL() {
	m.url = nil
	m.clearedFields[user.FieldURL] = true
}

// URLCleared returns if the field url was cleared in this mutation.
func (m *UserMutation) URLCleared() bool {
	return m.clearedFields[user.FieldURL]
}

// ResetURL reset all changes of the url field.
func (m *UserMutation) ResetURL() {
	m.url = nil
	delete(m.clearedFields, user.FieldURL)
}

// SetRaw sets the raw field.
func (m *UserMutation) SetRaw(jm json.RawMessage) {
	m.raw = &jm
}

// Raw returns the raw value in the mutation.
func (m *UserMutation) Raw() (r json.RawMessage, exists bool) {
	v := m.raw
	if v == nil {
		return
	}
	return *v, true
}

// ClearRaw clears the value of raw.
func (m *UserMutation) ClearRaw() {
	m.raw = nil
	m.clearedFields[user.FieldRaw] = true
}

// RawCleared returns if the field raw was cleared in this mutation.
func (m *UserMutation) RawCleared() bool {
	return m.clearedFields[user.FieldRaw]
}

// ResetRaw reset all changes of the raw field.
func (m *UserMutation) ResetRaw() {
	m.raw = nil
	delete(m.clearedFields, user.FieldRaw)
}

// SetDirs sets the dirs field.
func (m *UserMutation) SetDirs(h []http.Dir) {
	m.dirs = &h
}

// Dirs returns the dirs value in the mutation.
func (m *UserMutation) Dirs() (r []http.Dir, exists bool) {
	v := m.dirs
	if v == nil {
		return
	}
	return *v, true
}

// ClearDirs clears the value of dirs.
func (m *UserMutation) ClearDirs() {
	m.dirs = nil
	m.clearedFields[user.FieldDirs] = true
}

// DirsCleared returns if the field dirs was cleared in this mutation.
func (m *UserMutation) DirsCleared() bool {
	return m.clearedFields[user.FieldDirs]
}

// ResetDirs reset all changes of the dirs field.
func (m *UserMutation) ResetDirs() {
	m.dirs = nil
	delete(m.clearedFields, user.FieldDirs)
}

// SetInts sets the ints field.
func (m *UserMutation) SetInts(i []int) {
	m.ints = &i
}

// Ints returns the ints value in the mutation.
func (m *UserMutation) Ints() (r []int, exists bool) {
	v := m.ints
	if v == nil {
		return
	}
	return *v, true
}

// ClearInts clears the value of ints.
func (m *UserMutation) ClearInts() {
	m.ints = nil
	m.clearedFields[user.FieldInts] = true
}

// IntsCleared returns if the field ints was cleared in this mutation.
func (m *UserMutation) IntsCleared() bool {
	return m.clearedFields[user.FieldInts]
}

// ResetInts reset all changes of the ints field.
func (m *UserMutation) ResetInts() {
	m.ints = nil
	delete(m.clearedFields, user.FieldInts)
}

// SetFloats sets the floats field.
func (m *UserMutation) SetFloats(f []float64) {
	m.floats = &f
}

// Floats returns the floats value in the mutation.
func (m *UserMutation) Floats() (r []float64, exists bool) {
	v := m.floats
	if v == nil {
		return
	}
	return *v, true
}

// ClearFloats clears the value of floats.
func (m *UserMutation) ClearFloats() {
	m.floats = nil
	m.clearedFields[user.FieldFloats] = true
}

// FloatsCleared returns if the field floats was cleared in this mutation.
func (m *UserMutation) FloatsCleared() bool {
	return m.clearedFields[user.FieldFloats]
}

// ResetFloats reset all changes of the floats field.
func (m *UserMutation) ResetFloats() {
	m.floats = nil
	delete(m.clearedFields, user.FieldFloats)
}

// SetStrings sets the strings field.
func (m *UserMutation) SetStrings(s []string) {
	m.strings = &s
}

// Strings returns the strings value in the mutation.
func (m *UserMutation) Strings() (r []string, exists bool) {
	v := m.strings
	if v == nil {
		return
	}
	return *v, true
}

// ClearStrings clears the value of strings.
func (m *UserMutation) ClearStrings() {
	m.strings = nil
	m.clearedFields[user.FieldStrings] = true
}

// StringsCleared returns if the field strings was cleared in this mutation.
func (m *UserMutation) StringsCleared() bool {
	return m.clearedFields[user.FieldStrings]
}

// ResetStrings reset all changes of the strings field.
func (m *UserMutation) ResetStrings() {
	m.strings = nil
	delete(m.clearedFields, user.FieldStrings)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.url != nil {
		fields = append(fields, user.FieldURL)
	}
	if m.raw != nil {
		fields = append(fields, user.FieldRaw)
	}
	if m.dirs != nil {
		fields = append(fields, user.FieldDirs)
	}
	if m.ints != nil {
		fields = append(fields, user.FieldInts)
	}
	if m.floats != nil {
		fields = append(fields, user.FieldFloats)
	}
	if m.strings != nil {
		fields = append(fields, user.FieldStrings)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldURL:
		return m.URL()
	case user.FieldRaw:
		return m.Raw()
	case user.FieldDirs:
		return m.Dirs()
	case user.FieldInts:
		return m.Ints()
	case user.FieldFloats:
		return m.Floats()
	case user.FieldStrings:
		return m.Strings()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldURL:
		v, ok := value.(*url.URL)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	case user.FieldRaw:
		v, ok := value.(json.RawMessage)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRaw(v)
		return nil
	case user.FieldDirs:
		v, ok := value.([]http.Dir)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDirs(v)
		return nil
	case user.FieldInts:
		v, ok := value.([]int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetInts(v)
		return nil
	case user.FieldFloats:
		v, ok := value.([]float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFloats(v)
		return nil
	case user.FieldStrings:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStrings(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.clearedFields[user.FieldURL] {
		fields = append(fields, user.FieldURL)
	}
	if m.clearedFields[user.FieldRaw] {
		fields = append(fields, user.FieldRaw)
	}
	if m.clearedFields[user.FieldDirs] {
		fields = append(fields, user.FieldDirs)
	}
	if m.clearedFields[user.FieldInts] {
		fields = append(fields, user.FieldInts)
	}
	if m.clearedFields[user.FieldFloats] {
		fields = append(fields, user.FieldFloats)
	}
	if m.clearedFields[user.FieldStrings] {
		fields = append(fields, user.FieldStrings)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	return m.clearedFields[name]
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldURL:
		m.ClearURL()
		return nil
	case user.FieldRaw:
		m.ClearRaw()
		return nil
	case user.FieldDirs:
		m.ClearDirs()
		return nil
	case user.FieldInts:
		m.ClearInts()
		return nil
	case user.FieldFloats:
		m.ClearFloats()
		return nil
	case user.FieldStrings:
		m.ClearStrings()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldURL:
		m.ResetURL()
		return nil
	case user.FieldRaw:
		m.ResetRaw()
		return nil
	case user.FieldDirs:
		m.ResetDirs()
		return nil
	case user.FieldInts:
		m.ResetInts()
		return nil
	case user.FieldFloats:
		m.ResetFloats()
		return nil
	case user.FieldStrings:
		m.ResetStrings()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User edge %s", name)
}
