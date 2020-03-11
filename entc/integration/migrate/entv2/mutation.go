// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv2

import (
	"fmt"

	"github.com/facebookincubator/ent/entc/integration/ent"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/car"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/user"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCar   = "Car"
	TypeGroup = "Group"
	TypePet   = "Pet"
	TypeUser  = "User"
)

// CarMutation represents an operation that mutate the Cars
// nodes in the graph.
type CarMutation struct {
	config
	op            Op
	typ           string
	id            *int
	clearedFields map[string]bool
	owner         *int
	clearedowner  bool
}

var _ ent.Mutation = (*CarMutation)(nil)

// newCarMutation creates new mutation for $n.Name.
func newCarMutation(c config, op Op) *CarMutation {
	return &CarMutation{
		config:        c,
		op:            op,
		typ:           TypeCar,
		clearedFields: make(map[string]bool),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CarMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CarMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("entv2: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *CarMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetOwnerID sets the owner edge to User by id.
func (m *CarMutation) SetOwnerID(id int) {
	m.owner = &id
}

// ClearOwner clears the owner edge to User.
func (m *CarMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared returns if the edge owner was cleared.
func (m *CarMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerID returns the owner id in the mutation.
func (m *CarMutation) OwnerID() (id int, exists bool) {
	if m.owner != nil {
		return *m.owner, true
	}
	return
}

// OwnerIDs returns the owner ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *CarMutation) OwnerIDs() (ids []int) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner reset all changes of the owner edge.
func (m *CarMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Op returns the operation name.
func (m *CarMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Car).
func (m *CarMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *CarMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *CarMutation) Field(name string) (ent.Value, bool) {
	switch name {
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *CarMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Car field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *CarMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *CarMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *CarMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Car numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *CarMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *CarMutation) FieldCleared(name string) bool {
	return m.clearedFields[name]
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *CarMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Car nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *CarMutation) ResetField(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Car field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *CarMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, car.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *CarMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case car.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *CarMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *CarMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *CarMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, car.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *CarMutation) EdgeCleared(name string) bool {
	switch name {
	case car.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *CarMutation) ClearEdge(name string) error {
	switch name {
	case car.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Car unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *CarMutation) ResetEdge(name string) error {
	switch name {
	case car.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown Car edge %s", name)
}

// GroupMutation represents an operation that mutate the Groups
// nodes in the graph.
type GroupMutation struct {
	config
	op            Op
	typ           string
	id            *int
	clearedFields map[string]bool
}

var _ ent.Mutation = (*GroupMutation)(nil)

// newGroupMutation creates new mutation for $n.Name.
func newGroupMutation(c config, op Op) *GroupMutation {
	return &GroupMutation{
		config:        c,
		op:            op,
		typ:           TypeGroup,
		clearedFields: make(map[string]bool),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m GroupMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m GroupMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("entv2: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *GroupMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// Op returns the operation name.
func (m *GroupMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Group).
func (m *GroupMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *GroupMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *GroupMutation) Field(name string) (ent.Value, bool) {
	switch name {
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *GroupMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Group field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *GroupMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *GroupMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *GroupMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Group numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *GroupMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *GroupMutation) FieldCleared(name string) bool {
	return m.clearedFields[name]
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *GroupMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Group nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *GroupMutation) ResetField(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Group field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *GroupMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *GroupMutation) AddedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *GroupMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *GroupMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *GroupMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *GroupMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *GroupMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Group unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *GroupMutation) ResetEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Group edge %s", name)
}

// PetMutation represents an operation that mutate the Pets
// nodes in the graph.
type PetMutation struct {
	config
	op            Op
	typ           string
	id            *int
	clearedFields map[string]bool
}

var _ ent.Mutation = (*PetMutation)(nil)

// newPetMutation creates new mutation for $n.Name.
func newPetMutation(c config, op Op) *PetMutation {
	return &PetMutation{
		config:        c,
		op:            op,
		typ:           TypePet,
		clearedFields: make(map[string]bool),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PetMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PetMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("entv2: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *PetMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// Op returns the operation name.
func (m *PetMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Pet).
func (m *PetMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *PetMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *PetMutation) Field(name string) (ent.Value, bool) {
	switch name {
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *PetMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Pet field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *PetMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *PetMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *PetMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Pet numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *PetMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *PetMutation) FieldCleared(name string) bool {
	return m.clearedFields[name]
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *PetMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Pet nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *PetMutation) ResetField(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Pet field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *PetMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *PetMutation) AddedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *PetMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *PetMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *PetMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *PetMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *PetMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Pet unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *PetMutation) ResetEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Pet edge %s", name)
}

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	age           *int
	addage        *int
	name          *string
	nickname      *string
	phone         *string
	buffer        *[]byte
	title         *string
	new_name      *string
	blob          *[]byte
	state         *user.State
	clearedFields map[string]bool
	car           map[int]struct{}
	removedcar    map[int]struct{}
	pets          *int
	clearedpets   bool
}

var _ ent.Mutation = (*UserMutation)(nil)

// newUserMutation creates new mutation for $n.Name.
func newUserMutation(c config, op Op) *UserMutation {
	return &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]bool),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("entv2: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on User creation.
func (m *UserMutation) SetID(id int) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetAge sets the age field.
func (m *UserMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the age value in the mutation.
func (m *UserMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// AddAge adds i to age.
func (m *UserMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the age field in this mutation.
func (m *UserMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge reset all changes of the age field.
func (m *UserMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetName sets the name field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// ResetName reset all changes of the name field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetNickname sets the nickname field.
func (m *UserMutation) SetNickname(s string) {
	m.nickname = &s
}

// Nickname returns the nickname value in the mutation.
func (m *UserMutation) Nickname() (r string, exists bool) {
	v := m.nickname
	if v == nil {
		return
	}
	return *v, true
}

// ResetNickname reset all changes of the nickname field.
func (m *UserMutation) ResetNickname() {
	m.nickname = nil
}

// SetPhone sets the phone field.
func (m *UserMutation) SetPhone(s string) {
	m.phone = &s
}

// Phone returns the phone value in the mutation.
func (m *UserMutation) Phone() (r string, exists bool) {
	v := m.phone
	if v == nil {
		return
	}
	return *v, true
}

// ResetPhone reset all changes of the phone field.
func (m *UserMutation) ResetPhone() {
	m.phone = nil
}

// SetBuffer sets the buffer field.
func (m *UserMutation) SetBuffer(b []byte) {
	m.buffer = &b
}

// Buffer returns the buffer value in the mutation.
func (m *UserMutation) Buffer() (r []byte, exists bool) {
	v := m.buffer
	if v == nil {
		return
	}
	return *v, true
}

// ClearBuffer clears the value of buffer.
func (m *UserMutation) ClearBuffer() {
	m.buffer = nil
	m.clearedFields[user.FieldBuffer] = true
}

// BufferCleared returns if the field buffer was cleared in this mutation.
func (m *UserMutation) BufferCleared() bool {
	return m.clearedFields[user.FieldBuffer]
}

// ResetBuffer reset all changes of the buffer field.
func (m *UserMutation) ResetBuffer() {
	m.buffer = nil
	delete(m.clearedFields, user.FieldBuffer)
}

// SetTitle sets the title field.
func (m *UserMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the title value in the mutation.
func (m *UserMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// ResetTitle reset all changes of the title field.
func (m *UserMutation) ResetTitle() {
	m.title = nil
}

// SetNewName sets the new_name field.
func (m *UserMutation) SetNewName(s string) {
	m.new_name = &s
}

// NewName returns the new_name value in the mutation.
func (m *UserMutation) NewName() (r string, exists bool) {
	v := m.new_name
	if v == nil {
		return
	}
	return *v, true
}

// ClearNewName clears the value of new_name.
func (m *UserMutation) ClearNewName() {
	m.new_name = nil
	m.clearedFields[user.FieldNewName] = true
}

// NewNameCleared returns if the field new_name was cleared in this mutation.
func (m *UserMutation) NewNameCleared() bool {
	return m.clearedFields[user.FieldNewName]
}

// ResetNewName reset all changes of the new_name field.
func (m *UserMutation) ResetNewName() {
	m.new_name = nil
	delete(m.clearedFields, user.FieldNewName)
}

// SetBlob sets the blob field.
func (m *UserMutation) SetBlob(b []byte) {
	m.blob = &b
}

// Blob returns the blob value in the mutation.
func (m *UserMutation) Blob() (r []byte, exists bool) {
	v := m.blob
	if v == nil {
		return
	}
	return *v, true
}

// ClearBlob clears the value of blob.
func (m *UserMutation) ClearBlob() {
	m.blob = nil
	m.clearedFields[user.FieldBlob] = true
}

// BlobCleared returns if the field blob was cleared in this mutation.
func (m *UserMutation) BlobCleared() bool {
	return m.clearedFields[user.FieldBlob]
}

// ResetBlob reset all changes of the blob field.
func (m *UserMutation) ResetBlob() {
	m.blob = nil
	delete(m.clearedFields, user.FieldBlob)
}

// SetState sets the state field.
func (m *UserMutation) SetState(u user.State) {
	m.state = &u
}

// State returns the state value in the mutation.
func (m *UserMutation) State() (r user.State, exists bool) {
	v := m.state
	if v == nil {
		return
	}
	return *v, true
}

// ClearState clears the value of state.
func (m *UserMutation) ClearState() {
	m.state = nil
	m.clearedFields[user.FieldState] = true
}

// StateCleared returns if the field state was cleared in this mutation.
func (m *UserMutation) StateCleared() bool {
	return m.clearedFields[user.FieldState]
}

// ResetState reset all changes of the state field.
func (m *UserMutation) ResetState() {
	m.state = nil
	delete(m.clearedFields, user.FieldState)
}

// AddCarIDs adds the car edge to Car by ids.
func (m *UserMutation) AddCarIDs(ids ...int) {
	if m.car == nil {
		m.car = make(map[int]struct{})
	}
	for i := range ids {
		m.car[ids[i]] = struct{}{}
	}
}

// RemoveCarIDs removes the car edge to Car by ids.
func (m *UserMutation) RemoveCarIDs(ids ...int) {
	if m.removedcar == nil {
		m.removedcar = make(map[int]struct{})
	}
	for i := range ids {
		m.removedcar[ids[i]] = struct{}{}
	}
}

// RemovedCar returns the removed ids of car.
func (m *UserMutation) RemovedCarIDs() (ids []int) {
	for id := range m.removedcar {
		ids = append(ids, id)
	}
	return
}

// CarIDs returns the car ids in the mutation.
func (m *UserMutation) CarIDs() (ids []int) {
	for id := range m.car {
		ids = append(ids, id)
	}
	return
}

// ResetCar reset all changes of the car edge.
func (m *UserMutation) ResetCar() {
	m.car = nil
	m.removedcar = nil
}

// SetPetsID sets the pets edge to Pet by id.
func (m *UserMutation) SetPetsID(id int) {
	m.pets = &id
}

// ClearPets clears the pets edge to Pet.
func (m *UserMutation) ClearPets() {
	m.clearedpets = true
}

// PetsCleared returns if the edge pets was cleared.
func (m *UserMutation) PetsCleared() bool {
	return m.clearedpets
}

// PetsID returns the pets id in the mutation.
func (m *UserMutation) PetsID() (id int, exists bool) {
	if m.pets != nil {
		return *m.pets, true
	}
	return
}

// PetsIDs returns the pets ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// PetsID instead. It exists only for internal usage by the builders.
func (m *UserMutation) PetsIDs() (ids []int) {
	if id := m.pets; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetPets reset all changes of the pets edge.
func (m *UserMutation) ResetPets() {
	m.pets = nil
	m.clearedpets = false
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
	fields := make([]string, 0, 9)
	if m.age != nil {
		fields = append(fields, user.FieldAge)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.nickname != nil {
		fields = append(fields, user.FieldNickname)
	}
	if m.phone != nil {
		fields = append(fields, user.FieldPhone)
	}
	if m.buffer != nil {
		fields = append(fields, user.FieldBuffer)
	}
	if m.title != nil {
		fields = append(fields, user.FieldTitle)
	}
	if m.new_name != nil {
		fields = append(fields, user.FieldNewName)
	}
	if m.blob != nil {
		fields = append(fields, user.FieldBlob)
	}
	if m.state != nil {
		fields = append(fields, user.FieldState)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.Age()
	case user.FieldName:
		return m.Name()
	case user.FieldNickname:
		return m.Nickname()
	case user.FieldPhone:
		return m.Phone()
	case user.FieldBuffer:
		return m.Buffer()
	case user.FieldTitle:
		return m.Title()
	case user.FieldNewName:
		return m.NewName()
	case user.FieldBlob:
		return m.Blob()
	case user.FieldState:
		return m.State()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldNickname:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNickname(v)
		return nil
	case user.FieldPhone:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhone(v)
		return nil
	case user.FieldBuffer:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBuffer(v)
		return nil
	case user.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case user.FieldNewName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNewName(v)
		return nil
	case user.FieldBlob:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBlob(v)
		return nil
	case user.FieldState:
		v, ok := value.(user.State)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetState(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, user.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.clearedFields[user.FieldBuffer] {
		fields = append(fields, user.FieldBuffer)
	}
	if m.clearedFields[user.FieldNewName] {
		fields = append(fields, user.FieldNewName)
	}
	if m.clearedFields[user.FieldBlob] {
		fields = append(fields, user.FieldBlob)
	}
	if m.clearedFields[user.FieldState] {
		fields = append(fields, user.FieldState)
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
	case user.FieldBuffer:
		m.ClearBuffer()
		return nil
	case user.FieldNewName:
		m.ClearNewName()
		return nil
	case user.FieldBlob:
		m.ClearBlob()
		return nil
	case user.FieldState:
		m.ClearState()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldAge:
		m.ResetAge()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldNickname:
		m.ResetNickname()
		return nil
	case user.FieldPhone:
		m.ResetPhone()
		return nil
	case user.FieldBuffer:
		m.ResetBuffer()
		return nil
	case user.FieldTitle:
		m.ResetTitle()
		return nil
	case user.FieldNewName:
		m.ResetNewName()
		return nil
	case user.FieldBlob:
		m.ResetBlob()
		return nil
	case user.FieldState:
		m.ResetState()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.car != nil {
		edges = append(edges, user.EdgeCar)
	}
	if m.pets != nil {
		edges = append(edges, user.EdgePets)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeCar:
		ids := make([]ent.Value, 0, len(m.car))
		for id := range m.car {
			ids = append(ids, id)
		}
		return ids
	case user.EdgePets:
		if id := m.pets; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedcar != nil {
		edges = append(edges, user.EdgeCar)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeCar:
		ids := make([]ent.Value, 0, len(m.removedcar))
		for id := range m.removedcar {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedpets {
		edges = append(edges, user.EdgePets)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgePets:
		return m.clearedpets
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	case user.EdgePets:
		m.ClearPets()
		return nil
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeCar:
		m.ResetCar()
		return nil
	case user.EdgePets:
		m.ResetPets()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
