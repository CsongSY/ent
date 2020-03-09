// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package planet

import (
	"github.com/facebookincubator/ent"
)

const (
	// Label holds the string label denoting the planet type in the database.
	Label = "planet"
	// FieldID holds the string denoting the id field in the database.
	FieldID   = "id"   // FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name" // FieldAge holds the string denoting the age vertex property in the database.
	FieldAge  = "age"

	// EdgeNeighbors holds the string denoting the neighbors edge name in mutations.
	EdgeNeighbors = "neighbors"

	// Table holds the table name of the planet in the database.
	Table = "planets"
	// NeighborsTable is the table the holds the neighbors relation/edge. The primary key declared below.
	NeighborsTable = "planet_neighbors"
)

// Columns holds all SQL columns for planet fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldAge,
}

var (
	// NeighborsPrimaryKey and NeighborsColumn2 are the table columns denoting the
	// primary key for the neighbors relation (M2M).
	NeighborsPrimaryKey = []string{"planet_id", "neighbor_id"}
)

// Note that the variables below are initialized by the schema package
// on the initialization of the application. Therefore, the schema package
// should be imported in the main as follows:
//
//	import _ "github.com/facebookincubator/ent/entc/integration/privacy/ent/schema"
//
var (
	Hooks [1]ent.Hook
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
