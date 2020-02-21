// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package card

import (
	"time"

	"github.com/facebookincubator/ent"
)

const (
	// Label holds the string label denoting the card type in the database.
	Label = "card"
	// FieldID holds the string denoting the id field in the database.
	FieldID     = "id"     // FieldBoring holds the string denoting the boring vertex property in the database.
	FieldBoring = "boring" // FieldNumber holds the string denoting the number vertex property in the database.
	FieldNumber = "number" // FieldName holds the string denoting the name vertex property in the database.
	FieldName   = "name"

	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// EdgeBestFriend holds the string denoting the best_friend edge name in mutations.
	EdgeBestFriend = "best_friend"

	// Table holds the table name of the card in the database.
	Table = "cards"
	// FriendsTable is the table the holds the friends relation/edge. The primary key declared below.
	FriendsTable = "card_friends"
	// BestFriendTable is the table the holds the best_friend relation/edge.
	BestFriendTable = "cards"
	// BestFriendColumn is the table column denoting the best_friend relation/edge.
	BestFriendColumn = "card_best_friend"
)

// Columns holds all SQL columns for card fields.
var Columns = []string{
	FieldID,
	FieldBoring,
	FieldNumber,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Card type.
var ForeignKeys = []string{
	"card_best_friend",
}

var (
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"card_id", "friend_id"}
)

// Note that the variables below are initialized by the schema package
// on the initialization of the application. Therefore, the schema package
// should be imported in the main as follows:
//
//	import _ "github.com/facebookincubator/ent/entc/integration/hooks/ent/schema"
//
var (
	Hooks [1]ent.Hook
	// DefaultBoring holds the default value on creation for the boring field.
	DefaultBoring func() time.Time
	// DefaultNumber holds the default value on creation for the number field.
	DefaultNumber string
	// NumberValidator is a validator for the "number" field. It is called by the builders before save.
	NumberValidator func(string) error
)
