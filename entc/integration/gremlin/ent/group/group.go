// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package group

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID       = "id"        // FieldActive holds the string denoting the active vertex property in the database.
	FieldActive   = "active"    // FieldExpire holds the string denoting the expire vertex property in the database.
	FieldExpire   = "expire"    // FieldType holds the string denoting the type vertex property in the database.
	FieldType     = "type"      // FieldMaxUsers holds the string denoting the max_users vertex property in the database.
	FieldMaxUsers = "max_users" // FieldName holds the string denoting the name vertex property in the database.
	FieldName     = "name"

	// EdgeFiles holds the string denoting the files edge name in mutations.
	EdgeFiles = "files"
	// EdgeBlocked holds the string denoting the blocked edge name in mutations.
	EdgeBlocked = "blocked"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeInfo holds the string denoting the info edge name in mutations.
	EdgeInfo = "info"

	// FilesLabel holds the string label denoting the files edge type in the database.
	FilesLabel = "group_files"
	// BlockedLabel holds the string label denoting the blocked edge type in the database.
	BlockedLabel = "group_blocked"
	// UsersInverseLabel holds the string label denoting the users inverse edge type in the database.
	UsersInverseLabel = "user_groups"
	// InfoLabel holds the string label denoting the info edge type in the database.
	InfoLabel = "group_info"
)

var (
	// DefaultActive holds the default value on creation for the active field.
	DefaultActive bool
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// DefaultMaxUsers holds the default value on creation for the max_users field.
	DefaultMaxUsers int
	// MaxUsersValidator is a validator for the "max_users" field. It is called by the builders before save.
	MaxUsersValidator func(int) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
